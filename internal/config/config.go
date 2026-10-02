package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL    string `json:"db_url"`
	UserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {

	var config Config

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return config, fmt.Errorf("failed to get config file path: %w", err)
	}

	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	return config, nil

}

func (cfg *Config) SetUser(username string) error {

	cfg.UserName = username

	err := write(*cfg)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {

	userHomeDirectory, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user's home directory: %w", err)
	}

	configFilePath := filepath.Join(userHomeDirectory, configFileName)

	return configFilePath, nil
}

func write(cfg Config) error {

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("failed to get config file path: %w", err)
	}

	jsonData, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config struct: %w", err)
	}

	err = os.WriteFile(configFilePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON to file: %w", err)
	}

	return nil
}
