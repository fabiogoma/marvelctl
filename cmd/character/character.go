/*
Copyright © 2025 Fabio Gonçalves Martins <fabiogoma@gmail.com>
*/
package character

import (
	"fmt"

	"github.com/fabiogoma/marvelctl/cmd"
	"github.com/spf13/cobra"
)

// CharacterCmd represents the character command
var CharacterCmd = &cobra.Command{
	Use:                   "character",
	DisableFlagsInUseLine: true,
	Short:                 "Explore Marvel characters and their stories",
	Long: `The character command groups subcommands for exploring Marvel characters using
the official Marvel Developers API.

Use it to look up information about a specific character, discover their appearances
in comics, series, and events, or simply browse your favorite characters from
the Marvel universe.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Printf("Error printing help: %v\n", err)
		}
	},
}

func init() {
	cmd.RegisterSubCommand(CharacterCmd)
}
