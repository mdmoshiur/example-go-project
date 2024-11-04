package migration

import (
	"fmt"

	"github.com/mdmoshiur/example-go/internal/conn"
	"github.com/mdmoshiur/example-go/internal/logger"
	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "migration",
		Short: "Run database migrations",
		Long:  `Migration is a tool to generate and modify database tables`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := conn.ConnectDB(); err != nil {
				return fmt.Errorf("can't connect the database: %w", err)
			}
			return nil
		},
	}
)

// logDBFatal logs db fatal errors.
func logDBFatal(err error) {
	if err != nil {
		logger.Fatal(err)
	}
}
