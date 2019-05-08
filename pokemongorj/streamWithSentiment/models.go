package main

type MicrosftSentimentResponse struct {
	// LanguageDetection struct {
	// 	Documents []struct {
	// 		ID                string `json:"id"`
	// 		DetectedLanguages []struct {
	// 			Name        string  `json:"name"`
	// 			Iso6391Name string  `json:"iso6391Name"`
	// 			Score       float64 `json:"score"`
	// 		} `json:"detectedLanguages"`
	// 	} `json:"documents"`
	// 	Errors []interface{} `json:"errors"`
	// } `json:"languageDetection"`
	// KeyPhrases struct {
	// 	Documents []struct {
	// 		ID         string   `json:"id"`
	// 		KeyPhrases []string `json:"keyPhrases"`
	// 	} `json:"documents"`
	// 	Errors []interface{} `json:"errors"`
	// } `json:"keyPhrases"`
	Sentiment struct {
		Documents []struct {
			ID    string  `json:"id"`
			Score float64 `json:"score"`
		} `json:"documents"`
		Errors []interface{} `json:"errors"`
	} `json:"sentiment"`
	// Entities struct {
	// 	Documents []struct {
	// 		ID       string        `json:"id"`
	// 		Entities []interface{} `json:"entities"`
	// 	} `json:"documents"`
	// 	Errors []interface{} `json:"errors"`
	// } `json:"entities"`
}
