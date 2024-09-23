package config

import (
    "fmt"
    "os"
)

func UpdateConfig(key, value string) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        fmt.Println("Error fetching home directory:", err)
        return
    }

    configFile := homeDir + "/.githelperconfig"

    // Open the config file for appending
    file, err := os.OpenFile(configFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening config file:", err)
        return
    }
    defer file.Close()

    if _, err := file.WriteString(fmt.Sprintf("%s = %s\n", key, value)); err != nil {
        fmt.Println("Error writing to config file:", err)
    }

    fmt.Println("Config updated:", key, "=", value)
}
