package node

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Factom-Asset-Tokens/factom"
	"github.com/pegnet/pegnetd/config"
	log "github.com/sirupsen/logrus"
)

// DBlockSync iterates through dblocks and syncs the various chains
func (d *Pegnetd) DBlockSync(ctx context.Context) {
	fmt.Println("DBlockSync...")

	retryPeriod := d.Config.GetDuration(config.DBlockSyncRetryPeriod)
	isFirstSync := true

OuterSyncLoop:
	for {
		if isDone(ctx) {
			return // If the user does ctl+c or something
		}

		// Fetch the current highest height
		heights := new(factom.Heights)
		err := heights.Get(nil, d.FactomClient)
		if err != nil {
			log.WithError(err).WithFields(log.Fields{}).Errorf("failed to fetch heights")
			time.Sleep(retryPeriod)
			continue // Loop will just keep retrying until factomd is reached
		}

		if d.Sync.Synced >= heights.DirectoryBlock {
			// We are currently synced, nothing to do. If we are above it, the factomd could
			// be rebooted
			if d.Sync.Synced > heights.DirectoryBlock {
				log.Debugf("Factom node behind. database height = %d, factom height = %d", d.Sync.Synced, heights.DirectoryBlock)
			}

			if isFirstSync {
				isFirstSync = false
				log.WithField("height", d.Sync.Synced).Info("Node is up to date")
			}

			time.Sleep(retryPeriod) // TODO: Should we have a separate polling period?
			continue
		}

		var totalDur time.Duration
		var iterations int

		var longSync bool
		if isFirstSync || heights.DirectoryBlock-d.Sync.Synced > 1 {
			log.WithFields(log.Fields{
				"height":     d.Sync.Synced,
				"syncing-to": heights.DirectoryBlock,
			}).Infof("Starting sync job of %d blocks", heights.DirectoryBlock-d.Sync.Synced)
			longSync = true
		}

		begin := time.Now()
		lastReport := begin
		for d.Sync.Synced < heights.DirectoryBlock {
			start := time.Now()
			hLog := log.WithFields(log.Fields{"height": d.Sync.Synced + 1})
			if isDone(ctx) {
				return
			}

			// start transaction for all block actions
			tx, err := d.Pegnet.DB.BeginTx(ctx, nil)
			if err != nil {
				hLog.WithError(err).Errorf("failed to start transaction")
				continue
			}

			// We are not synced, so we need to iterate through the dblocks and sync them
			// one by one. We can only sync our current synced height +1
			if err := d.SyncBlock(ctx, tx, d.Sync.Synced+1); err != nil {
				hLog.WithError(err).Errorf("failed to sync height")
				time.Sleep(retryPeriod)

				// If we fail, we backout to the outer loop. This allows error handling on factomd state to be a bit
				// cleaner, such as a rebooted node with a different db. That node would have a new heights response.
				err = tx.Rollback()
				if err != nil {
					// TODO evaluate if we can recover from this point or not
					hLog.WithError(err).Fatal("unable to roll back transaction")
				}
				continue OuterSyncLoop
			}

			// Bump our sync, and march forward
			d.Sync.Synced++
			err = d.Pegnet.InsertSynced(tx, d.Sync)
			if err != nil {
				d.Sync.Synced--
				hLog.WithError(err).Errorf("unable to update synced metadata")
				err = tx.Rollback()
				if err != nil {
					// TODO evaluate if we can recover from this point or not
					hLog.WithError(err).Fatal("unable to roll back transaction")
				}
				continue OuterSyncLoop
			}

			err = tx.Commit()
			if err != nil {
				d.Sync.Synced--
				hLog.WithError(err).Errorf("unable to commit transaction")
				err = tx.Rollback()
				if err != nil {
					// TODO evaluate if we can recover from this point or not
					hLog.WithError(err).Fatal("unable to roll back transaction")
				}
			}

			elapsed := time.Since(start)
			hLog.WithFields(log.Fields{"took": elapsed}).Debugf("synced")

			iterations++
			totalDur += elapsed
			// update every 15 seconds
			if time.Since(lastReport) > time.Second*15 {
				lastReport = time.Now()
				toGo := heights.DirectoryBlock - d.Sync.Synced
				avg := totalDur / time.Duration(iterations)
				hLog.WithFields(log.Fields{
					"avg":        avg,
					"left":       time.Duration(toGo) * avg,
					"syncing-to": heights.DirectoryBlock,
					"elapsed":    time.Since(begin),
				}).Infof("sync stats")
			}
		}

		isFirstSync = false
		if longSync {
			longSync = false
			log.WithField("height", d.Sync.Synced).WithField("blocks-synced", iterations).Infof("Finished sync job")
		} else if d.Sync.Synced%6 == 0 {
			log.WithField("height", d.Sync.Synced).Infof("status report")
		}

		continue OuterSyncLoop
	}
}

// If SyncBlock returns no error, than that height was synced and saved. If any part of the sync fails,
// the whole sync should be rolled back and not applied. An error should then be returned.
// The context should be respected if it is cancelled
func (d *Pegnetd) SyncBlock(ctx context.Context, tx *sql.Tx, height uint32) error {
	fmt.Println("============================================================")
	//fLog := log.WithFields(log.Fields{"height": height})
	if isDone(ctx) { // Just an example about how to handle it being cancelled
		return context.Canceled
	}

	dblock := new(factom.DBlock)
	dblock.Height = height
	if err := dblock.Get(nil, d.FactomClient); err != nil {
		return err
	}

	// First, gather all entries we need from factomd
	oprEBlock := dblock.EBlock(OPRChain)
	if oprEBlock != nil {
		if err := multiFetch(oprEBlock, d.FactomClient); err != nil {
			return err
		}
	}
	transactionsEBlock := dblock.EBlock(TransactionChain)
	if transactionsEBlock != nil {
		if err := multiFetch(transactionsEBlock, d.FactomClient); err != nil {
			return err
		}
	}

	return errors.New("Skip the SyncBlock")
}

func multiFetch(eblock *factom.EBlock, c *factom.Client) error {
	err := eblock.Get(nil, c)
	if err != nil {
		return err
	}

	work := make(chan int, len(eblock.Entries))
	defer close(work)
	errs := make(chan error)
	defer close(errs)

	for i := 0; i < 8; i++ {
		go func() {
			// TODO: Fix the channels such that a write on a closed channel never happens.
			//		For now, just kill the worker go routine
			defer func() {
				recover()
			}()

			for j := range work {
				errs <- eblock.Entries[j].Get(nil, c)
			}
		}()
	}

	for i := range eblock.Entries {
		work <- i
	}

	count := 0
	for e := range errs {
		count++
		if e != nil {
			// If we return, we close the errs channel, and the working go routine will
			// still try to write to it.
			return e
		}
		if count == len(eblock.Entries) {
			break
		}
	}

	return nil
}

func isDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
