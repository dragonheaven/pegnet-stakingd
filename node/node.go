package node

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"

	"../config"
	"../node/pegnet"
	"github.com/Factom-Asset-Tokens/factom"
	"github.com/pegnet/pegnet/modules/grader"
	"github.com/spf13/viper"
)

var (
	PegnetActivation uint32 = 206421
)

type Pegnetd struct {
	FactomClient *factom.Client
	Config       *viper.Viper

	Sync   *pegnet.BlockSync
	Pegnet *pegnet.Pegnet
}

func NewPegnetStakingd(ctx context.Context, conf *viper.Viper) (*Pegnetd, error) {
	n := new(Pegnetd)
	n.FactomClient = FactomClientFromConfig(conf)
	n.Config = conf

	n.Pegnet = pegnet.New(conf)
	if err := n.Pegnet.Init(); err != nil {
		return nil, err
	}

	if sync, err := n.Pegnet.SelectSynced(ctx, n.Pegnet.DB); err != nil {
		if err == sql.ErrNoRows {
			n.Sync = new(pegnet.BlockSync)
			n.Sync.Synced = PegnetActivation
			log.Debug("connected to a fresh database")
		} else {
			return nil, err
		}
	} else {
		n.Sync = sync
	}
	/*
		err := n.Pegnet.CheckHardForks(n.Pegnet.DB)
		if err != nil {
			err = fmt.Errorf("pegnetd database hardfork check failed: %s", err.Error())
			if conf.GetBool(config.DisableHardForkCheck) {
				log.Warnf(err.Error())
			} else {
				return nil, err
			}
		}
	*/

	grader.InitLX()
	return n, nil
}

func FactomClientFromConfig(conf *viper.Viper) *factom.Client {
	cl := factom.NewClient()
	cl.FactomdServer = conf.GetString(config.Server)
	cl.WalletdServer = conf.GetString(config.Wallet)
	if config.WalletUser != "" {
		cl.Walletd.BasicAuth = true
		cl.Walletd.User = conf.GetString(config.WalletUser)
		cl.Walletd.Password = conf.GetString(config.WalletPass)
	}

	return cl
}
