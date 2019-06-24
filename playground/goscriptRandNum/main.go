package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/matryer/goscript"
)

func main() {
	c, err := config()
	if err != nil {
		log.Fatalln(err)
	}
	script := goscript.New(c.FuncString)
	defer script.Close()
	ex, err := script.Execute()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ex)
}

type Config struct {
	FuncString string `json:"func"`
}

func config() (c Config, err error) {
	raw, err := ioutil.ReadFile("config.json")
	if err != nil {
		return
	}
	return c, json.Unmarshal(raw, &c)
}
