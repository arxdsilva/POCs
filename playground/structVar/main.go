package main

import "fmt"

func main() {
	var fee struct {
		FeeVal  int    `json:"fee_val"`
		FeeName string `json:"fee_name"`
	}
	fmt.Println(fee)
}
