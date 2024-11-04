package conn

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mdmoshiur/example-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const slowQueriesThreshold = 500

// DB holds the database instance.
var db *gorm.DB

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             slowQueriesThreshold * time.Millisecond, // Slow SQL threshold
		LogLevel:                  logger.Info,                             // Log level
		IgnoreRecordNotFoundError: true,                                    // Ignore ErrRecordNotFound error for logger
		Colorful:                  false,                                   // Disable color
	},
)

// Ping tests if db connection is alive.
func Ping() error {
	d, _ := db.DB() // returns *sql.DB
	return d.Ping()
}

// Connect sets the db client of database using configuration cfg.
func Connect(cfg config.Database) error {
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)
	gormDB, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	})
	if err != nil {
		return err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}

	if cfg.MaxIdleConn != 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	}

	if cfg.MaxOpenConn != 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	}

	if cfg.MaxConnLifetime.Seconds() != 0 {
		sqlDB.SetConnMaxLifetime(cfg.MaxConnLifetime)
	}

	db = gormDB

	return nil
}

// DefaultDB returns default db of the application.
func DefaultDB() *gorm.DB {
	if Ping() != nil { // if connection lost, reconnect it
		_ = ConnectDB()
	}

	if config.App().Env == config.EnvDevelopment ||
		config.App().Env == config.EnvStaging {
		db.Debug()
	}

	return db
}

// ConnectDB sets the db client of database using default database configuration file.
func ConnectDB() error {
	cfg := config.DB()

	return Connect(cfg)
}
