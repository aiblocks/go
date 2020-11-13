package cmd

import (
	"fmt"
	stdLog "log"
	"os"

	"github.com/spf13/cobra"
	millennium "github.com/aiblocks/go/services/millennium/internal"
)

var (
	config, flags = millennium.Flags()

	rootCmd = &cobra.Command{
		Use:   "millennium",
		Short: "client-facing api server for the aiblocks network",
		Long:  "client-facing api server for the aiblocks network. It acts as the interface between AiBlocks Core and applications that want to access the AiBlocks network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams and more.",
		Run: func(cmd *cobra.Command, args []string) {
			millennium.NewAppFromFlags(config, flags).Serve()
		},
	}
)

func init() {
	err := flags.Init(rootCmd)
	if err != nil {
		stdLog.Fatal(err.Error())
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
