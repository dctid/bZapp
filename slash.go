package bZapp

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dctid/bZapp/model"
	"github.com/dctid/bZapp/view"
	"github.com/slack-go/slack"
	"log"
	"net/url"
	"os"
)

func Slash(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("bbbBody: %v", event.Body)

	body, err := url.ParseQuery(event.Body)
	if err != nil {
		log.Printf("Err parsing query: %v\n", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	triggerId := body["trigger_id"][0]
	modalRequest := view.NewSummaryModal(&model.Model{})
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
