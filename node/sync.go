package node

import (
	"../config"

	"context"
	"fmt"
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
		fmt.Println("retryPeriod:", retryPeriod, ", isFirstSync:", isFirstSync)
		continue OuterSyncLoop
	}
}

func isDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
