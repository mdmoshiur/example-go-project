package config

import (
	"time"

	"github.com/spf13/viper"
)

// JWTConfig stores the jwt configurations.
type JWTConfig struct {
	SecretKey                 []byte
	ExpirationDurationInHours time.Duration
}

var jwtCfg JWTConfig

// JWT returns the jwt configuration.
func JWT() JWTConfig {
	return jwtCfg
}

// loadDB loads the jwt configuration.
func loadJWT() {
	jwtCfg = JWTConfig{
		SecretKey:                 []byte(viper.GetString("jwt.secret_key")),
		ExpirationDurationInHours: viper.GetDuration("jwt.expiration_duration_in_hours") * time.Hour,
	}
}
