package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/T4ko0522/spotify-cli/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Configure Spotify Client ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your Spotify Client ID: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}
		clientID := strings.TrimSpace(input)
		if clientID == "" {
			return fmt.Errorf("client ID cannot be empty")
		}
		if err := config.Save(clientID); err != nil {
			return err
		}
		fmt.Println("Client ID saved. Run 'spt login' to authenticate.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
