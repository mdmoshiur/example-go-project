package cmd

import (
	"fmt"

	"github.com/mdmoshiur/example-go/internal/config"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version print the current build number and version information",
		Long:  `Version print the current build number and version information`,
		Run:   version,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func version(_ *cobra.Command, _ []string) {
	fmt.Printf("Version: %s\n", config.Version)
}
