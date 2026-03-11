package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var ClientID string

type configData struct {
	ClientID string `json:"client_id"`
}

func configPath() (string, error) {
	appData := os.Getenv("APPDATA")
	if appData == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("cannot determine home directory: %w", err)
		}
		appData = filepath.Join(home, ".config")
	}
	dir := filepath.Join(appData, "spty")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", fmt.Errorf("cannot create config directory: %w", err)
	}
	return filepath.Join(dir, "config.json"), nil
}

func Save(clientID string) error {
	path, err := configPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(configData{ClientID: clientID}, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal config: %w", err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("cannot write config file: %w", err)
	}
	return nil
}

func Load() error {
	// 1. Try config file
	path, err := configPath()
	if err == nil {
		if data, err := os.ReadFile(path); err == nil {
			var cfg configData
			if err := json.Unmarshal(data, &cfg); err == nil && cfg.ClientID != "" {
				ClientID = cfg.ClientID
				return nil
			}
		}
	}

	// 2. Fall back to environment variable
	ClientID = os.Getenv("SPOTIFY_CLIENT_ID")
	if ClientID != "" {
		return nil
	}

	return fmt.Errorf("SPOTIFY_CLIENT_ID is not configured. Run 'spty init' to set up")
}
