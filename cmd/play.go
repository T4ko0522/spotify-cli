package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Resume playback",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := spotifyPlayer.Play(cmd.Context()); err != nil {
			return err
		}
		fmt.Println("Playback resumed.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
