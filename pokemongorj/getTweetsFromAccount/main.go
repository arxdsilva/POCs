package main

import (
	"fmt"

	"github.com/arxdsilva/pokemongorj/config"
	"github.com/dghubble/go-twitter/twitter"
)

func main() {
	creds := config.GetCredentials()
	c, _ := config.GetClient(&creds)
	ts := tweets(c)
	for i, t := range ts {
		fmt.Println(i, t.FullText)
	}
}

func tweets(c *twitter.Client) (ts []twitter.Tweet) {
	ts, _, _ = c.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: "pokemongoapp",
		TweetMode:  "extended",
	})
	return
}
