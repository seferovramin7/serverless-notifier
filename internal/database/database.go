package database

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"serverless-notifier/internal/fetcher"
)

var dynamoClient *dynamodb.Client
var tableName = "linkedin_jobs"

// Initialize the DynamoDB client
func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}
	dynamoClient = dynamodb.NewFromConfig(cfg)
}

// CheckDatabase checks if the provided jobs already exist in the DynamoDB table
func CheckDatabase(jobs []fetcher.Job) ([]fetcher.Job, error) {
	var newJobs []fetcher.Job

	for _, job := range jobs {
		input := &dynamodb.GetItemInput{
			TableName: aws.String(tableName),
			Key: map[string]types.AttributeValue{
				"JobID": &types.AttributeValueMemberS{Value: job.Response.Jobs[0].JobPostingURL},
			},
		}

		result, err := dynamoClient.GetItem(context.TODO(), input)
		if err != nil {
			return nil, err
		}

		// If the item doesn't exist, add it to the new jobs list
		if result.Item == nil {
			newJobs = append(newJobs, job)
		}
	}

	return newJobs, nil
}

// StoreJobIDs stores the job IDs in the DynamoDB table
func StoreJobIDs(jobs []fetcher.Job) error {
	for _, job := range jobs {
		input := &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item: map[string]types.AttributeValue{
				"JobID":   &types.AttributeValueMemberS{Value: job.Response.Jobs[0].JobPostingURL},
				"JobData": &types.AttributeValueMemberS{Value: job.Response.Jobs[0].JobPostingURL}, // Adjust this based on your Job structure
			},
		}

		_, err := dynamoClient.PutItem(context.TODO(), input)
		if err != nil {
			return err
		}
	}

	return nil
}
