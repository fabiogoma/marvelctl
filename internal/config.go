package internal

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func GetConfigByKey(key string) (string, error) {
	validKeys := map[string]bool{
		"public_key":  true,
		"private_key": true,
	}

	if !validKeys[key] {
		return "", fmt.Errorf("invalid key: %s", key)
	}

	// Set up viper to use a config file
	// e.g: $HOME/.marvelctl.yaml

	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error finding home folder: %s", err)
	}

	viper.SetConfigName(".marvelctl.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(homePath)

	err = viper.ReadInConfig()
	if err != nil {
		return "", fmt.Errorf("error loading configuration file: %s", err)
	}

	return viper.GetString(key), nil
}
