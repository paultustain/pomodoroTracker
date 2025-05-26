package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFile = ".pomodoroconfig.json"

type Config struct {
	DBURL string `json:"db_url"`
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)

	return cfg, err
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homeDir, configFile)
	return fullPath, nil
}
