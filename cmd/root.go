/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	//"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"

	"github.com/dragonheaven/pegnet-stakingd/config"
	"github.com/dragonheaven/pegnet-stakingd/exit"
	log "github.com/sirupsen/logrus"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:              "pegnet-stakingd",
	Short:            "pegnet-stakingd is the pegnet staking daemon to track balances and SPRs",
	PersistentPreRun: always,
	PreRun:           ReadConfig,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Staking!!!!")
		/*
			// Handle ctl+c
			ctx, cancel := context.WithCancel(context.Background())
			exit.GlobalExitHandler.AddCancel(cancel)

			// Get the config
			conf := viper.GetViper()
			node, err := node.NewPegnetd(ctx, conf)
			if err != nil {
				log.WithError(err).Errorf("failed to launch pegnet node")
				os.Exit(1)
			}

			apiserver := srv.NewAPIServer(conf, node)
			go apiserver.Start(ctx.Done())

			// Run
			node.DBlockSync(ctx)
		*/
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// always is run before any command
func always(cmd *cobra.Command, args []string) {
	log.Infof("always is called")

	// Setup config reading
	viper.SetConfigName("pegnet-stakingd-conf")
	// Add as many config paths as we want to check
	viper.AddConfigPath("$HOME/.pegnet-stakingd")
	viper.AddConfigPath(".")

	// Also init some defaults
	viper.SetDefault(config.DBlockSyncRetryPeriod, time.Second*5)
	viper.SetDefault(config.SqliteDBPath, "$HOME/.pegnet-stakingd/mainnet/sql.db")
}

// ReadConfig can be put as a PreRun for a command that uses the config file
func ReadConfig(cmd *cobra.Command, args []string) {
	log.Infof("ReadConfig is called")
}

func init() {
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.newApp.yaml)")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
