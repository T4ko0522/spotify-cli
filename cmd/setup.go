package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/T4ko0522/spotify-cli/internal/auth"
	"github.com/T4ko0522/spotify-cli/internal/config"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure Client ID and authenticate with Spotify",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Step 1: Client ID 入力
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
		fmt.Println("Client ID saved.")

		// Step 2: config をロードして Login 実行
		if err := config.Load(); err != nil {
			return err
		}
		return auth.Login()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
