package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := `https://newsapi.org/v2/everything?q=League+of+legends&language=en&sortBy=relevancy&apiKey=%s&from=2019-06-29`
	url = fmt.Sprintf(url, os.Getenv("KEY"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(bodyBytes))
}
