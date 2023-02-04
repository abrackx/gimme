package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gimme",
	Short: "Gimme a database, quick!",
	Long: `A simple CLI for quickly spinning up database containers.

Usage:
	gimme db postgres
Will start up a postgres container on a random port.
	gimme dustbin
Used to clean up containers.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
