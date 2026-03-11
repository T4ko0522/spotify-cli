package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause playback",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := spotifyPlayer.Pause(cmd.Context()); err != nil {
			return err
		}
		fmt.Println("Playback paused.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}
