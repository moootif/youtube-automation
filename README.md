# youtube-automation

このプロジェクトは、YouTubeの自動化を目的としたアプリケーションです。HTTPサーバーを起動し、HTMLテンプレートをレンダリングしてユーザーに表示します。

## ファイル構成

- **main.go**: アプリケーションのエントリーポイント。HTTPサーバーを起動し、HTMLテンプレートを表示します。
- **configs/config.json**: アプリケーションの設定をJSON形式で保存します。`UploadDir`や`DefaultTitle`などの設定項目が含まれます。
- **templates/index.html**: ユーザーに表示されるウェブページのHTMLテンプレートです。
- **static/css/style.css**: HTMLページのスタイルを定義するスタイルシートです。
- **static/js/app.js**: HTMLページの動的な動作を制御するJavaScriptファイルです。
- **go.mod**: Goモジュールの設定ファイル。依存関係やモジュール名を管理します。
- **.gitignore**: Gitで無視するファイルやディレクトリを指定します。

## セットアップ手順

1. リポジトリをクローンします。
2. 必要な依存関係をインストールします。
   ```
   go mod tidy
   ```
3. アプリケーションを実行します。
   ```
   go run main.go
   ```
4. ブラウザで `http://localhost:8080` にアクセスします。

## 使用技術

- Go言語
- HTML/CSS/JavaScript
- JSON

このプロジェクトは、YouTubeの自動化に役立つツールを提供することを目的としています。

http://192.168.0.19:8080