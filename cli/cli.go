package cli

import (
	"fmt"
	"net/http"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

const (
	defaultName = "create-go-app"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var a app

		a.hc = *http.DefaultClient

		cmd.Flags().StringVarP(&a.name, "name", "n", "", "base project directory eg. github.com/spf13/")
		cmd.Flags().StringVarP(&a.template, "template", "t", "", "base project directory eg. github.com/spf13/")

		if a.name == defaultName {
			prompt := promptui.Prompt{
				Label: "Project Name (leave blank for create-go-app)",
			}

			name, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			if name != "" {
				a.name = name
			}
		}
		if a.template == "" {
			var res []BranchRes
			err := fetchJSON(GithubAPIHost+TemplateRepoPath+BranchesEndpoint, a.hc, &res)
			if err != nil {
				fmt.Printf("Failed to fetch branches: %v\n", err)
				return
			}
			branchNames := lo.Map[BranchRes, string](res, func(b BranchRes, _ int) string {
				return b.Name
			})

			prompt := promptui.Select{
				Label: "Select Project Template",
				Items: branchNames,
			}
			_, tmpl, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			a.template = tmpl
		}
		}
	},
}

type app struct {
	hc       http.Client
	name     string
	template string // references a branch name in https://github.com/cameronbrill/go-project-template
}

func Run() int {

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
