package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	millenniumclient "github.com/aiblocks/go/clients/millenniumclient"
	hlog "github.com/aiblocks/go/support/log"
)

var DatabaseURL string
var Client *millenniumclient.Client
var UseTestNet bool
var Logger = hlog.New()

var defaultDatabaseURL = getEnv("DB_URL", "postgres://localhost:5432/aiblocksticker01?sslmode=disable")

var rootCmd = &cobra.Command{
	Use:   "ticker",
	Short: "AiBlocks Development Foundation Ticker.",
	Long:  `A tool to provide AiBlocks Asset and Market data.`,
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(
		&DatabaseURL,
		"db-url",
		"d",
		defaultDatabaseURL,
		"database URL, such as: postgres://user:pass@localhost:5432/ticker",
	)
	rootCmd.PersistentFlags().BoolVar(
		&UseTestNet,
		"testnet",
		false,
		"use the AiBlocks Test Network, instead of the AiBlocks Public Network",
	)

	Logger.SetLevel(logrus.DebugLevel)
}

func initConfig() {
	if UseTestNet {
		Logger.Debug("Using AiBlocks Default Test Network")
		Client = millenniumclient.DefaultTestNetClient
	} else {
		Logger.Debug("Using AiBlocks Default Public Network")
		Client = millenniumclient.DefaultPublicNetClient
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
