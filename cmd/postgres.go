package cmd

import (
	"fmt"
	"gimme/docker"
	"gimme/docker/postgres"
	"gimme/template"
	"github.com/spf13/cobra"
	"os"
)

var projectType string
var format string

const SPRING_PROJECT_TYPE = "spring"
const YAML_FORMAT = "yaml"
const YML_FORMAT = "yml"

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
		container := postgres.Start()
		//TODO: Rework this, should probably validate before starting the container.
		switch projectType {
		case SPRING_PROJECT_TYPE:
			switch format {
			case YAML_FORMAT, YML_FORMAT:
				template.PrintSpringTemplate(container.Database)
				break
			default:
				fmt.Println("Format not supported")
			}
		default:
			fmt.Println("Project type not supported")
		}
	},
}

func init() {
	dbCmd.AddCommand(postgresCmd)
	postgresCmd.Flags().StringVarP(&projectType, "project-type", "p", "spring", "Connection project type to generate")
	postgresCmd.Flags().StringVarP(&format, "format", "f", "yml", "Format of the connection to generate")
}
