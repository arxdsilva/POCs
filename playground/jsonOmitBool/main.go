package main

// https://play.golang.org/p/dLXFKsUJRFD

import (
	"encoding/json"
	"fmt"
)

type A struct {
	B bool `json:"b,omitempty"`
	C bool `json:"c"`
}

func main() {
	a := A{}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}
