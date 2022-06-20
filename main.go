package main

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/code-raisan/gocolor"
)

func main() {
	config, err := configLoader(dev)
	if err != nil || config.Api == "" {
		fmt.Print(errorMsgs(5, "console", "\n[Log]"), err)
		os.Exit(0)
	}
	if len(os.Args) < 2 {
		fmt.Print(errorMsgs(0, "console", ""))
		os.Exit(0)
	}
	switch os.Args[1] {
	case "version":
		fmt.Printf("\n%s  Version %s\n", app_name, version)
		fmt.Printf("%s: %s\n", gocolor.Yellow("Repository"), repo)
		fmt.Printf("%s: `%s help`\n\n", gocolor.Cyan("Help"), command)
	case "help":
		fmt.Printf("コマンドの使用方法は `%s` を参照してください\nまた対応言語・言語コード一覧は `%s` を参照してください", gocolor.Purple(repo), gocolor.Purple("https://cloud.google.com/translate/docs/languages"))
	case "cat":
		fmt.Print(gocolor.Blue("にゃ～ん =^･ω･^="))
	default:
		isjson := false
		if os.Args[1] == "json" {
			isjson = true
		}
		if len(os.Args) < sif(isjson, 4, 3) {
			fmt.Print(errorMsgs(0, sifs(isjson, "json", "console"), ""))
			os.Exit(0)
		}

		//Spinner Start
		s := spinner.New(spinner.CharSets[35], 50*time.Millisecond)
		s.Start()

		resp, ecode, emsg, err := Translate(urlGen(config.Api, os.Args[sif(isjson, 2, 1)], url.QueryEscape(os.Args[sif(isjson, 3, 2)])))
		if err != nil || ecode != -1 {
			fmt.Print(errorMsgs(ecode, sifs(isjson, "json", "console"), sifs(emsg != "", emsg, "")))
			os.Exit(0)
		}

		//Spinner Stop
		s.Stop()

		// Out put json
		if isjson {
			fmt.Printf(`{"code": %s, "msg": "%s", "text": "%s"}`, "200", "success", resp.Text)
			os.Exit(0)
		}

		// Out put nomal
		fmt.Println(gocolor.Purple("[Before]"))
		fmt.Println(" ", os.Args[sif(isjson, 3, 2)])
		fmt.Print("  ↓\n")
		fmt.Println(gocolor.Green("[After: " + os.Args[sif(isjson, 2, 1)] + "]"))
		fmt.Println(" ", resp.Text)
	}
}
