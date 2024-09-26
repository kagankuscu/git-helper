/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
    all bool
    files []string
    message string
    yes bool
)
// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit (-a | -f <file_name>) -m <message> [-y]",
	Short: "Commit to changes to Git",
	Run: func(cmd *cobra.Command, args []string) {
        handleCommit()
	},
}

func init() {
    commitCmd.Flags().BoolVarP(&all, "all", "a", false, "Commit all files")
    commitCmd.Flags().StringSliceVarP(&files, "file", "f", []string{}, "File to commit")
    commitCmd.Flags().StringVarP(&message, "message", "m", "", "Commit message (required)")
    commitCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Push to remote")

    commitCmd.MarkFlagsOneRequired("all", "file")
    commitCmd.MarkFlagRequired("message")

	rootCmd.AddCommand(commitCmd)
}

func handleCommit() {
    if all {
        fmt.Println("Commiting all files...")
        exec.Command("git", "add", ".").Run()
    } else if len(files) > 0  {
        fmt.Printf("Adding file%s: %s\n", checkFiles(files), strings.Join(files, ", "))
        exec.Command("git", append([]string{"add"}, files...)...).Run()
    }

    out, _ := exec.Command("git", "commit", "-m", message).Output()
    fmt.Print(string(out))

    if yes {
        handleSync()
    }
}

func checkFiles(files []string) string {
    if len(files) > 1 {
        return "s"
    }
    return ""
}
