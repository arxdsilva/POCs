package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(fmt.Sprintf("%v %s", http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable)))
}
