package config

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Config struct {
    Remote        string
    DefaultBranch string
}

func LoadConfig() Config {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        fmt.Println("Error fetching home directory:", err)
        return Config{Remote: "origin", DefaultBranch: "main"} // Defaults
    }

    configFile := homeDir + "/.githelperconfig"
    file, err := os.Open(configFile)
    if err != nil {
        // Return default values if config file is missing
        return Config{Remote: "origin", DefaultBranch: "main"}
    }
    defer file.Close()

    config := Config{Remote: "origin", DefaultBranch: "main"} // Defaults
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "remote") {
            config.Remote = strings.TrimSpace(strings.Split(line, "=")[1])
        } else if strings.HasPrefix(line, "default_branch") {
            config.DefaultBranch = strings.TrimSpace(strings.Split(line, "=")[1])
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading config file:", err)
    }

    return config
}
