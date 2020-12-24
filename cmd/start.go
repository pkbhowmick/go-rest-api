package cmd

import (
	"github.com/pkbhowmick/go-rest-api/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "This command will start the api server",
	Long:  "This command will start the go-rest-api server",
	Run: func(cmd *cobra.Command, args []string) {
		api.StartServer()
	},
}
