/*
Copyright © 2025 Fabio Gonçalves Martins <fabiogoma@gmail.com>
*/
package config

import (
	"fmt"

	"github.com/fabiogoma/marvelctl/cmd"
	"github.com/spf13/cobra"
)

// ConfigCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:                   "config",
	DisableFlagsInUseLine: true,
	Short:                 "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Printf("Error printing help: %v\n", err)
		}
	},
}

func init() {
	cmd.RegisterSubCommand(ConfigCmd)
}
