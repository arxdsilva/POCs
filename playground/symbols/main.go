package main

import (
	"encoding/json"
	"fmt"
)

type SymbolsClientResp2 struct {
	Success bool `json:"success"`
	Symbols map[string]struct {
		Description string `json:"description"`
		Code        string `json:"code"`
	} `json:"symbols"`
}

var ref = `{
	"motd": {
		"msg": "If you or your company use this project or like what we doing, please consider backing us so we can continue maintaining and evolving this project.",
		"url": "https://exchangerate.host/#/donate"
	},
	"success": true,
	"symbols": {
		"AED": {
			"description": "United Arab Emirates Dirham",
			"code": "AED"
		}
	}
}`

func main() {
	s := &SymbolsClientResp2{}
	err := json.Unmarshal([]byte(ref), s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", s)
}
