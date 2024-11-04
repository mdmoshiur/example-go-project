package cmd

import (
	"os"

	"github.com/mdmoshiur/example-go/cmd/migration"
	"github.com/mdmoshiur/example-go/internal/config"
	"github.com/mdmoshiur/example-go/internal/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile                 string
	verbose, prettyPrintLog bool

	rootCmd = &cobra.Command{
		Use:   "example-go",
		Short: "example-go service ...",
		Long:  `example-go service ...`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	logger.DefaultLogger()
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file")
	rootCmd.PersistentFlags().BoolVarP(&prettyPrintLog, "pretty", "p", false, "pretty print verbose/logger")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	// set the value of viper config
	err := viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	if err != nil {
		return
	}

	// add migration rootCmd
	rootCmd.AddCommand(migration.RootCmd)
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}

func initConfig() {
	logger.Info("loading configurations...")
	if err := config.Init(cfgFile); err != nil {
		logger.Error("Failed to load configurations")
		logger.Fatal(err)
	}

	// logger as JSON instead of the default ASCII formatter
	logger.SetLogFormatter(&logrus.JSONFormatter{
		PrettyPrint: prettyPrintLog,
	})

	// by default logs only the warning severity or above
	// logger.SetLogLevel(logrus.WarnLevel)
	if config.App().Verbose { // if verbose set true show trace level
		logger.SetLogLevel(logrus.TraceLevel)
	}
	if verbose { // if -v flag pass override previous value
		logger.SetLogLevel(logrus.TraceLevel)
	}
}
