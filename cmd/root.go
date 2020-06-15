package cmd

import (
	"../config"
	"../exit"
	"../node"
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var rootCmd = &cobra.Command{
	Use:              "pegnet-stakingd",
	Short:            "pegnet-stakingd is the pegnet staking daemon to track balances and SPRs",
	PersistentPreRun: always,
	PreRun:           ReadConfig,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Staking!!!!")
		// Handle ctl+c
		ctx, cancel := context.WithCancel(context.Background())
		exit.GlobalExitHandler.AddCancel(cancel)

		// Get the config
		conf := viper.GetViper()
		node, err := node.NewPegnetStakingd(ctx, conf)
		if err != nil {
			log.WithError(err).Errorf("failed to launch pegnet staking node")
			os.Exit(1)
		}
		fmt.Println(node)

		/*
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
	// Setup config reading
	viper.SetConfigName("pegnet-stakingd-conf")
	// Add as many config paths as we want to check
	viper.AddConfigPath("$HOME/.pegnet-stakingd")
	viper.AddConfigPath(".")

	// Also init some defaults
	viper.SetDefault(config.DBlockSyncRetryPeriod, time.Second*5)
	viper.SetDefault(config.SqliteDBPath, "$HOME/.pegnet-stakingd/mainnet/sql.db")

	// Catch ctl+c
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		<-signalChan
		log.Info("Gracefully closing")
		exit.GlobalExitHandler.Close()

		log.Info("closing application")
		// If something is hanging, we have to kill it
		os.Exit(0)
	}()
}

// ReadConfig can be put as a PreRun for a command that uses the config file
func ReadConfig(cmd *cobra.Command, args []string) {
	err := viper.ReadInConfig()
	// If no config is found, we will attempt to make one
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		log.WithError(err).Fatal("config file not found")
	} else if err != nil {
		log.WithError(err).Fatal("failed to load config")
	}

	// Indicate which config was used
	log.Infof("Using config from %s", viper.ConfigFileUsed())

	initLogger()
}

func init() {
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.newApp.yaml)")
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initLogger() {
	switch strings.ToLower(viper.GetString(config.LoggingLevel)) {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	}
}
