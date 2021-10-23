package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thaibui2308/ai-hackathon/api"
	"github.com/thaibui2308/ai-hackathon/config"
	"github.com/thaibui2308/ai-hackathon/tui"
)

var SearchRepo = &cobra.Command{
	Use:   "f",
	Short: "Search for a specific pull request from a repository.",
	Long:  "grp f <YAML_CONFIG_FILE>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Not enough arguments to process the command")
			os.Exit(1)
		}
		Configurations, err := config.GetConfiguration(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		response, err := api.GetPullRequest(Configurations.Username, Configurations.Repository, Configurations.CommitID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		commit, err := api.GetStats(response[0].URL)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		author := response[0]

		tui.RenderUserInfo(author.Author.URL, commit.Stats, commit.Commit.Message, commit.Files, author.Commit.Verification)

	},
}

var MergePullRequest = &cobra.Command{
	Use:   "m",
	Short: "Merge pull request",
	Long:  "grp m <YAML_CONFIGFILE>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Not enough arguments to process the command")
			os.Exit(1)
		}
		Configurations, err := config.GetConfiguration(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		url := "https://api.github.com/repos/" + Configurations.Username + "/" + Configurations.Repository + "/pulls/" + Configurations.CommitID + "/merge"
		api.MergePullRequest(url)
	},
}

// launch all commands
func AddCommands() {
	RootCmd.AddCommand(SearchRepo)
	RootCmd.AddCommand(MergePullRequest)
}
