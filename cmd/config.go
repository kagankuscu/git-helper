/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"git-helper/config"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config [key] [value]",
	Short: "Configure Git Helper defaults",
	Long: `Set default remote, branch, and other Git Helper configurations.`,
    Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
        key := args[0]
        value := args[1]
        config.UpdateConfig(key, value)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
