package migrate

import (
	"github.com/akaporn-katip/go-project-structure-template/config"
	"github.com/akaporn-katip/go-project-structure-template/migrate"
	"github.com/spf13/cobra"
)

func NewMigrateCmd() *cobra.Command {
	var migrateCmd = &cobra.Command{
		Use: "migrate",
		RunE: func(cmd *cobra.Command, args []string) error {
			configVal, _ := cmd.Flags().GetString("config")
			cfg, _ := config.LoadWithPath(configVal)

			dbType := cfg.Database.Type
			dsn := cfg.Database.DSN

			mt, err := migrate.NewMigrator(dbType, dsn)
			if err != nil {
				return err
			}
			return mt.Up()

		},
	}

	return migrateCmd
}
