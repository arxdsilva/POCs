package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	s := Sentiment("Hoje nao vai ter mais brincadeira.")
	fmt.Printf("S: %+v", s)
}

func Sentiment(text string) (ans ST) {
	uriBase := "https://eastus.api.cognitive.microsoft.com/text/analytics/v2.1/sentiment"
	document := fmt.Sprintf("[{\"id\":0, \"text\":\"%s\"}]", text)
	rBody := strings.NewReader("{\"documents\": " + document + "}")
	req, err := http.NewRequest("POST", uriBase, rBody)
	if err != nil {
		log.Printf("Req error body: %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", os.Getenv("SENTIMENT_KEY"))
	c := &http.Client{Timeout: time.Second * 2}
	r, err := c.Do(req)
	if err != nil {
		log.Printf("Req error body: %v", err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}
	if r.StatusCode != 200 {
		fmt.Println("Wrong status code: ", r.StatusCode)
		fmt.Println("body: ", string(body))
		return
	}
	fmt.Println("body: ", string(body))
	ans = ST{}
	if err := json.Unmarshal(body, &ans); err != nil {
		log.Printf("Error Unmarshal body: %v", err)
		return
	}
	return
}

type ST struct {
	Documents []struct {
		ID    string  `json:"id"`
		Score float64 `json:"score"`
	} `json:"documents"`
	Errors []interface{} `json:"errors"`
}
