package cmd

import (
	"github.com/spf13/cobra"
	millennium "github.com/aiblocks/go/services/millennium/internal"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run millennium server",
	Long:  "serve initializes then starts the millennium HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		millennium.NewAppFromFlags(config, flags).Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
