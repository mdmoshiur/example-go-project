package config

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"

	_ "github.com/spf13/viper/remote"
)

// Init load configurations from config.yaml file
func Init(cfgFile string) error {
	spew.Dump(cfgFile)

	// local config file read
	viper.SetConfigName("config") // name of config file (without extension)
	if cfgFile != "" {
		viper.SetConfigName(cfgFile)
	}

	viper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")                     // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil { // Find and read the config file
		return err
	}

	// remote config file read

	//viper.SetEnvPrefix("candy_recruiter")
	//if err := viper.BindEnv("env"); err != nil {
	//	return err
	//}
	//
	//if err := viper.BindEnv("consul_url"); err != nil {
	//	return err
	//}
	//if err := viper.BindEnv("consul_path"); err != nil {
	//	return err
	//}
	//
	//consulURL := viper.GetString("consul_url")
	//consulPath := viper.GetString("consul_path")
	//if consulURL == "" {
	//	return errors.New("CONSUL_URL is missing")
	//}
	//if consulPath == "" {
	//	return errors.New("CONSUL_PATH is missing")
	//}
	//
	//if err := viper.AddRemoteProvider("consul", consulURL, consulPath); err != nil {
	//	return err
	//}
	//
	//viper.SetConfigType("yaml")
	//
	//if err := viper.ReadRemoteConfig(); err != nil {
	//	return fmt.Errorf(`%s named "%s"`, err.Error(), consulPath)
	//}

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
