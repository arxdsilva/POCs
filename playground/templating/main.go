package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
)

type Animal struct {
	Name string
	Data map[string]interface{}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	t, err := template.New("order.html").ParseFiles(wd + "/tmp/order.html")
	if err != nil {
		return
	}
	b := new(bytes.Buffer)
	err = t.Execute(b, nil)
	if err != nil {
		return
	}
	fmt.Println(b.String())
	return
}
