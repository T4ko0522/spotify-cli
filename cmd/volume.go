package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var volumeCmd = &cobra.Command{
	Use:     "volume [0-100]",
	Aliases: []string{"v"},
	Short:   "Get or set playback volume",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		if len(args) == 0 {
			state, err := spotifyPlayer.PlayerState(ctx)
			if err != nil {
				return err
			}
			fmt.Printf("Volume: %d%%\n", state.Device.Volume)
			return nil
		}
		vol, err := strconv.Atoi(args[0])
		if err != nil || vol < 0 || vol > 100 {
			return fmt.Errorf("volume must be a number between 0 and 100")
		}
		if err := spotifyPlayer.SetVolume(ctx, vol); err != nil {
			return err
		}
		fmt.Printf("Volume: %d%%\n", vol)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(volumeCmd)
}
