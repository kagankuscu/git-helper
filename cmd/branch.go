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
    Use: "switch [branch_name]",
    Short: "Switch to another branch",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        branchName := args[0]
        handleSwitchCommand(branchName)
    },
}

var deleteBranchCmd = &cobra.Command{
    Use: "delete [branch_name]",
    Short: "Delete a branch",
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        branchName := args[0]
        handleDeleteCommand(branchName)
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

func handleSwitchCommand(branchName string) {
    err := exec.Command("git", "checkout", branchName).Run()
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    yellow := color.New(color.FgYellow).SprintFunc()
    fmt.Printf("Switched to branch '%s'\n", yellow(branchName))
}

func handleDeleteCommand(branchName string) {
    out, err := exec.Command("git", "branch", "-D", branchName).CombinedOutput()
    if err != nil {
        fmt.Print(string(out))
        return
    }
    red := color.New(color.FgRed).SprintFunc()
    fmt.Printf("Deleted branch '%s'\n", red(branchName))
}
