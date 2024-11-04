package migration

import (
	"github.com/mdmoshiur/example-go/internal/conn"
	"github.com/mdmoshiur/example-go/internal/logger"
	"github.com/mdmoshiur/example-go/internal/migration"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Drop tables from database",
	Long:  `Drop tables from database`,
	Run:   downDatabase,
}

func init() {
	RootCmd.AddCommand(downCmd)
}

func downDatabase(_ *cobra.Command, _ []string) {
	confirmed := askForConfirmation("Migration down")
	if confirmed {
		logger.Info("Dropping database tables...")
		db := conn.DefaultDB()

		err := db.Transaction(func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(migration.Models...)
		})
		logDBFatal(err)

		logger.Info("Hola, Database tables dropped successfully!")
	}
}
