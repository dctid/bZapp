package bZapp

import (
	"context"
	"encoding/json"
	"github.com/slack-go/slack"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Title struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji"`
}

type Modal struct {
	Type           string                `json:"type"`
	TriggerID      string                `json:"trigger_id,omitempty"`
	CallbackID     string                `json:"callback_id,omitempty"` // Required
	State          string                `json:"state,omitempty"`       // Optional
	Title          *Title                `json:"title"`
	SubmitLabel    string                `json:"submit_label,omitempty"`
	NotifyOnCancel bool                  `json:"notify_on_cancel,omitempty"`
	Elements       []slack.DialogElement `json:"elements,omitempty"`
}

func NewModal(title *Title) *Modal {
	return &Modal{
		Type:  "modal",
		Title: title,
	}
}

func Slash(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//var headers = map[string]string{
	//	"Content-Type": "application/json",
	//}
	//var block = slack.NewTextBlockObject("plain_text", "HIII", false, false)


	titleText := slack.NewTextBlockObject("plain_text", "bZapp", true, false)
	submitText := slack.NewTextBlockObject("plain_text", "Submit", true, false)
	closeText := slack.NewTextBlockObject("plain_text", "Cancel", true, false)
	todayHeader := slack.NewContextBlock("", slack.NewTextBlockObject("mrkdwn", "*Today's Events*", false, false))

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewDividerBlock(),
			todayHeader,
			slack.NewDividerBlock(),
			//headerSection,
			//firstName,
			//lastName,
		},
	}

	var modalRequest slack.ModalViewRequest
	modalRequest.Type = "modal"
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks

	jsonBytes, err := json.Marshal(modalRequest)
	log.Printf("json %s", jsonBytes)

	postHeaders := http.Header{"Content-Type": {"application/json"}}

	_, err = Post("http://localhost:8080/api/scores", postHeaders, modalRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	//defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//println(string(body))

	return events.APIGatewayProxyResponse{

		StatusCode: 200,
	}, nil
}
