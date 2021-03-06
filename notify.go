package bZapp

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

// HandlerAPIGateway is an API Gateway Proxy Request handler function
type HandlerAPIGateway func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

// HandlerCloudWatch is a CloudWatchEvent handler function
type HandlerCloudWatch func(context.Context, events.CloudWatchEvent) error

// NotifyAPIGateway wraps a handler func and sends an SNS notification on error
func NotifyAPIGateway(h HandlerAPIGateway) HandlerAPIGateway {
	return func(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		r, err := h(ctx, e)
		notify(ctx, err)
		return r, err
	}
}

func notify(ctx context.Context, err error) {
	if err == nil {
		return
	}

	subj := fmt.Sprintf("ERROR %s", os.Getenv("AWS_LAMBDA_FUNCTION_NAME"))
	msg := fmt.Sprintf("%+v\n", err)
	log.Printf("%s %s\n", subj, msg)

	topic, topicSet := os.LookupEnv("NOTIFICATION_TOPIC")
	if topicSet {
		return
	}

	_, err = SNS.PublishWithContext(ctx, &sns.PublishInput{
		Message:  aws.String(msg),
		Subject:  aws.String(subj),
		TopicArn: aws.String(topic),
	})
	if err != nil {
		log.Printf("NotifyError SNS Publish error %+v\n", err)
	}
}
