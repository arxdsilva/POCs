package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("Blank.pdf")
	if err != nil {
		log.Fatal(err)
	}
	var str string
	for _, b := range b {
		str = str + fmt.Sprintf("%v,\n", b)
	}
	err = ioutil.WriteFile("out", []byte(str), 0700)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(str)
}
