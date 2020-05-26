package bZapp

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/slack-go/slack"
)

func Slash(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var block = slack.NewTextBlockObject("plain_text", "HIII", false, false)
	b, err := json.MarshalIndent(block, "", "    ")
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	resp, err := GET("https://www.google.com", nil)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)

	var headers = map[string]string{
		"Content-Type": "application/json",
	}
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		Body:       string(b),
		StatusCode: 200,
	}, nil
}
