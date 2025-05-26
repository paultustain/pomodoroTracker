package config

import (
	"os"
	"path/filepath"
)

const configFile = ".pomodoroconfig.json"

func GetConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homeDir, configFile)
	return fullPath, nil
}
