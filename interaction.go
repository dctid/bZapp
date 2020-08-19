package bZapp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/modal"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
	"io/ioutil"
	"log"
	"net/http"
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
	var payload modal.InteractionPayload
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
		return viewSubmission(payload)
	case slack.InteractionTypeBlockActions:
		return actionEvent(payload)
	case slack.InteractionTypeViewClosed:
		return viewClosed(payload)
	}

	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       fmt.Sprintf("Unimplemented Event Type: %v", payload.Type),
		StatusCode: 400,
	}, nil

}


func viewSubmission(payload modal.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	if len(payload.View.State.Values) == 1 && payload.View.State.Values["convo_input_id"]["conversation_select_action_id"].Type == "conversations_select" {
		return publishbZapp(payload)
	} else {
		return pushModalWithAddedEvent(payload)
	}
}

func publishbZapp(payload modal.InteractionPayload) (events.APIGatewayProxyResponse, error) {

	url := payload.ResponseUrls[0].ResponseUrl
	log.Printf("Response Urls: %s", url)
	headers := http.Header{"Content-type": []string{"application/json"}}

	todaysEvents, tomorrowsEvents := modal.ExtractEvents2(payload.View.Blocks.BlockSet)
	todaysSectionBlocks, tomorrowsSectionBlocks := modal.ConvertToEventsWithoutRemoveButton2(todaysEvents, tomorrowsEvents)
	tomorrowsSectionBlocks, tomorrowsSectionBlocks = modal.ReplaceEmptyEventsWithNoEventsYet(todaysSectionBlocks, tomorrowsSectionBlocks)
	eventBlocks := modal.BuildEventsBlock(todaysSectionBlocks, tomorrowsSectionBlocks)

	message := slack.NewBlockMessage(eventBlocks...)
	message.Text = "bZapp - Today's Standup Summary"
	message.ResponseType = slack.ResponseTypeInChannel

	post, err := Post(url, headers, message)
	if err != nil {
		log.Printf("Error: %s", err)
	} else {
		bytes, _ := ioutil.ReadAll(post.Body)
		log.Printf("Success: %s", string(bytes))
	}

	return events.APIGatewayProxyResponse{
		//Headers:    JsonHeaders(),
		StatusCode: 200,
	}, nil
}

func actionEvent(payload modal.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	log.Printf("action id %s\n", payload.ActionCallback.BlockActions[0].ActionID)
	switch payload.ActionCallback.BlockActions[0].ActionID {
	case modal.EditEventsActionId:
		return pushEditEventModal(payload)
	case modal.RemoveEventActionId:
		return removeEvent(payload)
	}
	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       "Unknown action type",
		StatusCode: 400,
	}, nil
}

func removeEvent(payload modal.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	log.Printf("remove starteddddddddddddddd	sss")
	//actionId := payload.ActionCallback.BlockActions[0].ActionID
	actionValue := payload.ActionCallback.BlockActions[0].Value
	todaysSectionBlocks, tomorrowsSectionBlocks := modal.RemoveEvent(payload.View.Blocks.BlockSet, actionValue)
	index := modal.ExtractInputIndex(payload.View.Blocks.BlockSet)

	modalRequest := modal.NewEditEventsModal(index, todaysSectionBlocks, tomorrowsSectionBlocks)
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
	log.Printf("Json bytes: %v\n", string(jsonBytes))
	//jsonBytes, err := json.Marshal(update)
	//indent, _ := json.MarshalIndent(update, "", "\t")
	//log.Printf("body sent to slack: %v", string(indent))

	return events.APIGatewayProxyResponse{
		Headers: JsonHeaders(),
		//Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func pushModalWithAddedEvent(payload modal.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	action := payload.View.State.Values[modal.AddEventDayInputBlock][modal.AddEventDayActionId]
	marshal, _ := json.Marshal(action)
	fmt.Printf("Add Event button pressed by user %s with value %v\n", payload.User.Name, string(marshal))

	index := modal.ExtractInputIndex(payload.View.Blocks.BlockSet)
	todaysEvents, tomorrowsEvents := modal.ExtractEvents2(payload.View.Blocks.BlockSet)

	newEvent := modal.BuildNewEventSectionBlock2(index, payload.View.State.Values)
	switch newEvent.Day {
	case modal.TodayOptionValue:
		todaysEvents = model.AddEventInOrder(newEvent, todaysEvents)
	case modal.TomorrowOptionValue:
		tomorrowsEvents = model.AddEventInOrder(newEvent, tomorrowsEvents)
	}

	//eventDay, newEvent := modal.BuildNewEventSectionBlock(index, payload.View.State.Values)


	//todaysSectionBlocks, tomorrowsSectionBlocks := modal.AddNewEventToDay(payload.View.Blocks.BlockSet, eventDay, newEvent)
	todaysSectionBlocks, tomorrowsSectionBlocks := modal.ConvertToEventsWithRemoveButton2(todaysEvents, tomorrowsEvents)
	todaysSectionBlocks, tomorrowsSectionBlocks = modal.ReplaceEmptyEventsWithNoEventsYet(todaysSectionBlocks, tomorrowsSectionBlocks)
	fmt.Printf("Addedsssss New gotasdf: %v, got1: %v\n", len(todaysSectionBlocks), len(tomorrowsSectionBlocks))


	modalRequest := modal.NewEditEventsModal(index + 1, todaysSectionBlocks, tomorrowsSectionBlocks)
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

func pushEditEventModal(payload modal.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Message button pressed by user %s with value %v\n", payload.User.Name, payload)

	todaysEvents, tomorrowsEvents := modal.ExtractEvents2(payload.View.Blocks.BlockSet)
	todaysSectionBlocks, tomorrowsSectionEvents := modal.ConvertToEventsWithRemoveButton2(todaysEvents, tomorrowsEvents)
	todaysSectionBlocks, tomorrowsSectionEvents = modal.ReplaceEmptyEventsWithNoEventsYet(todaysSectionBlocks, tomorrowsSectionEvents)
	index :=modal.ExtractInputIndex(payload.View.Blocks.BlockSet)

	modalRequest := modal.NewEditEventsModal(index +1, todaysSectionBlocks, tomorrowsSectionEvents)
	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	//viewResponse, err := api.UpdateView(modalRequest, payload.View.ExternalID, payload.Hash, payload.View.ID)
	viewResponse, err := api.PushView(payload.TriggerID, modalRequest)
	if err != nil {
		log.Printf("Err opening modalss: %v\n", err)
	} else {
		log.Printf("Success open modal: %v\n", viewResponse)
		indent, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success open modal2: %v", string(indent))
	}
	update := slack.NewUpdateViewSubmissionResponse(&modalRequest)
	jsonBytes, err := json.Marshal(update)
	log.Printf("Json bytes: %v\n", format.PrettyJsonNoError(string(jsonBytes)))

	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		StatusCode: 200,
	}, nil
}

func viewClosed(payload modal.InteractionPayload) (events.APIGatewayProxyResponse, error){
	if payload.View.Title.Text == "bZapp - Edit Events" && !payload.IsCleared  {
		return returnToSummaryModal(payload)
	}
	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       fmt.Sprintf("Unimplemented Event Type: %v", payload.Type),
		StatusCode: 400,
	}, nil
}


func returnToSummaryModal(payload modal.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	todaysEvents, tomorrowsEvents := modal.ExtractEvents2(payload.View.Blocks.BlockSet)
	todaysSectionBlocks, tomorrowsSectionBlocks := modal.ConvertToEventsWithoutRemoveButton2(todaysEvents, tomorrowsEvents)
	todaysSectionBlocks, tomorrowsSectionBlocks = modal.ReplaceEmptyEventsWithNoEventsYet(todaysSectionBlocks, tomorrowsSectionBlocks)

	modalRequest := modal.NewSummaryModal(todaysSectionBlocks, tomorrowsSectionBlocks)
	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.UpdateView(modalRequest, payload.View.ExternalID, payload.Hash, payload.View.RootViewID)
	if err != nil {
		log.Printf("Err opening modal: %v\n", err)
	} else {
		indent, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success return to summary modal: %v", string(indent))
	}
	update := slack.NewUpdateViewSubmissionResponse(&modalRequest)
	jsonBytes, err := json.Marshal(update)
	log.Printf("json return to summary modal: %v\n", format.PrettyJsonNoError(string(jsonBytes)))

	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		StatusCode: 200,
	}, nil
}