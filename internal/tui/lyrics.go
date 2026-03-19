package tui

import (
	"context"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/T4ko0522/spotify-cli/internal/lyrics"
	"github.com/zmb3/spotify/v2"
)

var (
	lyricsCurrentStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10"))
	lyricsNormalStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	lyricsTitleStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	lyricsHelpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	lyricsErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
)

type lyricsTickMsg time.Time

type lyricsDataMsg struct {
	result *lyrics.Result
	err    error
}

type lyricsPlayingMsg struct {
	playing *spotify.CurrentlyPlaying
}

type LyricsModel struct {
	client   *spotify.Client
	playing  *spotify.CurrentlyPlaying
	lyrics   *lyrics.Result
	err      error
	width    int
	height   int
	quitting bool
	trackID  spotify.ID
	loading  bool
}

func NewLyricsModel(client *spotify.Client) LyricsModel {
	return LyricsModel{
		client:  client,
		width:   80,
		height:  24,
		loading: true,
	}
}

func (m LyricsModel) Init() tea.Cmd {
	return tea.Batch(m.fetchPlaying, lyricsTickCmd())
}

func lyricsTickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return lyricsTickMsg(t)
	})
}

func (m LyricsModel) fetchPlaying() tea.Msg {
	playing, err := m.client.PlayerCurrentlyPlaying(context.Background())
	if err != nil {
		return lyricsPlayingMsg{}
	}
	return lyricsPlayingMsg{playing: playing}
}

func fetchLyrics(trackName, artistName string, durationMs int) tea.Cmd {
	return func() tea.Msg {
		result, err := lyrics.Fetch(trackName, artistName, durationMs)
		return lyricsDataMsg{result: result, err: err}
	}
}

func (m LyricsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case lyricsTickMsg:
		return m, tea.Batch(m.fetchPlaying, lyricsTickCmd())
	case lyricsPlayingMsg:
		m.playing = msg.playing
		if msg.playing != nil && msg.playing.Item != nil {
			newID := msg.playing.Item.ID
			if newID != m.trackID {
				m.trackID = newID
				m.loading = true
				m.lyrics = nil
				m.err = nil
				item := msg.playing.Item
				return m, fetchLyrics(
					item.Name,
					formatArtists(item.Artists),
					int(item.Duration),
				)
			}
		}
	case lyricsDataMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
			m.lyrics = nil
		} else {
			m.lyrics = msg.result
			m.err = nil
		}
	}
	return m, nil
}

func (m LyricsModel) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder

	// Header
	if m.playing != nil && m.playing.Item != nil {
		item := m.playing.Item
		b.WriteString(lyricsTitleStyle.Render(fmt.Sprintf("♪ %s - %s", item.Name, formatArtists(item.Artists))))
		b.WriteString("\n")
		b.WriteString(strings.Repeat("─", min(m.width, 60)))
		b.WriteString("\n\n")
	} else {
		b.WriteString("Nothing is currently playing.\n\n")
		b.WriteString(lyricsHelpStyle.Render("Press q to quit"))
		b.WriteString("\n")
		return b.String()
	}

	if m.loading {
		b.WriteString("Loading lyrics...\n")
		b.WriteString("\n")
		b.WriteString(lyricsHelpStyle.Render("Press q to quit"))
		b.WriteString("\n")
		return b.String()
	}

	if m.err != nil {
		b.WriteString(lyricsErrorStyle.Render(fmt.Sprintf("Lyrics not available: %v", m.err)))
		b.WriteString("\n\n")
		b.WriteString(lyricsHelpStyle.Render("Press q to quit"))
		b.WriteString("\n")
		return b.String()
	}

	if m.lyrics == nil {
		b.WriteString("No lyrics found.\n\n")
		b.WriteString(lyricsHelpStyle.Render("Press q to quit"))
		b.WriteString("\n")
		return b.String()
	}

	// Available lines for lyrics display (header=3, help=2, bottom margin=1)
	maxLines := m.height - 6
	if maxLines < 5 {
		maxLines = 5
	}

	if m.lyrics.Synced {
		m.renderSyncedLyrics(&b, maxLines)
	} else {
		m.renderPlainLyrics(&b, maxLines)
	}

	b.WriteString("\n")
	b.WriteString(lyricsHelpStyle.Render("Press q to quit"))
	b.WriteString("\n")

	return b.String()
}

func (m LyricsModel) renderSyncedLyrics(b *strings.Builder, maxLines int) {
	lines := m.lyrics.SyncedLyrics
	progress := 0
	if m.playing != nil {
		progress = int(m.playing.Progress)
	}

	currentIdx := lyrics.CurrentLineIndex(lines, progress)

	// Center the current line in the viewport
	startIdx := currentIdx - maxLines/2
	if startIdx < 0 {
		startIdx = 0
	}
	endIdx := startIdx + maxLines
	if endIdx > len(lines) {
		endIdx = len(lines)
		startIdx = endIdx - maxLines
		if startIdx < 0 {
			startIdx = 0
		}
	}

	for i := startIdx; i < endIdx; i++ {
		line := lines[i].Text
		if line == "" {
			line = "♪"
		}
		if i == currentIdx {
			b.WriteString(lyricsCurrentStyle.Render("▸ " + line))
		} else {
			b.WriteString(lyricsNormalStyle.Render("  " + line))
		}
		b.WriteString("\n")
	}
}

func (m LyricsModel) renderPlainLyrics(b *strings.Builder, maxLines int) {
	lines := strings.Split(m.lyrics.PlainLyrics, "\n")
	for i, line := range lines {
		if i >= maxLines {
			b.WriteString(lyricsNormalStyle.Render(fmt.Sprintf("  ... (%d more lines)", len(lines)-i)))
			b.WriteString("\n")
			break
		}
		if line == "" {
			b.WriteString("\n")
		} else {
			b.WriteString(lyricsNormalStyle.Render("  " + line))
			b.WriteString("\n")
		}
	}
}

func RunLyrics(client *spotify.Client) error {
	p := tea.NewProgram(NewLyricsModel(client), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
