package main

// Get the text to analyze
// Pass the text into the AWS Client
// Fetch the results of the analyzer
// Update the JSON file
// add in nlp property to the JSON.

// Convert the string to the new struc
// - whats currently in JSON string
// - plus nlp.

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/service/comprehend"
)

// Media - Used in new JSON
type Media struct {
	FeaturedMedia int
	WPMediaLink   []struct {
		Href string `json:"href"`
	}
}

// NLP - Inner JSON housing all Natural Launguage Processing Results
type NLP struct {
	Entities   comprehend.DetectEntitiesOutput
	KeyPhrases comprehend.DetectKeyPhrasesOutput
	Sentiment  comprehend.DetectSentimentOutput
}

// Payload - New payload created from the items we only care about.
type Payload struct {
	ID           int
	Date         string
	LastModified string
	Title        string
	Content      string
	Media        Media
	NLP          NLP
}

func main() {

	// Get the text to analyze
	testFileBucket := "jor-el-data-lake/data-raw-blogs/"
	testFileKey := "Parts of Speech Tagging N-Grams.json"

	testText := getS3Object(testFileKey, testFileBucket)
	body := testText.Body

	// Convert to string
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	string := buf.String()

	// Create the string to a struc
	blogJSON := Payload{}
	by := []byte(string)
	json.Unmarshal(by, &blogJSON)

	entityResults := getEntities(string)
	keyPhrasesResults := getKeyPhrases(string)
	sentimentResults := getSentiment(string)

	blogJSON.NLP = NLP{
		Sentiment:  *sentimentResults,
		KeyPhrases: *keyPhrasesResults,
		Entities:   *entityResults,
	}

	// Convert back to JSON
	var jsonData []byte
	jsonData, err := json.Marshal(blogJSON)

	str := fmt.Sprintf("%s", jsonData)
	fmt.Println(str)
	fmt.Println(err)

	// Bucket to save to
	newFilePath := fmt.Sprintf("%s.nlp", testFileKey)
	saveToS3(newFilePath, str)
}
