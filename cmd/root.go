/*
Copyright © 2025 Fabio Gonçalves Martins <fabiogoma@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "marvelctl",
	Short: "A CLI tool to explore the Marvel universe",
	Long: `marvelctl is a command-line tool that connects to the official Marvel Developers API
to fetch information about your favorite characters, comics, and more.

Use marvelctl to quickly look up heroes, discover their stories, and explore
the Marvel universe without leaving your terminal.

Before using the tool, make sure to configure your Marvel API public and private
keys with 'marvelctl config set'.`,
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
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.marvelctl.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func RegisterSubCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
