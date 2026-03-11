package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ImgSizePreset struct {
	Cols int
	Rows int
}

var ImgSizePresets = map[string]ImgSizePreset{
	"small":  {Cols: 16, Rows: 8},
	"medium": {Cols: 20, Rows: 10},
	"large":  {Cols: 28, Rows: 14},
}

var ImgSizeNames = []string{"small", "medium", "large"}

var (
	ClientID string
	ImgSize  = "medium"
	ImgCols  = ImgSizePresets["medium"].Cols
	ImgRows  = ImgSizePresets["medium"].Rows
)

type configData struct {
	ClientID string `json:"client_id"`
	ImgSize  string `json:"img_size,omitempty"`
}

func applyPreset(name string) {
	if p, ok := ImgSizePresets[name]; ok {
		ImgSize = name
		ImgCols = p.Cols
		ImgRows = p.Rows
	}
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
	dir := filepath.Join(appData, "spt")
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
			if err := json.Unmarshal(data, &cfg); err == nil {
				if cfg.ClientID != "" {
					ClientID = cfg.ClientID
				}
				if cfg.ImgSize != "" {
					applyPreset(cfg.ImgSize)
				}
				if ClientID != "" {
					return nil
				}
			}
		}
	}

	// 2. Fall back to environment variable
	ClientID = os.Getenv("SPOTIFY_CLIENT_ID")
	if ClientID != "" {
		return nil
	}

	return fmt.Errorf("SPOTIFY_CLIENT_ID is not configured. Run 'spt init' to set up")
}

func SaveSettings(imgSize string) error {
	path, err := configPath()
	if err != nil {
		return err
	}

	var cfg configData
	if data, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(data, &cfg)
	}

	cfg.ImgSize = imgSize

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal config: %w", err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("cannot write config file: %w", err)
	}

	applyPreset(imgSize)
	return nil
}
