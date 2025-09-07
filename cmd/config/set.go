/*
Copyright © 2025 Fabio Gonçalves Martins <fabiogoma@gmail.com>
*/
package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:                   "set <key> <value>",
	Args:                  cobra.ExactArgs(2),
	DisableFlagsInUseLine: true,
	Short:                 "Set a configuration parameter (e.g., public_key or private_key",
	Long: `The 'set' command allows you to update or add configuration parameters to your marvelctl config file.

Use this command to securely store your Marvel API credentials, such as 'public_key' and 'private_key'.

Example usage:

  marvelctl config set public_key YOUR_PUBLIC_KEY
  marvelctl config set private_key YOUR_PRIVATE_KEY

The configuration is saved in $HOME/.marvelctl.yaml and is used by other marvelctl commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]

		validKeys := map[string]bool{
			"public_key":  true,
			"private_key": true,
		}

		if !validKeys[key] {
			fmt.Printf("Invalid key: %s\n", key)
			return
		}
		// Set up viper to use a config file
		// e.g: $HOME/.marvelctl.yaml

		homePath, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error finding home folder: %s\n", err)
		}
		configFile := fmt.Sprintf("%s/.marvelctl.yaml", homePath)

		viper.SetConfigName(".marvelctl.yaml")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(homePath)

		// Read existing config (if any)
		err = viper.ReadInConfig()
		if err != nil {
			// If the config file doesn't exist, create it.
			viper.Set(key, value)
			err = viper.SafeWriteConfigAs(configFile)
			if err != nil {
				fmt.Printf("Error creating config: %v\n", err)
				return
			}
		} else {
			// If the config file already exists, update it.
			viper.Set(key, value)
			err = viper.WriteConfig()
			if err != nil {
				fmt.Printf("Error updating config: %v\n", err)
				return
			}
		}

		fmt.Printf("Set %s = %s\n", key, value)
	},
}

func init() {
	ConfigCmd.AddCommand(setCmd)
}
