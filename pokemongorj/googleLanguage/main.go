package main

import (
	"context"
	"fmt"
	"os"

	language "cloud.google.com/go/language/apiv1"
	"google.golang.org/api/option"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

func main() {
	keys := `{}`
	os.Setenv("GOOGLE_KEYS", keys)
	lang()
}

func lang() {
	ctx := context.Background()
	client, err := language.NewClient(ctx, option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_KEYS"))))
	fmt.Println(err)
	fmt.Printf("%+v", client)
	r, err := analyzeSentiment(ctx, client, "ola mundo!")
	fmt.Println(err)
	fmt.Printf("%+v", r)
	if r.DocumentSentiment.Score >= 0 {
		fmt.Println("Sentiment: positive")
	} else {
		fmt.Println("Sentiment: negative")
	}
}

func analyzeSentiment(ctx context.Context, client *language.Client, text string) (*languagepb.AnalyzeSentimentResponse, error) {
	return client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
	})
}
