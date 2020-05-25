package bZapp

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-lambda-go/events"
)

func TestSlash(t *testing.T) {

	r, err := Slash(context.Background(), events.APIGatewayProxyRequest{
		Body: `{"text": "none"}`,
	})
	assert.NoError(t, err)

	assert.EqualValues(t,
		events.APIGatewayProxyResponse{
			Body: "{\"text\": \"hi\"}",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			StatusCode: 200,
		},
		r,
	)
}
