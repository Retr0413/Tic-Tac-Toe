# Go製GUI三目並べ（Tic-Tac-Toe）

Go言語とFyneライブラリを使ったGUI三目並べアプリケーションです。

---

## 必要条件

- Go 1.18 以上（[Go公式サイト](https://go.dev/dl/)からインストールできます）
- Fyne（Go用のGUIライブラリ）

---

## セットアップ手順

1. リポジトリをクローン or ソースコードをダウンロードして任意のディレクトリへ移動

2. Goモジュールの初期化  
   （すでに`go.mod`がある場合はこの手順は不要です）

   ```bash
   go mod init tictactoe

    追加パッケージ（Fyne）のインストール
    go get fyne.io/fyne/v2

    依存関係を整理
    go mod tidy

    アプリの起動方法
    go run Tic-Tac-Toe.go
