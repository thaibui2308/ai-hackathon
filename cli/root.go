package cli

import (
	"github.com/spf13/cobra"
)

// RootCmd main cobra command
var RootCmd = &cobra.Command{
	Use:   "gpr-cli",
	Short: "A cli to check for any pull requests from your terminal.",
}
