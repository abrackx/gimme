package cmd

import (
	"gimme/docker"
	"gimme/template"
	"github.com/AlecAivazis/survey/v2"
	"github.com/docker/docker/api/types"
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
		container := selectContainerPrompt()
		docker.DeleteContainer(container)
	},
}

func init() {
	rootCmd.AddCommand(dustbinCmd)
}

func selectContainerPrompt() types.Container {
	data := map[string]types.Container{}
	for _, container := range docker.GetContainers() {
		data[strings.Join(container.Names, ", ")] = container
	}
	containerNames := make([]string, 0, len(data))
	for key := range data {
		containerNames = append(containerNames, key)
	}

	if len(containerNames) == 0 {
		println("No containers to remove")
		os.Exit(0)
	}

	survey.SelectQuestionTemplate = template.SurveyTemplate()

	var result string
	prompt := &survey.Select{
		Message: "Which container do you want to delete?",
		Options: containerNames,
	}
	err := survey.AskOne(prompt, &result)
	if err != nil {
		panic(err)
	}

	return data[result]
}
