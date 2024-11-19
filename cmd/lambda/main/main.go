package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"serverless-notifier/internal/pipeline"
)

func handler(ctx context.Context) (string, error) {
	err := pipeline.Run()
	if err != nil {
		return "Pipeline execution failed", err
	}
	return "Pipeline executed successfully", nil
}

func main() {
	lambda.Start(handler)
}
