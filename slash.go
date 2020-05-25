package bZapp

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func Slash(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	resp, err := http.Get("https://www.google.com")
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var headers = map[string]string{
		"Content-Type": "application/json",
	}
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		Body:       fmt.Sprintf("{\"text\": \"%v\"}", body),
		StatusCode: 200,
	}, nil
}
