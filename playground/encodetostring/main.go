package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println(basicAuthString("5f45115cb9e7", "920b8f18"))
}

func basicAuthString(u, p string) string {
	up := u + ":" + p
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(up))
}
