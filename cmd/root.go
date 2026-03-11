package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/T4ko0522/spotify-cli/internal/auth"
	"github.com/T4ko0522/spotify-cli/internal/config"
	"github.com/T4ko0522/spotify-cli/internal/player"
	spotify "github.com/zmb3/spotify/v2"
	"github.com/spf13/cobra"
)

var (
	spotifyPlayer *player.Player
	spotifyClient *spotify.Client
)

var rootCmd = &cobra.Command{
	Use:   "spty",
	Short: "Spotify CLI controller",
	Long:  "A command-line tool to control Spotify playback.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// init command doesn't need config yet
		if cmd.Name() == "init" {
			return nil
		}
		if err := config.Load(); err != nil {
			return err
		}
		// login command doesn't need an existing token
		if cmd.Name() == "login" {
			return nil
		}
		ctx := context.Background()
		authenticator, token, err := auth.GetClient(ctx)
		if err != nil {
			return fmt.Errorf("%w\nRun 'spty login' to authenticate", err)
		}
		httpClient := authenticator.Client(ctx, token)
		spotifyClient = spotify.New(httpClient)
		spotifyPlayer = player.New(spotifyClient)
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "login" || spotifyClient == nil {
			return nil
		}
		// Re-save token in case it was refreshed during this session
		_, token, err := auth.GetClient(context.Background())
		if err != nil {
			return nil // not critical
		}
		_ = auth.SaveToken(token)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
