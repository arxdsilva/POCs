package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/arxdsilva/webhook"
)

func main() {
	secret := []byte("f0247600-0574-11ea-8d71-362b9e155667")
	body := "a message"
	sig := webhook.SignBody(secret, []byte(body))
	signature := fmt.Sprintf("sha1=%x", sig)
	req, _ := http.NewRequest("POST", "", strings.NewReader(body))
	req.Header.Add("x-hub-signature", signature)
	h, err := webhook.Parse(secret, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("signed: ", h.SignedBy(secret))
	fmt.Println("Signature: ", h.Signature)
}
