/*
Copyright © 2025 Fabio Gonçalves Martins <fabiogoma@gmail.com>
*/
package character

import (
	"fmt"

	"github.com/fabiogoma/marvelctl/internal/marvel"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info [character name]",
	Args:  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Short: "Get detailed information about a Marvel character by name",
	Long: `The "character info" command retrieves detailed information about a Marvel character
using the official Marvel Developers API.

Provide the character's name as an argument to fetch data such as their
description, number of available comics, series appearances, and events.
This command requires your Marvel API keys to be configured with
"marvelctl config set".`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := marvel.NewClient()
		if err != nil {
			fmt.Printf("Error creating Marvel client: %s\n", err)
			return
		}

		characterName := args[0]
		characterInfo, err := marvel.GetCharacterByName(characterName, client)
		if err != nil {
			fmt.Printf("Error fetching character info: %s\n", err)
			return
		}

		if len(characterInfo.Data.Results) > 0 {
			character := characterInfo.Data.Results[0]
			fmt.Println("Name:", character.Name)
			fmt.Println("Description:", character.Description)
			fmt.Println("Comics:", character.Comics.Available)
			fmt.Println("Series:", character.Series.Available)
			fmt.Println("Stories:", character.Stories.Available)
		} else {
			fmt.Println("Character not found.")
		}
	},
}

func init() {
	CharacterCmd.AddCommand(infoCmd)
}
