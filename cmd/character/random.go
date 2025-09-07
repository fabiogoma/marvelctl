/*
Copyright © 2025 Fabio Gonçalves Martins <fabiogoma@gmail.com>
*/
package character

import (
	"fmt"
	"math/rand"

	"github.com/fabiogoma/marvelctl/internal/marvel"
	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:                   "random",
	Args:                  cobra.ExactArgs(0),
	DisableFlagsInUseLine: true,
	Short:                 "Get a random Marvel character.",
	Long: `Fetches a random character from the Marvel universe using the Marvel API.

This is useful when you’re looking for inspiration, exploring unfamiliar
characters, or just want to discover something new without searching by name.

Each time you run the command, a different character may be returned.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := marvel.NewClient()
		if err != nil {
			fmt.Printf("Error creating Marvel client: %v\n", err)
			return
		}

		randomOffset, err := marvel.GetRandomCharactersOffset(client)
		if err != nil {
			fmt.Printf("Error fetching random offset: %v\n", err)
			return
		}

		numberOfCharacters := len(randomOffset.Data.Results)
		randomCharacter := randomOffset.Data.Results[rand.Intn(numberOfCharacters)]

		fmt.Println("Name:", randomCharacter.Name)
		fmt.Println("Description:", randomCharacter.Description)
		fmt.Println("Comics:", randomCharacter.Comics.Available)
		fmt.Println("Series:", randomCharacter.Series.Available)
		fmt.Println("Stories:", randomCharacter.Stories.Available)
	},
}

func init() {
	CharacterCmd.AddCommand(randomCmd)
}
