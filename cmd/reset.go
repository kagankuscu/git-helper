/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
    discard bool
    undo bool
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Manage to Git reset",
	Run: func(cmd *cobra.Command, args []string) {
        handleReset()
	},
}

func init() {
	resetCmd.Flags().BoolVarP(&discard, "discard", "d", false, "Discard all changes")
	resetCmd.Flags().BoolVarP(&undo, "undo", "u", false, "Undo last commit")
	rootCmd.AddCommand(resetCmd)
}

func handleReset() {
   if discard {
       fmt.Println("Discard all files")
       exec.Command("git", "reset", "--hard", "HEAD").Run()
   } else if undo {
       fmt.Println("Undo last commit")
       exec.Command("git", "reset", "--soft", "HEAD~1").Run()
   } else {
       fmt.Println("You must either specify --discard or --undo to add files.")
   }
}
