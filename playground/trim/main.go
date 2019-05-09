package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "2 dia(s) 2 hora(s) 3 minuto(s)"
	fmt.Println(strings.Trim(s, " "))
	fmt.Println(strings.TrimSpace(s))
}
