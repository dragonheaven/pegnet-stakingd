package pegnet

import (
	"database/sql"
	"github.com/spf13/viper"
)

type Pegnet struct {
	Config *viper.Viper

	// This is the sqlite db to store state
	DB *sql.DB
}

func New(conf *viper.Viper) *Pegnet {
	p := new(Pegnet)
	p.Config = conf
	return p
}
