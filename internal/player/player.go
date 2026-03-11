package player

import (
	"context"
	"fmt"
	"strings"

	"github.com/zmb3/spotify/v2"
)

type Player struct {
	Client *spotify.Client
}

func New(client *spotify.Client) *Player {
	return &Player{Client: client}
}

func (p *Player) Play(ctx context.Context) error {
	if err := p.Client.Play(ctx); err != nil {
		return fmt.Errorf("failed to resume playback: %w", err)
	}
	return nil
}

func (p *Player) Pause(ctx context.Context) error {
	if err := p.Client.Pause(ctx); err != nil {
		return fmt.Errorf("failed to pause playback: %w", err)
	}
	return nil
}

func (p *Player) Next(ctx context.Context) error {
	if err := p.Client.Next(ctx); err != nil {
		return fmt.Errorf("failed to skip to next track: %w", err)
	}
	return nil
}

func (p *Player) Previous(ctx context.Context) error {
	if err := p.Client.Previous(ctx); err != nil {
		return fmt.Errorf("failed to skip to previous track: %w", err)
	}
	return nil
}

func (p *Player) SetVolume(ctx context.Context, percent int) error {
	if err := p.Client.Volume(ctx, percent); err != nil {
		return fmt.Errorf("failed to set volume: %w", err)
	}
	return nil
}

func (p *Player) NowPlaying(ctx context.Context) (*spotify.CurrentlyPlaying, error) {
	playing, err := p.Client.PlayerCurrentlyPlaying(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get currently playing: %w", err)
	}
	return playing, nil
}

func (p *Player) PlayerState(ctx context.Context) (*spotify.PlayerState, error) {
	state, err := p.Client.PlayerState(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get player state: %w", err)
	}
	return state, nil
}

func (p *Player) Devices(ctx context.Context) ([]spotify.PlayerDevice, error) {
	devices, err := p.Client.PlayerDevices(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get devices: %w", err)
	}
	return devices, nil
}

func FormatArtists(artists []spotify.SimpleArtist) string {
	names := make([]string, len(artists))
	for i, a := range artists {
		names[i] = a.Name
	}
	return strings.Join(names, ", ")
}

func FormatProgress(ms, total int) string {
	current := fmt.Sprintf("%d:%02d", ms/60000, (ms/1000)%60)
	end := fmt.Sprintf("%d:%02d", total/60000, (total/1000)%60)
	return current + " / " + end
}
