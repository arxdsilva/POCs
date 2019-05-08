package main

import (
	"fmt"

	"github.com/arxdsilva/pokemongorj/config"
	"github.com/dghubble/go-twitter/twitter"
)

func main() {
	creds := config.GetCredentials()
	c, _ := config.GetClient(&creds)
	streamDemux(c)
	// params := &twitter.StreamFilterParams{
	// 	Track:         []string{"pokemon go"},
	// 	Language:      []string{"pt"},
	// 	StallWarnings: twitter.Bool(true),
	// }
	// stream, _ := c.Streams.Filter(params)
	// likeStream(stream, c)
}

func likeStream(s *twitter.Stream, c *twitter.Client) {
	for message := range s.Messages {
		t := message.(*twitter.Tweet)
		p := &twitter.StatusShowParams{TweetMode: "extended"}
		tweet, _, err := c.Statuses.Show(t.ID, p)
		fmt.Println(tweet.FullText, "\n")
		tt, _, err := c.Favorites.Create(&twitter.FavoriteCreateParams{ID: t.ID})
		fmt.Println(tt.ID, err, "\n")
	}
}

func streamDemux(c *twitter.Client) {
	params := &twitter.StreamFilterParams{
		Track:         []string{"pokemon go"},
		Language:      []string{"en"},
		StallWarnings: twitter.Bool(true),
	}
	s, err := c.Streams.Filter(params)
	fmt.Println("Starting Stream...", err)
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		p := &twitter.StatusShowParams{TweetMode: "extended"}
		tweet, _, err := c.Statuses.Show(tweet.ID, p)
		fmt.Println(tweet.FullText, "\n")
		tt, _, err := c.Favorites.Create(&twitter.FavoriteCreateParams{ID: tweet.ID})
		fmt.Println(tt.ID, err, "\n")
	}
	demux.HandleChan(s.Messages)
	fmt.Println("Stopping Stream...")
	s.Stop()
	streamDemux(c)
}
