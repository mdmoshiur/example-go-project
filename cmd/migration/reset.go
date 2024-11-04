package migration

import (
	"github.com/mdmoshiur/example-go/internal/conn"
	"github.com/mdmoshiur/example-go/internal/logger"
	"github.com/mdmoshiur/example-go/internal/migration"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets the tables in database",
	Long:  `Resets the tables in database`,
	Run:   resetDatabase,
}

func init() {
	RootCmd.AddCommand(resetCmd)
}

func resetDatabase(_ *cobra.Command, _ []string) {
	confirmed := askForConfirmation("Migration reset")
	if confirmed {
		logger.Info("Resetting database...")
		db := conn.DefaultDB()

		err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Migrator().DropTable(migration.Models...); err != nil {
				return err
			}

			return tx.AutoMigrate(migration.Models...)
		})
		logDBFatal(err)

		logger.Info("Hola, Database resettled successfully!")
	}
}
