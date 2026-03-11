package cmd

import (
	"github.com/T4ko0522/spotify-cli/internal/auth"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate with Spotify",
	RunE: func(cmd *cobra.Command, args []string) error {
		return auth.Login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
