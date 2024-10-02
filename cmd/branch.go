/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"git-helper/ui/switchBranch"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var availableSubCommands = []string{
    "list",
    "create",
    "switch",
    "delete",
}

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Manage Git branches",
    ValidArgs: availableSubCommands,
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Use a subcommand: %s\n", strings.Join(availableSubCommands, ", "))

	},
}

var listBranchCmd = &cobra.Command{
    Use: "list",
    Short: "List all branches",
    Run: func(cmd *cobra.Command, args []string) {
        handleListBranches()
    },
}

var createBranchCmd = &cobra.Command{
    Use: "create [branch_name]",
    Short: "Creata a new branch",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        branchName := args[0]
        handleCreateCommand(branchName)
    },
}

var switchBranchCmd = &cobra.Command{
    Use: "switch",
    Short: "Switch to another branch",
    Run: func(cmd *cobra.Command, args []string) {
        handleSwitchCommand()
    },
}

var deleteBranchCmd = &cobra.Command{
    Use: "delete",
    Short: "Delete a branch",
    Run: func(cmd *cobra.Command, args []string) {
        handleDeleteCommand()
    },
}

func init() {
    branchCmd.AddCommand(listBranchCmd)
    branchCmd.AddCommand(createBranchCmd)
    branchCmd.AddCommand(switchBranchCmd)
    branchCmd.AddCommand(deleteBranchCmd)
	rootCmd.AddCommand(branchCmd)
}

func handleListBranches() {
    out, err := exec.Command("git", "branch").Output()
    if err != nil {
        fmt.Println("Error listing brancehs:", err)
        return
    }
    fmt.Print(string(out))
}

func handleCreateCommand(branchName string) {
    err := exec.Command("git", "checkout", "-b", branchName).Run()
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    green := color.New(color.FgGreen).SprintFunc()
    fmt.Printf("Switched to a new branch '%s'\n", green(branchName))
}

func handleSwitchCommand() {
    option := switchBranch.Option{
        Title: "Switch Branch",
        Mode: "switch",
    }
    p := tea.NewProgram(switchBranch.InitialModel(option), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}

func handleDeleteCommand() {
    option := switchBranch.Option{
        Title: "Delete Branch",
        Mode: "delete",
    }
    p := tea.NewProgram(switchBranch.InitialModel(option), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
