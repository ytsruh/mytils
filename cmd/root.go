package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my",
	Short: "A collection of CLI utilities packaged together to make your life easier & customise your computer.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n" + "Please choose a sub command to run\n" + "\n")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here, will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.my.yaml)")

	// Cobra also supports local flags, which will only run when this action is called directly.
	pomoCmd.Flags().DurationP("time", "t", 25*time.Minute, "Pomodoro timer duration")

	// Add subcommands
	rootCmd.AddCommand(licenseCmd)
	rootCmd.AddCommand(tmuxCmd)
	rootCmd.AddCommand(pomoCmd)
}
