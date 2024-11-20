package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"serverless-notifier/internal/pipeline"
)

type Event struct {
	Query       string `json:"query"`
	Page        int    `json:"page"`
	LocationID  string `json:"locationId"`
	SortBy      string `json:"sortBy"`
	FunctionIDs string `json:"functionIds"`
	PostedAgo   int    `json:"postedAgo"`
	Experience  int    `json:"experience"`
}

func handler(ctx context.Context, event Event) (string, error) {
	err := pipeline.Run(event.Query, event.Page, event.LocationID, event.SortBy, event.FunctionIDs, event.PostedAgo, event.Experience)

	if err != nil {
		return "Pipeline execution failed", err
	}
	return "Pipeline executed successfully", nil
}

func main() {
	lambda.Start(handler)
}
