package cmd

import (
	"fmt"
	"time"

	"github.com/T4ko0522/spotify-cli/internal/player"
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:     "next",
	Aliases: []string{"n"},
	Short:   "Skip to next track",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := spotifyPlayer.Next(cmd.Context()); err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)
		playing, err := spotifyPlayer.NowPlaying(cmd.Context())
		if err != nil {
			fmt.Println("Skipped to next track.")
			return nil
		}
		if playing.Item != nil {
			fmt.Printf("Now playing: %s - %s\n", playing.Item.Name, player.FormatArtists(playing.Item.Artists))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)
}
