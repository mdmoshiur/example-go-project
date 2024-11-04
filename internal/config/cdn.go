package config

import (
	"time"

	"github.com/spf13/viper"
)

// CDNCfg stores the cdn service configurations.
type CDNCfg struct {
	Host      string
	Token     string
	Directory string
	Timeout   time.Duration
}

var cdnCfg CDNCfg

// CDN returns the cdn service configurations.
func CDN() CDNCfg {
	return cdnCfg
}

// loadCDN loads the cdn service configurations.
func loadCDN() {
	cdnCfg = CDNCfg{
		Host:      viper.GetString("cdn.host"),
		Token:     viper.GetString("cdn.token"),
		Directory: viper.GetString("cdn.directory"),
		Timeout:   viper.GetDuration("cdn.timeout") * time.Second,
	}
}
