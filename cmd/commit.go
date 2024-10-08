/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	listfiles "git-helper/ui/commit/list-files"
	"git-helper/utils"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	all     bool
	file    bool
	message string
	yes     bool
)

type Options struct {
	files         *listfiles.Output
}

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit (-a | -f) -m <message> [-y]",
	Short: "Commit to changes to Git",
	Run: func(cmd *cobra.Command, args []string) {
		handleCommit()
	},
}

func init() {
	commitCmd.Flags().BoolVarP(&all, "all", "a", false, "Commit all files")
	commitCmd.Flags().BoolVarP(&file, "file", "f", false, "Choose file to commit")
	commitCmd.Flags().StringVarP(&message, "message", "m", "", "Commit message (required)")
	commitCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Push to remote")

	commitCmd.MarkFlagsOneRequired("all", "file")
	commitCmd.MarkFlagsRequiredTogether("all", "message")

	rootCmd.AddCommand(commitCmd)
}

func handleCommit() {
	option := Options{}

	if all {
		fmt.Println("Commiting all files...")
		exec.Command("git", "add", ".").Run()
	}

	if file {
		fOut, errOut := exec.Command("git", "ls-files", utils.GetGitDirectory(), "-o", "-m", "--exclude-standard").CombinedOutput()
		if errOut != nil {
			color.Red("Error: %v", errOut.Error())
			return
		}
		files := strings.Split(strings.TrimSpace(string(fOut)), "\n")
		option.files = &listfiles.Output{Output: files}

		tprogram := tea.NewProgram(listfiles.InitialListFiles(option.files, "Please select files to commit"), tea.WithAltScreen())
		if _, err := tprogram.Run(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

        if len(option.files.Selected) == 0 {
            fmt.Println("Please select file.")
            os.Exit(1)
        }

		color.Green("Adding file%s: %s\n", checkFiles(option.files.Selected), strings.Join(option.files.Selected, ", "))
		exec.Command("git", append([]string{"add"}, option.files.Selected...)...).Run()

        if option.files.Message == "" {
           color.Yellow("Restore file%s: %s\n", checkFiles(option.files.Selected), strings.Join(option.files.Selected, ", "))
           fmt.Println("Please add commit message.")
           exec.Command("git", append([]string{"restore", "--staged"}, option.files.Selected...)...).Run()
           os.Exit(1)
        }
	}

	var msg string
	if message != "" {
		msg = message
	} else {
		msg = option.files.Message
	}

	out, _ := exec.Command("git", "commit", "-m", msg).Output()
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
