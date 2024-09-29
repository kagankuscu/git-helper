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
    dir := strings.Split(gitFile, ".")[0]
    return dir
}

func CheckError(err error) {
    if err != nil {
        fmt.Println("Error:", err)
    }
}
