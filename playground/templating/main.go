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

func (a *Animal) Eat(food string) string {
	return os.Getenv(food)
}

func (a *Animal) Seeds() string {
	return "asidaisdadkapdadkao"
}

func main() {
	a := Animal{Name: "Cacatua", Data: map[string]interface{}{"a": "b"}}
	if err := run(a); err != nil {
		log.Fatal(err)
	}
}

func run(a Animal) (err error) {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	t, err := template.New("order.html").ParseFiles(wd + "/tmp/order.html")
	if err != nil {
		return
	}
	b := new(bytes.Buffer)
	err = t.Execute(b, &a)
	if err != nil {
		return
	}
	fmt.Println(b.String())
	return
}
