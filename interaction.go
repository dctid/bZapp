package bZapp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dctid/bZapp/modal"
	"github.com/dctid/bZapp/test"
	"github.com/slack-go/slack"
	"log"
	"net/url"
	"os"
)

func Interaction(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Printf("Body: %v", event.Body)

	m, err := url.ParseQuery(event.Body)
	if err != nil {
		log.Printf("Err parsing query: %v\n", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	} else {
		log.Printf("Body parsed: %v\n", m)
	}
	var payload slack.InteractionCallback
	err = json.Unmarshal([]byte(m["payload"][0]), &payload)
	if err != nil {
		fmt.Printf("Could not parse action response JSON: %v\n", err)
		return events.APIGatewayProxyResponse{Headers: JsonHeaders(),
			Body:       "Bad payload",
			StatusCode: 400,
		}, err
	}
	switch payload.Type {
	case slack.InteractionTypeViewSubmission:
		return pushModalWithAddedEvent(payload)
	case slack.InteractionTypeBlockActions:
		return actionEvent(payload)
	}

	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       fmt.Sprintf("Unimplemented Event Type: %v", payload.Type),
		StatusCode: 400,
	}, nil

}

func actionEvent(payload slack.InteractionCallback) (events.APIGatewayProxyResponse, error) {
	log.Printf("action id %s\n", payload.ActionCallback.BlockActions[0].ActionID)
	switch payload.ActionCallback.BlockActions[0].ActionID {
	case modal.EditEventsActionId:
		return pushEditEventModal(payload)
	case modal.RemoveEventActionId:
		return removeEvent(payload)
	}
	return 	 events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       "Unknown action type",
		StatusCode: 400,
	}, nil
}

func removeEvent(payload slack.InteractionCallback) (events.APIGatewayProxyResponse, error) {
	log.Printf("remove startedddddddd	sss")
	//actionId := payload.ActionCallback.BlockActions[0].ActionID
	actionValue := payload.ActionCallback.BlockActions[0].Value
	todaysSectionBlocks, tomorrowsSectionBlocks := modal.RemoveEvent(payload.View.Blocks.BlockSet, actionValue)
	modalRequest := modal.NewEditEventsModal(todaysSectionBlocks, tomorrowsSectionBlocks)
	//update := slack.NewUpdateViewSubmissionResponse(&modalRequest)

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.UpdateView(modalRequest, payload.View.ExternalID, payload.Hash, payload.View.ID)
	if err != nil {
		log.Printf("Err opening modal: %v\n", err)
	} else {
		log.Printf("Success open modal: %v\n", viewResponse)
		indent, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success open modal2: %v", string(indent))
	}
	update := slack.NewUpdateViewSubmissionResponse(&modalRequest)
	jsonBytes, err := json.Marshal(update)
	log.Printf("Json bytes: %v\n", jsonBytes)
	//jsonBytes, err := json.Marshal(update)
	//indent, _ := json.MarshalIndent(update, "", "\t")
	//log.Printf("body sent to slack: %v", string(indent))

	return events.APIGatewayProxyResponse{
		Headers: JsonHeaders(),
		//Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func pushModalWithAddedEvent(payload slack.InteractionCallback) (events.APIGatewayProxyResponse, error) {
	action := payload.View.State.Values[modal.AddEventDayInputBlock][modal.AddEventDayActionId]
	marshal, _ := json.Marshal(action)
	fmt.Printf("Add Event button pressed by user %s with value %v\n", payload.User.Name, string(marshal))

	eventDay, newEvent := modal.BuildNewEventSectionBlock(payload.View.State.Values)

	todaysSectionBlocks, tomorrowsSectionBlocks := modal.AddNewEventToDay(payload.View.Blocks.BlockSet, eventDay, newEvent)
	fmt.Printf("Addedss New got: %v, got1: %v\n", len(todaysSectionBlocks), len(tomorrowsSectionBlocks))
	modalRequest := modal.NewSummaryModal(todaysSectionBlocks, tomorrowsSectionBlocks)
	//modalRequest.PrivateMetadata = "test metadata"

	//api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true))
	//viewResponse, err := api.OpenView(payload.TriggerID, modalRequest)
	////interaction, err := api.UpdateView(modalRequest, payload.View.ExternalID, payload.Hash, payload.View.ID)
	//if err != nil {
	//	log.Printf("Err opening modal: %v\n", err)
	//} else {
	//	log.Printf("Success open modal: %v\n", viewResponse)
	//	indent, _ := json.MarshalIndent(viewResponse, "", "\t")
	//	log.Printf("Success open modal2: %v", string(indent))
	//}
	update := slack.NewUpdateViewSubmissionResponse(&modalRequest)
	jsonBytes, err := json.Marshal(update)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers:    JsonHeaders(),
			Body:       "Error processing request",
			StatusCode: 500,
		}, err
	}
	indent, _ := json.MarshalIndent(update, "", "\t")
	log.Printf("body sent to slack: %v", string(indent))

	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func pushEditEventModal(payload slack.InteractionCallback) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Message button pressed by user %s with value %v\n", payload.User.Name, payload)

	todaysEvents, tomorrowsEvents := modal.ExtractEvents(payload.View.Blocks.BlockSet)
	todaysEvents, tomorrowsEvents = modal.ConvertToEventsWithRemoveButton(todaysEvents, tomorrowsEvents)
	todaysEvents, tomorrowsEvents = modal.ReplaceEmptyEventsWithNoEventsYet(todaysEvents, tomorrowsEvents)

	modalRequest := modal.NewEditEventsModal(todaysEvents, tomorrowsEvents)
	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.UpdateView(modalRequest, payload.View.ExternalID, payload.Hash, payload.View.ID)
	if err != nil {
		log.Printf("Err opening modal: %v\n", err)
	} else {
		log.Printf("Success open modal: %v\n", viewResponse)
		indent, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success open modal2: %v", string(indent))
	}
	update := slack.NewUpdateViewSubmissionResponse(&modalRequest)
	jsonBytes, err := json.Marshal(update)
	log.Printf("Json bytes: %v\n", test.PrettyJsonNoError(string(jsonBytes)))

	return events.APIGatewayProxyResponse{
		Headers: JsonHeaders(),
		StatusCode: 200,
	}, nil
}
