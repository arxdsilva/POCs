package main

// https://play.golang.org/p/80A1qpHIvW0

import (
	"fmt"
	"time"
)

// dont forget to export TZ='America/Sao_Paulo'
// before running `go run main.go`
func main() {
	fmt.Println(time.Now())
}
