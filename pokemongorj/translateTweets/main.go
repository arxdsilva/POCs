package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"cloud.google.com/go/translate"
	"github.com/arxdsilva/pokemongorj/config"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/text/language"
)

type YandexResponse struct {
	Code int      `json:"code"`
	Lang string   `json:"lang"`
	Text []string `json:"text"`
}

func main() {
	creds := config.GetCredentials()
	c, _ := config.GetClient(&creds)
	ts := tweets(c)
	for i, t := range ts {
		fmt.Println(i, t.FullText)
		fmt.Println(i, TranslateYandex(t.FullText))
	}
}

func tweets(c *twitter.Client) (ts []twitter.Tweet) {
	ts, _, _ = c.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: "pokemongoapp",
		TweetMode:  "extended",
	})
	return
}

func trslt(msg string) (t string) {
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	target, err := language.Parse("pt")
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}
	translations, err := client.Translate(ctx, []string{msg}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}
	t = translations[0].Text
	return
}

func TranslateYandex(text string) (t string) {
	apiKey := os.Getenv("YANDEX_KEY")
	req, _ := http.NewRequest("GET", "https://translate.yandex.net/api/v1.5/tr.json/translate", nil)
	q := url.Values{}
	q.Add("lang", "pt")
	q.Add("key", apiKey)
	q.Add("format", "plain")
	q.Add("text", text)
	req.URL.RawQuery = q.Encode()
	c := http.DefaultClient
	r, _ := c.Do(req)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}
	fmt.Println(string(body))
	yr := YandexResponse{}
	if err := json.Unmarshal(body, &yr); err != nil {
		panic(err)
	}
	if len(yr.Text) > 0 {
		t = yr.Text[0]
	}
	return
}
