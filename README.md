# ConsoleTranslate

翻訳をコンソール上で行えるツール

<div align="center">

![ConsoleTranslate](./docs/image_v2.gif)

</div>

### ⚠ 1系(1.x.x)とは仕様が1部異なります

# ⭐ 使い方

```bash
# ヘルプを表示
translate help

# 翻訳先の言語を指定して翻訳
translate en こんにちは

# 空白のあるテキストを翻訳
translate en "こんにちは 世界"

# 結果をJSONで出力
translate json en こんにちは
```

対応している言語の言語コード一覧は [言語サポート  |  Cloud Translation  |  Google Cloud](https://cloud.google.com/translate/docs/languages) を参照

# ⚡ インストール

#### ⚠ 設定ファイル(config.json)は1系(1.x.x)と同じなので1系をお使いの方はバイナリを置き換えるだけで使用できます

1. [Latest Release](https://github.com/Ablaze-MIRAI/ConsoleTranslate/releases) からビルド済みバイナリをダウンロード

2. ダウンロードしたバイナリを解凍して好きなディレクトリへ移動

3. バイナリを置いたディレクトリにパスを通す

4. *DeepL(v2.0.3より対応)* または *GoogleTranslate(GAS)* のAPIキーを発行してください。
  
   [APIキー発行の手順(DeepL・Google共通)](./API.md)を参考にしてください。

5. `config.json`を同じ階層に作成し`4.`で発行したAPIキーを設定します。

- Google Translate(GAS)でAPIキーを発行した場合

```json
{"api": "[ここにAPIキーを設定]"}
```

- DeepLでAPIキーを発行した場合

アカウントタイプはDeepL Freeの場合は`free`をDeepL Proの場合は`pro`を設定してください。

APIキーは[APIキー発行の手順(DeepL・Google共通)](./API.md)で発行したAPIキー(認証キー)を設定してください

```json
{"api": "deepl,[アカウントタイプ],[APIキー]"}
```

### お疲れ様でした。これでご使用いただけます🎉

# 🌠 開発

### ⚠ 開発は1系(1.x.x)から大きく変更があります

```bash
# Clone this repository

# Linux
cp config.go.sample config.go
cp config.json.sample config.json

# Windows (PowerShell 7.x)
copy config.go.sample config.go
copy config.json.sample config.json

go mod tidy
```

**開発モードに切り替える** *(これを行わないと`go run`が使用できません)*

`config.go`以下のように書き換えてください

```go
// config.go [Line: 6~9]

repo     = "https://github.com/Ablaze-MIRAI/ConsoleTranslate"
command  = "translate"
dev      = true //false
```

## 🔧 ビルド

**開発モードに切り替えていた場合は戻してください**

```bash
# Linux
cp config.go.sample config.go
cp config.json.sample config.json
go build -o translate *.go


# Windows (PowerShell 7.x)
copy config.go.sample config.go
copy config.json.sample config.json
go build -o translate.exe .
```

# 💌 Special thanks

[Comamoca](https://github.com/Comamoca)

[お餅のCreeeper](https://github.com/creeper-0910)

[nexryai](https://github.com/nexryai)
