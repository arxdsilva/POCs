package main

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/arxdsilva/pocs/playground/docker-volume/config"
)

func main() {
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatal("FromEnv ", err)
	}
	data, err := ioutil.ReadFile(cfg.PrivateKeyFile)
	if err != nil {
		log.Fatal("ReadFile ", err)
	}
	keyData, _ := pem.Decode(data)
	if keyData == nil {
		log.Fatal("Decode: no key data")
	}
	fmt.Println(string(data))
}
