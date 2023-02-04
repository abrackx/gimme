package cmd

import (
	"fmt"
	"gimme/docker"
	"gimme/docker/postgres"
	"gimme/template"
	"github.com/spf13/cobra"
)

const (
	POSTGRES = "postgres"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Start a new database container",
	Long: `Start a new database container on a random port.

Usage:
	gimme db postgres
`,
	Run: func(cmd *cobra.Command, args []string) {
		docker.IsDockerRunning()
		//TODO: Output to different project types ex: --project-type spring --format yml
		for _, arg := range args {
			switch arg {
			case POSTGRES:
				{
					container := postgres.Start()
					template.PrintSpringTemplate(container.Database)
				}
			default:
				fmt.Printf("Looks like you didn't choose a valid database! Current supported databases are: %s\n", POSTGRES)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
