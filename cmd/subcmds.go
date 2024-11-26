package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"ytsruh.com/my/pkg/license"
	"ytsruh.com/my/pkg/password"
	"ytsruh.com/my/pkg/pomo"
	"ytsruh.com/my/pkg/tmux"
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
		tmux.Run()
	},
}

var pomoCmd = &cobra.Command{
	Use:     "pomo",
	Aliases: []string{"pomodoro", "timer"},
	Short:   "A pomodoro timer",
	Long:    `A pomodoro timer to help you stay focused on your work.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Can ignore error as default value of 25 minutes is set
		timeout, _ := cmd.Flags().GetDuration("time")
		pomo.Run(timeout)
	},
}

var pwdCmd = &cobra.Command{
	Use:     "password",
	Aliases: []string{"pass", "pwd"},
	Short:   "Generate a random password",
	Long:    `Generate a random password.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Can ignore errors as default values of 16 characters and false are set
		enc, _ := cmd.Flags().GetBool("encoded")
		length, _ := cmd.Flags().GetInt("length")
		// Run the password generator
		pwd := password.Run(length, enc)
		fmt.Println("\n" + pwd + "\n")
	},
}
