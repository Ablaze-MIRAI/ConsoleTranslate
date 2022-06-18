package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/code-raisan/gocolor"
)

const (
	// DEV MODE
	dev      = false
	app_name = "ConsoleTranslate"
	version  = "1.0.0"
	command  = "translate"
	git_repo = "https://github.com/"
)

func main() {
	if isset(1, os.Args) && os.Args[1] == "help" {
		if isset(2, os.Args) && os.Args[2] == "api" {
			fmt.Printf("APIの設定は %s を参照してください", gocolor.Purple(git_repo))
			os.Exit(0)
		} else {
			fmt.Printf("Example\n\n  %s <テキスト> -t [翻訳先] (-f [翻訳元]:任意)\n\n %s -t, --to : 翻訳先の言語コードを指定\n %s -f, --from : 翻訳元の言語コードを指定\n\n 対応している言語の言語コード一覧は `%s` を参照", command, gocolor.Red("(必須)"), gocolor.Cayn("(任意)"), gocolor.Purple("https://cloud.google.com/translate/docs/languages"))
		}
	} else if isset(1, os.Args) && os.Args[1] == "version" {
		fmt.Printf("\n%s v%s\nGithub: %s\nHelp: `%s`\n\n", app_name, version, git_repo, gocolor.Cayn(command+" help"))
	} else {
		conf, err := loadConfig(dev)
		if err != nil {
			fmt.Printf("%s: 設定ファイルの読み込みに失敗\n設定は `%s` を参照してください", gocolor.Red("Error"), gocolor.Purple(git_repo))
			os.Exit(0)
		}
		to, to_fi := getFlag("-t", "--to", os.Args)
		from, from_fi := getFlag("-f", "--from", os.Args)
		flags := []int{to_fi, to_fi + 1, from_fi, from_fi + 1}
		var args []string
		for i, v := range os.Args {
			if i != 0 {
				if !contains(flags, i) {
					args = append(args, v)
				}
			}
		}
		if to == "" {
			fmt.Printf("%s: 必要な引数がありません\n詳細は `%s` を参照してください。", gocolor.Red("Error"), gocolor.Cayn(command+" help"))
			os.Exit(0)
		}

		response := &struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
			Text string `json:"text"`
		}{}

		resp, err := http.Get(urlGen(to, from, args[0], conf.Api))
		if err != nil {
			fmt.Printf("%s: リクエストに失敗しました\nインターネットの接続、APIの設定等を確認してください\n[Log]%s", gocolor.Red("Error"), err)
			os.Exit(0)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Printf("%s: リクエストに失敗しました\n[Log]HTTP Status: `%s`", gocolor.Red("Error"), resp.Status)
			os.Exit(0)
		}

		body, _ := io.ReadAll(resp.Body)
		if err := json.Unmarshal(body, response); err != nil {
			fmt.Printf("%s: リクエストの解析に失敗しました\n[Log]%s", gocolor.Red("Error"), err)
			os.Exit(0)
		}
		if response.Msg == "unexpected" {
			fmt.Printf("%s: 翻訳に失敗しました\n翻訳に対応している言語は `%s` を参照してください。\n[Log]API Error", gocolor.Red("Error"), gocolor.Purple("https://cloud.google.com/translate/docs/languages"))
			os.Exit(0)
		}

		var lang_info string
		if from == "" {
			lang_info = "auto"
		} else {
			lang_info = from
		}
		fmt.Printf("%s\n %s", gocolor.Purple("[Before: "+lang_info+"]"), args[0])
		fmt.Print("\n  ↓\n")
		fmt.Printf("%s\n %s", gocolor.Green("[After: "+to+"]"), response.Text)
		os.Exit(0)
	}
}
