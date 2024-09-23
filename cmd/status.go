/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
    tracked bool
    modified bool
    untracked bool
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the git status",
	Run: func(cmd *cobra.Command, args []string) {
        handleStatus()
	},
}

func init() {
    statusCmd.Flags().BoolVarP(&tracked, "tracked", "t", false, "Show all tracked files") 
    statusCmd.Flags().BoolVarP(&modified, "modified", "m", false, "Show all modified files") 
    statusCmd.Flags().BoolVarP(&untracked, "untracked", "u", false, "Show all untracked files") 
	rootCmd.AddCommand(statusCmd)
}

func handleStatus() {
    if tracked {
        out, err := exec.Command("git", "ls-files").Output()
        if err != nil {
            fmt.Println("Error: ", err)
            return
        }
        fmt.Println("Tracked files:")
        fmt.Print(string(out))
    } else if modified {
        out, err := exec.Command("git", "ls-files", "-m").Output()
        if err != nil {
            fmt.Println("Error: ", err)
            return
        }
        fmt.Println("Modified files:")
        fmt.Print(string(out))
    } else if untracked {
        out, err := exec.Command("git", "ls-files", "-o").Output()
        if err != nil {
            fmt.Println("Error: ", err)
            return
        }
        fmt.Println("Untracked files:")
        fmt.Print(string(out))
    } else {
        getAll()
    }
}

func getAll() {
    out, err := exec.Command("git", "status", "--porcelain").Output()
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }

    lines := strings.Split(string(out), "\n")
    for _, line := range lines {
        if len(line) == 0 {
            continue
        }
        status := line[:2]
        file := line[3:]

        switch status {
        case " M":
            color.Yellow("Modified: %s", file)
        case "A ":
            color.Green("Added: %s", file)
        case " D":
            color.Red("Deleted: %s", file)
        case "??":
            color.White("Untracked: %s", file)
        default:
            fmt.Println(line)
        }
    }
}
