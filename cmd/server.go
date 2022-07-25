package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xiaowenxia/nat/server"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Long:  `server`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
