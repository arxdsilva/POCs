package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "some err", http.StatusBadRequest)
	})
	http.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received: ", r.Method)
		time.Sleep(3 * time.Second)
		http.Error(w, "some err", http.StatusBadRequest)
	})
	http.ListenAndServe(":3333", nil)
}
