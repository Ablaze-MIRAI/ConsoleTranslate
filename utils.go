package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/code-raisan/gocolor"
)

func errorMsgs(ErrorNumber int, OutPutType string, custom string) string {
	errors := [][]string{
		{"400", "arguments_error", "必要な引数がありません\n詳細は" + gocolor.Cyan("`"+command+" help`") + "を参照してください"},
		{"400", "request_error", "リクエストに失敗しました\nインターネットの接続、APIの設定等を確認してください"},
		{"400", "respose_error", "リクエストに失敗しました"},
		{"400", "parse_error", "リクエストの解析に失敗しました"},
		{"500", "translation_error", "翻訳に失敗しました\n翻訳に対応している言語は `" + gocolor.Purple("https://cloud.google.com/translate/docs/languages") + "` を参照してください。"},
		{"400", "config_error", "設定ファイルの読み込み失敗しました\n設定の方法は `" + gocolor.Purple(repo) + "` を参照してください"}}

	if OutPutType == "console" {
		return gocolor.Red("Error["+strconv.Itoa(ErrorNumber)+"]: ") + errors[ErrorNumber][2] + custom
	} else if OutPutType == "json" {
		return `{"code": "` + errors[ErrorNumber][0] + `", "msg": "` + errors[ErrorNumber][1] + `"}`
	}
	return ""
}

type Config struct {
	Api string `json:"api"`
}

func configLoader(dev bool) (*Config, error) {
	// Set Path
	var path string
	if dev {
		path = "./config.dev.json"
	} else {
		prosess_path, err := os.Executable()
		if err != nil {
			return nil, err
		}
		path = filepath.Dir(prosess_path) + "/config.json"
	}

	// Load file
	var cfg Config
	config_json, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// Set File close
	defer config_json.Close()
	err = json.NewDecoder(config_json).Decode(&cfg)
	if err != nil {
		return &cfg, err
	}

	return &cfg, nil
}

func sif(target bool, t int, f int) int {
	if target {
		return t
	}
	return f
}

func sifs(target bool, t string, f string) string {
	if target {
		return t
	}
	return f
}

func urlGen(Api string, Tolang string, Text string) string {
	return Api + "?to=" + Tolang + "&text=" + Text
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Text string `json:"text"`
}

func Translate(URL string) (*Response, int, string, error) {
	// Start Request
	resp, err := http.Get(URL)
	if err != nil {
		return nil, 1, "", err
	}

	// Set Connection close
	defer resp.Body.Close()

	// Purse Response
	if resp.StatusCode != 200 {
		return nil, 2, resp.Status, nil
	}

	// Purse Response Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 3, "", err
	}

	// Purse Json
	var result *Response
	if err := json.Unmarshal([]byte(string(body)), &result); err != nil {
		return nil, 3, "", err
	}

	// Request Result
	if result.Msg == "unexpected" {
		return nil, 4, "", nil
	}

	return result, -1, "", nil
}
