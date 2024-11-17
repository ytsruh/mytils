package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"ytsruh.com/my/license"
)

var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Create a license file",
	Long:  `All software has licenses. Generate a license file for your project by following a simple prompt.`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := license.RunPrompt()
		if err != nil {
			fmt.Printf("Generator failed %v\n", err)
			os.Exit(1)
		}

		err = license.GenTemplate(result)

		if err != nil {
			fmt.Printf("Template failed to generate %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Template has been generated & saved to your path. \n")
	},
}

var tmuxCmd = &cobra.Command{
	Use:   "tmux",
	Short: "A tmux cheat sheet",
	Long:  `A list of commands for Tmux in a cheat sheet format.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("A list of commands for Tmux in a cheat sheet format. \n")
	},
}
