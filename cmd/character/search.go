/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package character

import (
	"fmt"

	"github.com/fabiogoma/marvelctl/internal/marvel"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:                   "search [character name]",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Short:                 "Search for Marvel characters by name or keyword.",
	Long: `Use the search subcommand to find Marvel characters by name or partial match.
You can provide a full character name (e.g., "Spider-Man") or just part of the
name (e.g., "Spider") to retrieve all matching characters from the Marvel API.

This is useful when you want to explore characters without knowing their
exact names, or when browsing for related results. The command supports
filtering, pagination, and other options to refine your search.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := marvel.NewClient()
		if err != nil {
			fmt.Printf("Error creating Marvel client: %s\n", err)
			return
		}

		characterName := args[0]
		characterInfo, err := marvel.SearchCharacterByName(characterName, client)
		if err != nil {
			fmt.Printf("Error fetching character info: %s\n", err)
			return
		}

		if len(characterInfo.Data.Results) > 0 {
			for _, character := range characterInfo.Data.Results {
				fmt.Println("Name:", character.Name)
				fmt.Println("ID:", character.ID)
				fmt.Println("")
			}
		} else {
			fmt.Println("Character not found.")
		}
	},
}

func init() {
	CharacterCmd.AddCommand(searchCmd)
}
