package main

import (
	"fmt"

	"github.com/arxdsilva/pokemongorj/config"
	"github.com/dghubble/go-twitter/twitter"
)

func main() {
	creds := config.GetCredentials()
	c, _ := config.GetClient(&creds)
	params := &twitter.StreamFilterParams{
		Track:         []string{"pokemon go"},
		Language:      []string{"pt"},
		StallWarnings: twitter.Bool(true),
	}
	stream, _ := c.Streams.Filter(params)
	for message := range stream.Messages {
		t := message.(*twitter.Tweet)
		p := &twitter.StatusShowParams{TweetMode: "extended"}
		tweet, _, err := c.Statuses.Show(t.ID, p)
		fmt.Println(tweet.FullText, err)
	}
}
