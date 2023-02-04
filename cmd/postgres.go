package cmd

import (
	"fmt"
	"gimme/docker"
	"gimme/docker/postgres"
	"gimme/template"
	"github.com/spf13/cobra"
	"os"
)

// dbCmd represents the db command
var postgresCmd = &cobra.Command{
	Use:     "postgres",
	Aliases: []string{"postgresql"},
	Short:   "Start a new postgres container",
	Long: `Start a new postgres container on a random port.

Usage:
	gimme db postgres
`,
	Run: func(cmd *cobra.Command, args []string) {
		if !docker.IsDockerRunning() {
			fmt.Println("Docker is not running! Start docker then rerun this command.")
			os.Exit(1)
		}
		//TODO: Output to different project types ex: --project-type spring --format yml
		container := postgres.Start()
		template.PrintSpringTemplate(container.Database)
	},
}

func init() {
	dbCmd.AddCommand(postgresCmd)
}
