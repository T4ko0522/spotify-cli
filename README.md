# Spotify CLI + TUI

Spotify の再生をコマンドラインから操作するツールです。

![スクリーンショット](assets/readme.png)

## ✨ 機能

- 再生 / 一時停止 / 曲のスキップ
- 現在再生中の曲情報の表示（TUI）
- アルバムアート表示（WezTerm 対応）
- 音量の調整（TUI / コマンド）
- 接続デバイス一覧の表示
- アクティブデバイスの自動検出・選択
- 画像サイズプリセットの設定（small / medium / large）

## 📋 必要要件

- Go 1.25 以上
- [Spotify Developer](https://developer.spotify.com/) で作成したアプリの Client ID
- リダイレクト URI: `http://127.0.0.1:8888/callback`

## 📦 インストール

### 🔨 ソースからビルド

```bash
go install github.com/T4ko0522/spotify-cli@latest
```

### 🪟 MSI インストーラー（Windows）

[Release](https://github.com/T4ko0522/spotify-cli/Releases) から `spt.msi` をダウンロードして実行してください。

## 🚀 セットアップ

Client ID の設定と Spotify 認証を行います:

```bash
spt setup
```

## 📖 使い方

### 🖥️ TUI モード

引数なしで実行すると、現在再生中の曲情報をリアルタイム表示する TUI が起動します。
WezTerm ではアルバムアートも表示されます。

```bash
spt
```

### 📋 コマンド一覧

| コマンド | エイリアス | 説明 |
|---|---|---|
| `spt` | | TUI を起動（Now Playing） |
| `spt setup` | | Client ID 設定 & Spotify 認証 |
| `spt play` | `spt p` | 再生を再開 |
| `spt pause` | `spt stop`, `spt s` | 一時停止 |
| `spt next` | `spt n` | 次の曲へ |
| `spt back` | `spt b` | 前の曲へ |
| `spt now` | | 現在再生中の曲を表示 |
| `spt volume` | `spt v` | 音量調整 TUI を起動 |
| `spt volume [0-100]` | `spt v [0-100]` | 音量を設定 |
| `spt devices` | `spt d` | 利用可能なデバイス一覧を表示 |
| `spt settings` | | 画像サイズプリセットを変更 |

### ⚙️ 設定

`spt settings` で画像サイズプリセットを変更できます。矢印キーで選択し、Enter で確定します。

| プリセット | サイズ |
|---|---|
| small | 16×8 |
| medium | 20×10（デフォルト） |
| large | 28×14 |

## 📄 ライセンス

[Apache 2.0 LICENSE](https://github.com/T4ko0522/Spotify-CLI/blob/main/LICENSE)
