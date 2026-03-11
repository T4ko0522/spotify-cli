package cmd

import (
	"fmt"

	"github.com/T4ko0522/spotify-cli/internal/player"
	"github.com/spf13/cobra"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Show currently playing track",
	RunE: func(cmd *cobra.Command, args []string) error {
		playing, err := spotifyPlayer.NowPlaying(cmd.Context())
		if err != nil {
			return err
		}
		if playing.Item == nil {
			fmt.Println("Nothing is currently playing.")
			return nil
		}
		item := playing.Item
		fmt.Printf("Track:    %s\n", item.Name)
		fmt.Printf("Artist:   %s\n", player.FormatArtists(item.Artists))
		fmt.Printf("Album:    %s\n", item.Album.Name)
		fmt.Printf("Progress: %s\n", player.FormatProgress(int(playing.Progress), int(item.Duration)))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)
}
