/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize with remote",
	Run: func(cmd *cobra.Command, args []string) {
        handleSync()
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

func handleSync() {
    err := exec.Command("git", "pull").Run()
    if err != nil {
        color.Red("There is no tracking information for the current branch.")
        return
    }

    out, err1 := exec.Command("git", "push").CombinedOutput()
    if err1 != nil {
        color.Red(string(out))
        return
    }
    fmt.Print(string(out))
}
