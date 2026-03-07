package migrate

import "github.com/spf13/cobra"

func NewMigrateCmd() *cobra.Command {
	var serveCmd = &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			// daemon.Serve()
		},
	}
	serveCmd.Flags().StringP("config", "c", "./config/config.yaml", "Config path")
	return serveCmd
}
