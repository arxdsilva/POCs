package main

import (
	"fmt"
	"time"
)

func main() {
	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Now().In(location))
}
