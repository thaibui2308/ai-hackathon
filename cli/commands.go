package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thaibui2308/ai-hackathon/api"
	"github.com/thaibui2308/ai-hackathon/tui"
)

var SearchRepo = &cobra.Command{
	Use:   "f",
	Short: "Search for a specific pull requests of a repository.",
	Long:  "grp -s <USERNAME> <REPO_NAME> <PULL_REQUEST_ID>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("Not enough arguments to process")
			os.Exit(1)
		}
		username := args[0]
		repository := args[1]
		pId := args[2]

		response, err := api.GetPullRequest(username, repository, pId)
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
		tui.RenderUserInfo(author.Author.URL, commit.Stats, commit.Commit.Message, commit.Files)
		fmt.Println(commit.Stats)
	},
}

// launch all commands
func AddCommands() {
	RootCmd.AddCommand(SearchRepo)
}
