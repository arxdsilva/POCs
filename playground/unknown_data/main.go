package main

// playground
// https://play.golang.org/p/dPmHZkw2QwB

import (
	"encoding/json"
	"fmt"
)

type emptyData struct{}
type emptyData2 interface{}

type fullJson struct {
	Name  string     `json:"name"`
	Data  emptyData  `json:"data"`
	Data2 emptyData2 `json:"data2"`
}

func main() {
	data := `{"data":{"abc":1},"data2":{"abc":1},"name":"arthur"}`
	fj := &fullJson{}
	err := json.Unmarshal([]byte(data), fj)
	if err != nil {
		fmt.Println("Hello, err: ", err)
	}
	fmt.Printf("Hello, data %+v\n", fj)
	m, err := json.Marshal(fj.Data2)
	if err != nil {
		fmt.Println("Hello, err: ", err)
	}
	fmt.Printf("Hello, marshal %v\n", string(m))
}

// Hello, data &{Name:arthur Data:{} Data2:map[abc:1]}
// Hello, marshal {"abc":1}
// Program exited.
