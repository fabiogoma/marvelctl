/*
Copyright © 2025 Fabio Gonçalves Martins <fabiogoma@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// comicsCmd represents the comics command
var comicsCmd = &cobra.Command{
	Use:   "comics",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("comics called")
	},
}

func init() {
	rootCmd.AddCommand(comicsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// comicsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// comicsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
