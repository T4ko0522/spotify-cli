package cmd

import (
	"errors"
	"fmt"

	"github.com/T4ko0522/spotify-cli/internal/player"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause playback",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := spotifyPlayer.Pause(cmd.Context()); err != nil {
			if errors.Is(err, player.ErrAlreadyPaused) {
				fmt.Println("Already paused.")
				return nil
			}
			return err
		}
		fmt.Println("Playback paused.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}
