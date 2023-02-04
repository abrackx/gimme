package cmd

import (
	"fmt"
	"gimme/docker"
	"github.com/docker/docker/api/types"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// dustbinCmd represents the dustbin command
var dustbinCmd = &cobra.Command{
	Use:   "dustbin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		categoryPromptContent := promptContent{
			"Something bad happened.",
			"Which container do you want to delete?",
		}
		docker.DeleteContainer(promptGetSelect(categoryPromptContent))
	},
}

func init() {
	rootCmd.AddCommand(dustbinCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dustbinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dustbinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type promptContent struct {
	errorMsg string
	label    string
}

func promptGetSelect(pc promptContent) types.Container {
	data := map[string]types.Container{}
	for _, container := range docker.GetContainers() {
		data[strings.Join(container.Names, ", ")] = container
	}
	containerNames := make([]string, 0, len(data))
	for key := range data {
		containerNames = append(containerNames, key)
	}

	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label: pc.label,
			Items: containerNames,
		}

		index, result, err = prompt.Run()
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return data[result]
}
