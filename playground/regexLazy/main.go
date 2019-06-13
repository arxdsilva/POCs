package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://play.golang.org/p/RrLP3Q6LSi3
func main() {
	r, _ := http.Get("https://pastebin.com/raw/er9Rmr2T")
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close() //  must close
	fmt.Println(string(bodyBytes))
}
