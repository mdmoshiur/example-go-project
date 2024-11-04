package migration

import (
	"github.com/mdmoshiur/example-go/internal/conn"
	"github.com/mdmoshiur/example-go/internal/logger"
	"github.com/mdmoshiur/example-go/internal/migration"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var (
	upCmd = &cobra.Command{
		Use:   "up",
		Short: "Populate tables in database",
		Long:  `Populate tables in database`,
		Run:   upDatabase,
	}
)

func init() {
	RootCmd.AddCommand(upCmd)
}

func upDatabase(_ *cobra.Command, _ []string) {
	confirmed := askForConfirmation("Migration up")
	if confirmed {
		logger.Info("Populating database...")
		db := conn.DefaultDB()

		err := db.Transaction(func(tx *gorm.DB) error {
			return tx.AutoMigrate(migration.Models...)
		})
		logDBFatal(err)

		logger.Info("Hola, Database populated successfully!")
	}
}
