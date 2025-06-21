package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFile = ".pomodoroconfig.json"

type Config struct {
	DBURL string `json:"db_url"`
}

func Read() (Config, error) {
	configPath, err := GetConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)

	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

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
