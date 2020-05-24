package gofaas

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func Slash(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var headers = map[string]string{
		"Content-Type": "application/json",
	}
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		Body:       fmt.Sprintf("{\"text\": \"%v\"}", "hi"),
		StatusCode: 200,
	}, nil
}
