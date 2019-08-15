package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

var comprehendSession, _ = session.NewSession(&aws.Config{
	Region:      aws.String(awsRegion),
	Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
})

var comprehendClient = comprehend.New(comprehendSession)

// Get sentiment
func getSentiment(text string) *comprehend.DetectSentimentOutput {
	results, err := comprehendClient.DetectSentiment(&comprehend.DetectSentimentInput{
		LanguageCode: aws.String("en"),
		Text:         aws.String(text),
	})

	if err != nil {
		fmt.Println("Error", err)
	}

	return results
}

// Get Key Phrases
func getKeyPhrases(text string) *comprehend.DetectKeyPhrasesOutput {
	results, err := comprehendClient.DetectKeyPhrases(&comprehend.DetectKeyPhrasesInput{
		LanguageCode: aws.String("en"),
		Text:         aws.String(text),
	})

	if err != nil {
		fmt.Println("Error", err)
	}

	return results
}

// Get Entities
func getEntities(text string) *comprehend.DetectEntitiesOutput {
	results, err := comprehendClient.DetectEntities(&comprehend.DetectEntitiesInput{
		LanguageCode: aws.String("en"),
		Text:         aws.String(text),
	})

	if err != nil {
		fmt.Println("Error", err)
	}

	return results
}
