package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var ClientID string

func Load() error {
	_ = godotenv.Load() // .env is optional; environment variables take precedence
	ClientID = os.Getenv("SPOTIFY_CLIENT_ID")
	if ClientID == "" {
		return fmt.Errorf("SPOTIFY_CLIENT_ID is not set. Create a .env file or set the environment variable")
	}
	return nil
}
