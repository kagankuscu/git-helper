package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetGitDirectory() string {
    out, err := exec.Command("git", "rev-parse", "--git-dir").Output()
    CheckError(err)

    gitFile := string(out)
    splited := strings.Split(gitFile, ".")
    dir := splited[0]
    if len(splited) > 1 {
        dir = "."
    }
    return dir
}

func CheckError(err error) {
    if err != nil {
        fmt.Println("Error:", err)
    }
}
