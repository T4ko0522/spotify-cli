# Spotify CLI + TUI

English | [日本語](README.ja.md)

A command-line tool to control Spotify playback.

![Screenshot](assets/readme.png)

## ✨ Features

- Play / Pause / Skip tracks
- Now Playing display (TUI)
- Album art display (WezTerm supported)
- Synced lyrics display (via [LRCLIB](https://lrclib.net))
- Volume control (TUI / command)
- List available devices
- Auto-detect and select active device
- Image size presets (small / medium / large)

## 📋 Requirements

- Client ID from an app created on [Spotify Developer](https://developer.spotify.com/)
- Redirect URI: `http://127.0.0.1:8888/callback`

## 📦 Installation

### Windows

Download `spt.msi` from [Releases](https://github.com/T4ko0522/Spotify-CLI/releases) and run it.

### macOS / Linux

#### From source (requires Go 1.25+)

```bash
go install github.com/T4ko0522/spotify-cli@latest
```

#### From binary

Download the appropriate binary for your platform from [Releases](https://github.com/T4ko0522/Spotify-CLI/releases) and place it in your `$PATH`.

## 🚀 Setup

Configure your Client ID and authenticate with Spotify:

```bash
spt setup
```

## 📖 Usage

### 🖥️ TUI Mode

Run without arguments to launch the TUI, which shows the currently playing track in real time. Album art is displayed in WezTerm.

```bash
spt
```

### 📋 Command Reference

| Command | Alias | Description |
|---|---|---|
| `spt` | | Launch TUI (Now Playing) |
| `spt --lyrics` | `spt -l` | Show synced lyrics |
| `spt setup` | | Configure Client ID & Spotify auth |
| `spt play` | `spt p` | Resume playback |
| `spt stop` | `spt s` | Pause |
| `spt next` | `spt n` | Next track |
| `spt back` | `spt b` | Previous track |
| `spt now` | | Show currently playing track |
| `spt volume` | `spt v` | Launch volume control TUI |
| `spt volume [0-100]` | `spt v [0-100]` | Set volume |
| `spt devices` | `spt d` | List available devices |
| `spt settings` | | Change image size preset |

### ⚙️ Settings

Use `spt settings` to change the image size preset. Use arrow keys to select and Enter to confirm.

| Preset | Size |
|---|---|
| small | 16×8 |
| medium | 20×10 (default) |
| large | 28×14 |

## 📄 License

[Apache 2.0 LICENSE](https://github.com/T4ko0522/Spotify-CLI/blob/main/LICENSE)
