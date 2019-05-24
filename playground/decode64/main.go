package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	s := "UAHSUAHSUAHSUAHSUAH"
	sDec, _ := b64.StdEncoding.DecodeString(s)
	fmt.Println(string(sDec))
}
