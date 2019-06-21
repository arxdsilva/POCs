package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var t struct {
	Map  map[string][]string `json:"map"`
	Name string              `json:"name"`
}

func main() {
	m := make(map[string][]string)
	m["pikachu"] = []string{"raichu"}
	t.Map = m
	t.Name = "ash"
	b, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
