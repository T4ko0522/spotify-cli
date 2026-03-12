package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/T4ko0522/spotify-cli/internal/auth"
	"github.com/T4ko0522/spotify-cli/internal/config"
	"github.com/T4ko0522/spotify-cli/internal/player"
	"github.com/T4ko0522/spotify-cli/internal/tui"
	spotify "github.com/zmb3/spotify/v2"
	"github.com/spf13/cobra"
)

var (
	spotifyPlayer *player.Player
	spotifyClient *spotify.Client
)

var rootCmd = &cobra.Command{
	Use:           "spt",
	Short:         "Spotify CLI controller",
	Long:          "A command-line tool to control Spotify playback.",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return tui.Run(spotifyClient)
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// setup/settings commands handle config on their own
		if cmd.Name() == "setup" || cmd.Name() == "settings" {
			return nil
		}
		if err := config.Load(); err != nil {
			return err
		}
		ctx := context.Background()
		httpClient, err := auth.GetClient(ctx)
		if err != nil {
			return fmt.Errorf("%w\nRun 'spt setup' to authenticate", err)
		}
		spotifyClient = spotify.New(httpClient)
		spotifyPlayer = player.New(spotifyClient)
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == "setup" || cmd.Name() == "settings" || spotifyClient == nil {
			return nil
		}
		// Save token in case it was refreshed during this session
		_ = auth.PersistToken()
		return nil
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
