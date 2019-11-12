package main

import (
	"fmt"

	"github.com/arxdsilva/webhook"
)

func main() {
	secret := []byte("f0247600-0574-11ea-8d71-362b9e155667")
	body := []byte("a message")
	sig := webhook.SignBody(secret, body)
	fmt.Printf("signature: %x", sig)
}
