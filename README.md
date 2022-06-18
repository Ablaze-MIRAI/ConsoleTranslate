# ConsoleTranslate

翻訳をコンソール上で行えるツール

<div align="center">

![ConsoleTranslate](./docs/main.gif)

</div>

# ⭐ 使い方

```powershell
- ヘルプを表示
translate help

- 翻訳先の言語を指定して翻訳
translate -t en こんにちは

- 空白のあるテキストを翻訳
translate -t en "こんにちは 世界"

- 翻訳元の言語も指定する
translate -t en -f ja こんにちは
```

対応している言語の言語コード一覧は [言語サポート  |  Cloud Translation  |  Google Cloud](https://cloud.google.com/translate/docs/languages) を参照

# ⚡ インストール

1. [APIキー発行](./API.md)を参考にGAS(Google Apps Script)でAPIキーを発行

2. [Latest Release](https://github.com/Ablaze-MIRAI/ConsoleTranslate/releases) からビルド済みバイナリをダウンロード

3. ダウンロードしたバイナリを解凍して好きなディレクトリへ移動

4. バイナリを置いたディレクトリにパスを通す

5. `config.json` の `api` に `1.` で発行したAPIキーを設定

```json
{"api": "<<ここにAPIキーを設定>>"}
```

### お疲れ様でした。これでご使用いただけます🎉

# 🌠 開発

```bash
go mod tidy
```

**開発モードに切り替える** *(これを行わないと`go run`が使用できません)*

以下のように書き換えてください

```go
// main.go [Line: 13~19]
const (
    //DEV MODE
	dev = true //false
	app_name = "ConsoleTranslate"
	version  = "1.0.0"
	command  = "translate"
	git_repo = "https://github.com/Ablaze-MIRAI/ConsoleTranslate"
)
```

## 🔧 ビルド

**開発モードに切り替えていた場合は戻してください**

```
# Linux
go build -o translate *.go

# Windows
go build -o translate.exe .
```

# 💌 Special thanks

[Comamoca](https://github.com/Comamoca)

[お餅のCreeeper](https://github.com/creeper-0910)

[nexryai](https://github.com/nexryai)
