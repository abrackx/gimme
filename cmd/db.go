package cmd

import (
	"fmt"
	"gimme/docker"
	"gimme/docker/postgres"
	"github.com/spf13/cobra"
)

const (
	POSTGRES = "postgres"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		docker.IsDockerRunning()
		for _, arg := range args {
			switch arg {
			case POSTGRES:
				postgres.Start()
			default:
				fmt.Printf("Looks like you didn't choose a valid database! Current supported databases are: %s\n", POSTGRES)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
