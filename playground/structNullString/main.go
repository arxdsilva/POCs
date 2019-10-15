package main

import (
	"encoding/json"
	"fmt"
)

type a struct {
	Name *string `json:"name"`
}

func main() {
	aa := &a{}
	b, err := json.Marshal(aa)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
