# spotify cli

Spotify の再生をコマンドラインから操作するツールです。

## 機能

- 再生 / 一時停止 / 曲のスキップ
- 現在再生中の曲情報の表示
- 音量の調整
- 接続デバイス一覧の表示
- アクティブデバイスの自動検出・選択

## 必要要件

- Go 1.25 以上
- [Spotify Developer](https://developer.spotify.com/) で作成したアプリの Client ID
- リダイレクト URI: `http://127.0.0.1:8888/callback`

## インストール

### ソースからビルド

```bash
go install github.com/T4ko0522/spotify-cli@latest
```

### MSI インストーラー（Windows）

[Release](https://github.com/T4ko0522/spotify-cli/Releases) から `spt.msi` をダウンロードするか、ローカルでビルドしてください:

```bat
build.bat
```

## セットアップ

1. Client ID を設定します:

```bash
spt init
```

2. Spotify にログインします:

```bash
spt login
```

## 使い方

| コマンド | エイリアス | 説明 |
|---|---|---|
| `spt init` | | Client ID を設定 |
| `spt login` | | Spotify で認証 |
| `spt play` | `spt p` | 再生を再開 |
| `spt pause` | | 一時停止 |
| `spt next` | `spt n` | 次の曲へ |
| `spt back` | `spt b` | 前の曲へ |
| `spt now` | | 現在再生中の曲を表示 |
| `spt volume` | `spt v` | 現在の音量を表示 |
| `spt volume [0-100]` | `spt v [0-100]` | 音量を設定 |
| `spt devices` | `spt d` | 利用可能なデバイス一覧を表示 |

## ライセンス

MIT
