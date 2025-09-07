/*
Copyright © 2025 Fabio Gonçalves Martins <fabiogoma@gmail.com>
*/
package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:                   "reset",
	Args:                  cobra.ExactArgs(0),
	DisableFlagsInUseLine: true,
	Short:                 "Reset the configuration by removing the file",
	Long: `The "config reset" command clears your marvelctl configuration by removing
the file. This removes stored Marvel API keys and any other
custom settings, allowing you to start fresh.

Use this command with caution, as all saved values will be lost and you will
need to run "marvelctl config set" again to provide your API keys before using
other commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		homePath, err := os.UserHomeDir()

		if err != nil {
			fmt.Printf("Error finding home folder: %s\n", err)
		}
		configFile := fmt.Sprintf("%s/.marvelctl.yaml", homePath)
		err = os.Remove(configFile)
		if err != nil {
			fmt.Printf("Error removing configuration: %s\n", err)
			return
		}
		fmt.Println("Configuration file removed")
	},
}

func init() {
	ConfigCmd.AddCommand(resetCmd)
}
