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
    oneline bool
    commits bool
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Manage of Git log",
	Run: func(cmd *cobra.Command, args []string) {
        handleLog()
	},
}

func init() {
	logCmd.Flags().BoolVarP(&oneline, "oneline", "o", false, "Show single line of log")
	logCmd.Flags().BoolVarP(&commits, "commits", "c", false, "Show commits that have not been pushed yet")
	rootCmd.AddCommand(logCmd)
}

func handleLog() {
    if oneline {
        out, err := exec.Command("git", "log", "--pretty=format:%h-%an-%ar-%s").Output()
        if err != nil {
            fmt.Println("Error: ", err)
            return
        }
        yellow := color.New(color.FgYellow).SprintFunc()
        green := color.New(color.FgGreen).SprintfFunc()
        cyan := color.New(color.FgCyan).SprintfFunc()
        lines := strings.Split(string(out), "\n")

        for _, line := range lines {
            splLine := strings.Split(string(line), "-")
            fmt.Printf("%s - %s, %s : %s\n", yellow(splLine[0]), green(splLine[1]), cyan(splLine[2]), splLine[3])
        }
    } else if commits {
        out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
        if err != nil {
            fmt.Println("Error reading current branch: ", err)
            return
        }

        branchName := strings.TrimSpace(string(out))

        logOut, _ := exec.Command("git", "log", fmt.Sprintf("%s..HEAD", branchName)).CombinedOutput()
        fmt.Print(string(logOut))
    } else {
        out, err := exec.Command("git", "log").Output()
        if err != nil {
            fmt.Println("Error: ", err)
            return
        }

        lines := strings.Split(string(out), "\n")
        for i, line := range lines {
            if i == 0 || i % 6 == 0{
                color.Yellow(line)
                continue
            }
            fmt.Println(line)
        }
    }
}
