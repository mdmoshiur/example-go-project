package config

import (
	"log"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// Init load configurations from specified config file.
func Init(configFile string) error {
	viper.SetEnvPrefix("example_go")
	if err := viper.BindEnv("env"); err != nil {
		return err
	}

	env := viper.GetString("env")
	if env == "" {
		log.Fatal("environment variable 'env' is missing")
	}

	if env == EnvDevelopment { // local config file read
		viper.SetConfigName(configFile) // name of config file (without extension)

		viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath(".")    // optionally look for config in the working directory

		if err := viper.ReadInConfig(); err != nil { // Find and read the config file
			log.Fatalf("Error reading config file, %s", err.Error())
		}
	} else { // remote config file read
		if err := viper.BindEnv("consul_url"); err != nil {
			return err
		}
		if err := viper.BindEnv("consul_path"); err != nil {
			return err
		}

		consulURL := viper.GetString("consul_url")
		if consulURL == "" {
			log.Fatal("CONSUL_URL missing")
		}

		consulPath := viper.GetString("consul_path")
		if consulPath == "" {
			log.Fatal("CONSUL_PATH missing")
		}

		_ = viper.AddRemoteProvider("consul", consulURL, consulPath)
		viper.SetConfigType("yaml")

		err := viper.ReadRemoteConfig()
		if err != nil {
			log.Fatalf("%s named \"%s\"", err.Error(), consulPath)
		}
	}

	initConfig()
	return nil
}

// initConfig loads all the configurations
func initConfig() {
	loadApp()
	loadDB()
	loadRedis()
	loadJWT()
	loadCDN()
}
