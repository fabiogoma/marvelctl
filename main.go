/*
Copyright © 2025 Fabio Gonçalves Martins <fabiogoma@gmail.com>
*/
package main

import (
	"github.com/fabiogoma/marvelctl/cmd"
	_ "github.com/fabiogoma/marvelctl/cmd/character" // This triggers the init() in hero
	_ "github.com/fabiogoma/marvelctl/cmd/config"    // This triggers the init() in config
)

func main() {
	cmd.Execute()
}
