package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/translate"
	"github.com/arxdsilva/pokemongorj/config"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/text/language"
)

func main() {
	creds := config.GetCredentials()
	c, _ := config.GetClient(&creds)
	client, err := NewClient(&Config{
		Addr:       "localhost:6379",
		Password:   "",
		DB:         0,
		MaxRetries: 3,
	})
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(client.cli.FlushAll().Err())
	ts := tweets(c)
	for i, t := range ts {
		if i == 0 || i == 2 {
			fmt.Println(i, " SET ", client.Set(1))
		}
		if i > 2 {
			continue
		}
		r, err := client.Find(int64(i))
		fmt.Println(i, " FIND ", r, t.ID, err)
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
