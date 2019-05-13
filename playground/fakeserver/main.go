package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "some err", http.StatusBadRequest)
	})
	http.ListenAndServe(":3333", nil)
}
