/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"git-helper/utils"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var validCommands = []string {
    "add",
    "remove",
    "list",
}

var (
    rootDir = utils.GetGitDirectory()
    gitignoreFilename = ".gitignore"
    path = fmt.Sprintf("%s/%s", rootDir, gitignoreFilename)
)

// gitignoreCmd represents the gitignore command
var gitignoreCmd = &cobra.Command{
	Use:   "gitignore",
	Short: "Managed to gitignore file",
    Args: cobra.ExactArgs(1),
    ValidArgs: availableSubCommands,
}


var addSubCommand = &cobra.Command{
	Use:   "add <path>",
	Short: "Add to gitignore file",
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        handleAddCommand(args[0])
	},
}

var removeSubCommand = &cobra.Command{
	Use:   "remove <path>",
	Short: "Remove to gitignore file",
    Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        if id, err := strconv.Atoi(args[0]); err == nil {
            handleRemoveCommand(id)
            return
        }
        fmt.Println("Please enter a number.")
	},
}

var listSubCommand = &cobra.Command{
	Use:   "list",
	Short: "Managed to gitignore file",
    Run: func(cmd *cobra.Command, args []string) {
        handleListCommand()
    },
}

func handleAddCommand(path string) {
    file := openFile()
    defer file.Close()
    
    _, err := file.WriteString(path + "\n")
    utils.CheckError(err)
    fmt.Printf("%s was added succesfully.\n", path)
}

func handleRemoveCommand(id int) {
    file, err := os.ReadFile(path)
    utils.CheckError(err)
    lines := strings.Split(string(file), "\n")

    if id >= len(lines) {
        fmt.Println("Please enter a valid id.")
        return
    }

    fmt.Printf("%s was remove succesfully.\n", lines[id - 1])
    lines = utils.RemoveIndex(lines, id - 1)

    f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
    defer f.Close()

    _, err = f.WriteString(strings.Join(lines, "\n"))
    utils.CheckError(err)
}

func handleListCommand() {
    file, err := os.ReadFile(path)
    utils.CheckError(err)
    lines := strings.Split(strings.TrimSpace(string(file)), "\n")

    for index, line := range lines {
        fmt.Printf(".gitignore:%d:%s\n", index + 1, line)
    }
}

func openFile() *os.File {
    file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    utils.CheckError(err)
    return file
}

func init() {
    gitignoreCmd.AddCommand(addSubCommand)
    gitignoreCmd.AddCommand(removeSubCommand)
    gitignoreCmd.AddCommand(listSubCommand)
	rootCmd.AddCommand(gitignoreCmd)
}
