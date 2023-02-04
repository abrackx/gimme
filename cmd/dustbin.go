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

var dustbinCmd = &cobra.Command{
	Use:   "dustbin",
	Short: "Clean up containers",
	Long: `Select which containers you would like to delete.

Usage:
	gimme dustbin`,
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
