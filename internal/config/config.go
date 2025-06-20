package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const configFile = ".pomodoroconfig.json"

func Read()

// NEed to read config file path and use the values inside config.
// Currently using the link to the path not the value

func GetConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	fmt.Printf("")
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homeDir, configFile)
	return fullPath, nil
}
