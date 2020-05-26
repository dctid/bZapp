package bZapp

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

// Dashboard returns a dashboard HTML page
func Dashboard(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Key: %s", os.Getenv("SLACK_KEY"))
	return events.APIGatewayProxyResponse{
		Body: string("<html><body><h1>bZapp dashboard</h1></body></html>\n"),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
		StatusCode: 200,
	}, nil
}
