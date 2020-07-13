package main

import (
	"context"
	"fmt"
	"os"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

func translateText(projectID string, sourceLang string, targetLang string, text string) error {
	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return fmt.Errorf("NewTranslationClient: %v", err)
	}
	defer client.Close()
	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", projectID),
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
		MimeType:           "text/plain",
		Contents:           []string{text},
	}
	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return fmt.Errorf("TranslateText: %v", err)
	}
	for _, translation := range resp.GetTranslations() {
		fmt.Printf("Translated text: %v\n", translation.GetTranslatedText())
	}
	return nil
}

func main() {
	id := os.Getenv("GOOGLE_ID")
	text := "When will these Team GO Rocket Grunts learn that nothing is stronger than a united team of Pokémon GO Trainers?! I guess we’ll just have to remind them! Bíceps flexionados Friendly reminder to do your best to defeat a #TeamGORocket Grunt today!"
	err := translateText(id, "en", "pt-br", text)
	fmt.Println("err: ", err)
}
