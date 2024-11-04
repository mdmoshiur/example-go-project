package config

import (
	"time"

	"github.com/spf13/viper"
)

// Version contains application version
var Version = "1.0.0"

// represents environment level

const (
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

// Application represents application configurations
type Application struct {
	HTTPPort int    `json:"http_port"`
	Verbose  bool   `json:"debug"`
	Env      string `json:"env"`

	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	IdleTimeout  time.Duration `json:"idle_timeout"`
	HTTPTimeout  time.Duration `json:"http_timeout"`

	PaginationPageSize int `json:"pagination_page_size"`
}

var app Application

// App contains app configurations
func App() Application {
	return app
}

// loadApp loads the application configurations.
func loadApp() {
	app = Application{
		HTTPPort: viper.GetInt("app.http_port"),
		Verbose:  viper.GetBool("app.verbose"),
		Env:      viper.GetString("app.env"),

		ReadTimeout:  viper.GetDuration("app.read_timeout") * time.Second,
		WriteTimeout: viper.GetDuration("app.write_timeout") * time.Second,
		IdleTimeout:  viper.GetDuration("app.idle_timeout") * time.Second,
		HTTPTimeout:  viper.GetDuration("app.http_timeout") * time.Second,

		PaginationPageSize: viper.GetInt("app.pagination_page_size"),
	}
}
