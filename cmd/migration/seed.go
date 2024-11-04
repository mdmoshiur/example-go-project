package migration

import (
	"github.com/mdmoshiur/example-go/internal/conn"
	"github.com/mdmoshiur/example-go/internal/logger"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var (
	seedCmd = &cobra.Command{
		Use:   "seed",
		Short: "Seed tables in database",
		Long:  `Seed tables in database`,
		Run:   seedDatabase,
	}
)

func init() {
	RootCmd.AddCommand(seedCmd)
}

func seedDatabase(_ *cobra.Command, args []string) {
	confirmed := askForConfirmation("Migration seed")
	if confirmed {
		// no arguments return
		if len(args) == 0 {
			return
		}

		//hasAllArgument := helper.InSlice[string]("all", args)

		logger.Info("Seeding database...")
		db := conn.DefaultDB()

		// Disable foreign key constraints
		db.Exec("SET FOREIGN_KEY_CHECKS=0")

		err := db.Transaction(func(tx *gorm.DB) error {
			// users seeder if arguments has `users`
			//if hasAllArgument || helper.InSlice[string]("users", args) {
			//	if err := usersSeeder(tx); err != nil {
			//		return err
			//	}
			//}

			// return nil will commit the whole transaction
			return nil
		})

		logDBFatal(err)

		logger.Info("Hola, Database populated successfully!")
	}
}

//func usersSeeder(tx *gorm.DB) error {
//	var users []domain.User
//	var roleIDs []uint8
//	if err := tx.Model(&domain.Role{}).Pluck("id", &roleIDs).Error; err != nil {
//		return fmt.Errorf("error in users seeder role ids find: %w", err)
//	}
//
//	// password of all user is 12345
//	password, _ := usecase.GenPasswordHash("12345")
//
//	// generate unique values
//	faker.SetGenerateUniqueValues(true)
//
//	for i := 0; i < 50; i++ {
//		phone := faker.Phonenumber()
//		status := uint8(rand.Intn(3))
//		user := domain.User{
//			Name:     faker.Name(),
//			Email:    faker.Email(),
//			Phone:    &phone,
//			RoleID:   &roleIDs[rand.Intn(len(roleIDs))],
//			Status:   &status,
//			Password: password,
//		}
//		users = append(users, user)
//	}
//
//	if err := tx.Create(users).Error; err != nil {
//		return fmt.Errorf("error in users seeder: %w", err)
//	}
//
//	return nil
//}
