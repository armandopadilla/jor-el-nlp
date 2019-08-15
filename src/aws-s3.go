package main

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Set up the session and client for use.
var s3Session, _ = session.NewSession(&aws.Config{
	Region:      aws.String(awsRegion),
	Credentials: credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, ""),
})

var s3Client = s3.New(s3Session)

/**
 * Fetch the posts and save to SQS
 **/
func saveToS3(key string, payload string) {
	buffer := []byte(payload)

	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("jor-el-data-lake/data-raw-blogs/"),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buffer),
	})

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	return
}

/**
 * Get the s3 object
 **/
func getS3Object(key string, bucket string) *s3.GetObjectOutput {
	results, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		fmt.Println("Error", err)
	}

	return results

}
