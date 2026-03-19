package auth

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/T4ko0522/spotify-cli/internal/config"
	"golang.org/x/oauth2"
)

func tokenPath() (string, error) {
	dir, err := config.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "token.json"), nil
}

func SaveToken(token *oauth2.Token) error {
	path, err := tokenPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal token: %w", err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("cannot write token file: %w", err)
	}
	return nil
}

func LoadToken() (*oauth2.Token, error) {
	path, err := tokenPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("no saved token found. Run 'spt login' first")
	}
	var token oauth2.Token
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, fmt.Errorf("corrupt token file: %w", err)
	}
	return &token, nil
}
