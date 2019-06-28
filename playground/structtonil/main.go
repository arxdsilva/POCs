package main

import (
	"fmt"
	"net/http"
)

func main() {
	var c *http.Client
	fmt.Println(c == nil)
	fmt.Println((&http.Client{}) == c)
}
