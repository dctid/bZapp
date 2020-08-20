package bZapp

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

func VerifyRequestInterceptor(h HandlerAPIGateway) HandlerAPIGateway {
	return func(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		if signed := verifySigning(e.Body, e.Headers); signed {
			r, err := h(ctx, e)
			return r, err
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 401,
		}, errors.New("unauthorized")
	}
}

func verifySigning(body string, header map[string]string) bool {
	headerMap := convertHeaders(header)
	signingSecret, isSet := os.LookupEnv("SLACK_SIGNING_SECRET")
	log.Printf("Slack signing secret set: %v", isSet)
	if isSet {
		sv, err := slack.NewSecretsVerifier(headerMap, signingSecret)
		if err != nil {
			log.Printf("[ERROR] Fail to verify SigningSecret: %v", err)
			return false
		} else {
			log.Printf("Created signer %v", sv)
		}
		sv.Write([]byte(body))
		if err := sv.Ensure(); err != nil {
			log.Printf("[ERROR] Fail to verify SigningSecret: %v", err)
			return false
		} else {
			log.Println("Passed verify")
		}

	} else {
		log.Println("Signing secret not set, returning true")
	}
	return true
}

func convertHeaders(headers map[string]string) http.Header {
	result := http.Header{}
	for key, value := range headers {
		result.Add(key, value)
	}
	return result
}
