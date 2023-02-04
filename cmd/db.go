package cmd

import (
	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Start a new database container",
	Long: `Start a new database container on a random port.

Usage:
	gimme db postgres
`,
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
