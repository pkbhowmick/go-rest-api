package cmd

import (
	"github.com/pkbhowmick/go-rest-api/api"
	"github.com/spf13/cobra"
)

var port string

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "This flag sets the port of the server")
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "This command will start the api server",
	Long:  "This command will start the go-rest-api server",
	Run: func(cmd *cobra.Command, args []string) {
		api.SetFlags(port)
		api.StartServer()
	},
}
