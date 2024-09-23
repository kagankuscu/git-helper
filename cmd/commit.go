/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
    all bool
    files []string
    message string
)
// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit to changes to Git",
	Run: func(cmd *cobra.Command, args []string) {
        handleCommit()
	},
}

func init() {
    commitCmd.Flags().BoolVarP(&all, "all", "a", false, "Commit all files")
    commitCmd.Flags().StringSliceVarP(&files, "file", "f", []string{}, "File to commit")
    commitCmd.Flags().StringVarP(&message, "message", "m", "", "Commit message")

	rootCmd.AddCommand(commitCmd)
}

func handleCommit() {
    if message == "" {
        fmt.Println("Commit message is required. Use --message or -m to specify one.")
        os.Exit(1)
    }

    if all {
        fmt.Println("Commiting all files...")
        exec.Command("git", "add", ".").Run()
    } else if len(files) > 0  {
        fmt.Printf("Adding file%s: %s\n", checkFiles(files), strings.Join(files, ", "))
        exec.Command("git", append([]string{"add"}, files...)...).Run()
    } else {
        fmt.Println("You must either specify --all or --file to add files.")
        os.Exit(1)
    }

    out, _ := exec.Command("git", "commit", "-m", message).Output()
    fmt.Print(string(out))
}

func checkFiles(files []string) string {
    if len(files) > 1 {
        return "s"
    }
    return ""
}
