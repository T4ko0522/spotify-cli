package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/T4ko0522/spotify-cli/internal/config"
	"github.com/spf13/cobra"
)

var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Interactively change TUI settings",
	Long:  "Interactively change TUI image size preset.",
	RunE: func(cmd *cobra.Command, args []string) error {
		_ = config.Load()

		fmt.Printf("\n現在の画像サイズ: %s\n\n", config.ImgSize)
		fmt.Println("プリセット:")
		for i, name := range config.ImgSizeNames {
			p := config.ImgSizePresets[name]
			marker := "  "
			if name == config.ImgSize {
				marker = "* "
			}
			fmt.Printf("  %s%d. %s (%dx%d)\n", marker, i+1, name, p.Cols, p.Rows)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\n番号を選択 [1-%d]: ", len(config.ImgSizeNames))
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("\n変更はありません。")
			return nil
		}

		var idx int
		if _, err := fmt.Sscanf(input, "%d", &idx); err != nil || idx < 1 || idx > len(config.ImgSizeNames) {
			return fmt.Errorf("1-%d の番号を入力してください", len(config.ImgSizeNames))
		}

		selected := config.ImgSizeNames[idx-1]
		if selected == config.ImgSize {
			fmt.Println("\n変更はありません。")
			return nil
		}

		if err := config.SaveSettings(selected); err != nil {
			return fmt.Errorf("failed to save settings: %w", err)
		}

		p := config.ImgSizePresets[selected]
		fmt.Printf("\n設定を保存しました: %s (%dx%d)\n", selected, p.Cols, p.Rows)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(settingsCmd)
}
