package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	ns := News("pokemon go")
	for i, v := range ns.Value {
		fmt.Println()
		fmt.Printf("%v.%v	|	%v\n", i, v.Name, v.URL)
	}
}

func News(subj string) (ans NewsAnswer) {
	req, _ := http.NewRequest("GET", "https://api.cognitive.microsoft.com/bing/v7.0/news/search", nil)
	q := url.Values{}
	q.Add("mkt", "pt-br")
	q.Add("q", subj)
	q.Add("freshness", "day")
	req.Header.Add("Ocp-Apim-Subscription-Key", os.Getenv("BING_KEY"))
	req.URL.RawQuery = q.Encode()
	c := http.DefaultClient
	r, _ := c.Do(req)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}
	ans = NewsAnswer{}
	if err := json.Unmarshal(body, &ans); err != nil {
		panic(err)
	}
	return
}

type NewsAnswer struct {
	ReadLink     string `json: "readLink"`
	QueryContext struct {
		OriginalQuery string `json: "originalQuery"`
		AdultIntent   bool   `json: "adultIntent"`
	} `json: "queryContext"`
	TotalEstimatedMatches int `json: totalEstimatedMatches"`
	Sort                  []struct {
		Name       string `json: "name"`
		ID         string `json: "id"`
		IsSelected bool   `json: "isSelected"`
		URL        string `json: "url"`
	} `json: "sort"`
	Value []struct {
		Name  string `json: "name"`
		URL   string `json: "url"`
		Image struct {
			Thumbnail struct {
				ContentUrl string `json: "thumbnail"`
				Width      int    `json: "width"`
				Height     int    `json: "height"`
			} `json: "thumbnail"`
			Description string `json: "description"`
			Provider    []struct {
				Type string `json: "_type"`
				Name string `json: "name"`
			} `json: "provider"`
			DatePublished string `json: "datePublished"`
		} `json: "image"`
	} `json: "value"`
}
