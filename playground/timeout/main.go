package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "http://localhost:3333/timeout"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println(1, err)
		return
	}
	client := &http.Client{Timeout: time.Second}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(2, err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(3, err)
		return
	}
	fmt.Println(err)
	fmt.Println(string(body))
}
