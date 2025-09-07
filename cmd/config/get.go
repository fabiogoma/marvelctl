/*
Copyright © 2025 Fabio Gonçalves Martins <fabiogoma@gmail.com>
*/
package config

import (
	"fmt"

	"github.com/fabiogoma/marvelctl/internal"
	"github.com/spf13/cobra"
)

// getCmd represents the show command
var getCmd = &cobra.Command{
	Use:                   "get <key>",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Short:                 "Retrieve a configuration value",
	Long: `The "config get" command reads a value from your marvelctl configuration file.
Provide the key name as an argument (for example, "public_key" or "private_key")
to display the stored value.

If no value is set, a helpful message will be shown. Sensitive keys such as
"private_key" may be partially masked when displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value, err := internal.GetConfigByKey(key)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s = %s\n", key, value)
	},
}

func init() {
	ConfigCmd.AddCommand(getCmd)

}
