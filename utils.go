package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func getFlag(k string, kf string, s []string) (string, int) {
	for i, v := range s {
		if k == v {
			return s[i+1], i
		}
	}
	for i, v := range s {
		if kf == v {
			return s[i+1], i
		}
	}
	return "", 0
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func isset(k int, s []string) bool {
	return len(s) >= k+1
}

type config struct {
	Api string `json:"api"`
}

func loadConfig(dev bool) (*config, error) {
	var configfp string
	if !dev {
		runp, err := os.Executable()
		if err != nil {
			return nil, err
		}
		rund := filepath.Dir(runp)
		configfp = rund + "/config.json"
	} else {
		configfp = "./config.json"
	}
	//configfp = "" // カスタムパスで上書きする場合
	f, err := os.Open(configfp)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Text string `json:"text"`
}

func HttpRequest(api_url string) (*Response, error) {
	var response *Response
	resp, err := http.Get(api_url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, err
	}

	body, _ := io.ReadAll(resp.Body)

	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}
	return response, nil
}

func urlGen(to string, from string, text string, base string) string {
	var query string
	if from == "" {
		query = "?to=" + to + "&text=" + url.QueryEscape(text)
	} else {
		query = "?to=" + to + "&from=" + from + "&text=" + url.QueryEscape(text)
	}
	return base + query
}
