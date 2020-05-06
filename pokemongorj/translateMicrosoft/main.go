package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/arxdsilva/pokemonGoRJ/config"
	"github.com/dghubble/go-twitter/twitter"
)

type MicosoftResponse []struct {
	DetectedLanguage struct {
		Language string  `json:"language"`
		Score    float64 `json:"score"`
	} `json:"detectedLanguage"`
	Translations []struct {
		Text string `json:"text"`
		To   string `json:"to"`
	} `json:"translations"`
}

func main() {
	creds := config.GetCredentials()
	c, _ := config.GetClient(&creds)
	ts := tweets(c)
	fmt.Println(0, ts[1].FullText)
	fmt.Println(0, TranslateMicroSoft(ts[1].FullText))
	// for i, t := range ts {
	// }
}

func tweets(c *twitter.Client) (ts []twitter.Tweet) {
	ts, _, _ = c.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: "pokemongoapp",
		TweetMode:  "extended",
	})
	return
}

func TranslateMicroSoft(text string) (t string) {
	bdStr := fmt.Sprintf(`[{"Text":"%s"}]`, text)
	bd := []byte(bdStr)
	req, _ := http.NewRequest("POST", "https://api.cognitive.microsofttranslator.com/translate", bytes.NewBuffer(bd))
	q := url.Values{}
	q.Add("api-version", "3.0")
	q.Add("to", "pt")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Ocp-Apim-Subscription-Key", os.Getenv("MICROSOFT_KEY"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", string(len(bdStr)))
	c := http.DefaultClient
	r, _ := c.Do(req)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}
	fmt.Println(string(body))
	yr := MicosoftResponse{}
	if err := json.Unmarshal(body, &yr); err != nil {
		panic(err)
	}
	if len(yr) > 0 {
		if len(yr[0].Translations) > 0 {
			t = yr[0].Translations[0].Text
		}
	}
	return
}
