package bZapp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dctid/bZapp/model"
	"github.com/dctid/bZapp/view"
	"github.com/slack-go/slack"
	"log"
	"net/url"
	"os"
)

func Slash(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("bbbBody: %v", event.Body)

	command, err := slashCommandParse(event.Body)

	if err != nil {
		fmt.Printf("Could not parse slash command: %v\n", err)
		return events.APIGatewayProxyResponse{Headers: JsonHeaders(),
			Body:       "Bad payload",
			StatusCode: 400,
		}, err
	}

	triggerId := command.TriggerID
	log.Printf("Channel id: %s", command.ChannelID)
	table := os.Getenv("DYNAMODB_TABLE_NAME")

	withContext, err := DynamoDB.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Id": &dynamodb.AttributeValue{
				S: aws.String(command.ChannelID),
			},
			"token": &dynamodb.AttributeValue{
				S: aws.String("bs toekn"),
			},
			"username": &dynamodb.AttributeValue{
				S: aws.String("user"),
			},
		},
		TableName:                   aws.String("bZappTable"),
	})
	if err != nil {
		log.Printf("Couldn't save model %s", err)
	}
	log.Printf("Put resutl: %v", withContext)

	log.Printf("Table: %s", table)
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": &dynamodb.AttributeValue{
				S: aws.String(command.ChannelID),
			},
		},
		TableName: aws.String("bZappTable"),
	}

	currentModel, err := DynamoDB.GetItemWithContext(ctx, input)
	if err != nil {
		log.Printf("Couldn't get model %s", err)
	}
	log.Printf("model %+v", currentModel)

	modalRequest := view.NewSummaryModal(&model.Model{ChannelId: command.ChannelID})
	requestAsJson, _ := json.MarshalIndent(modalRequest, "", "\t")
	log.Printf("Body sent to slack to open modal: %v", string(requestAsJson))

	getenv := os.Getenv("SLACK_TOKEN")
	log.Printf("token: %s, trigger: %s", getenv, triggerId)
	api := slack.New(getenv, slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.OpenView(triggerId, modalRequest)

	statusCode := 200
	if err != nil {
		statusCode = 500
		log.Printf("Err opening view: %v", err)
		indent, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success opening bZap modal: %v", string(indent))
	} else {
		indent, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success opening bZap modal: %v", string(indent))
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
	}, nil
}

func slashCommandParse(bodyStr string) (slack.SlashCommand, error) {
	body, err := url.ParseQuery(bodyStr)
	if err != nil {
		return slack.SlashCommand{}, err
	}

	return slack.SlashCommand{
		Token:          body.Get("token"),
		TeamID:         body.Get("team_id"),
		TeamDomain:     body.Get("team_domain"),
		EnterpriseID:   body.Get("enterprise_id"),
		EnterpriseName: body.Get("enterprise_name"),
		ChannelID:      body.Get("channel_id"),
		ChannelName:    body.Get("channel_name"),
		UserID:         body.Get("user_id"),
		UserName:       body.Get("user_name"),
		Command:        body.Get("command"),
		Text:           body.Get("text"),
		ResponseURL:    body.Get("response_url"),
		TriggerID:      body.Get("trigger_id"),
	}, nil
}
