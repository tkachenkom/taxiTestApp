package main

import (
	"database/sql"
	"fmt"
	dbx "github.com/go-ozzo/ozzo-dbx"
	app "github.com/tkachenkom/taxiTestApp"
	"github.com/tkachenkom/taxiTestApp/db/migrator"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/tkachenkom/taxiTestApp/config"
)

// Migrator method necessar for mariadb migrations
type Migrator func(*sql.DB, migrator.MigrateDir) (int, error)

// MigrateDB setup
func MigrateDB(direction string, dbClient *dbx.DB, migratorFn Migrator) (int, error) {
	applied, err := migratorFn(dbClient.DB(), migrator.MigrateDir(direction))

	return applied, errors.Wrap(err, "failed to apply migrations")
}

func main() {
	apiConfig := config.New()
	log := apiConfig.Log()

	rootCmd := &cobra.Command{}

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "run command",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := MigrateDB("up", apiConfig.DB().DBX(), migrator.Migrations.Migrate)

			if err != nil {
				log.WithError(err).Error("migration failed")
				return
			}

			api := app.New(apiConfig)
			if err := api.Start(); err != nil {
				panic(errors.Wrap(err, "failed to start api"))
			}
		},
	}

	rootCmd.AddCommand(runCmd)
	if err := rootCmd.Execute(); err != nil {
		log.WithField("cobra", "read").Error(fmt.Sprintf("failed to read command %s", err.Error()))
		return
	}
}
