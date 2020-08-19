package bZapp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/view"
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
	var payload view.InteractionPayload
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


func viewSubmission(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	if len(payload.View.State.Values) == 1 && payload.View.State.Values["convo_input_id"]["conversation_select_action_id"].Type == "conversations_select" {
		return publishbZapp(payload)
	} else {
		return pushModalWithAddedEvent(payload)
	}
}

func publishbZapp(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {

	url := payload.ResponseUrls[0].ResponseUrl
	log.Printf("Response Urls: %s", url)
	headers := http.Header{"Content-type": []string{"application/json"}}

	todaysEvents, tomorrowsEvents := view.ExtractEvents(payload.View.Blocks.BlockSet)
	todaysSectionBlocks, tomorrowsSectionBlocks := view.ConvertToEventsWithoutRemoveButton(todaysEvents, tomorrowsEvents)
	eventBlocks := view.BuildEventsBlock(todaysSectionBlocks, tomorrowsSectionBlocks)

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

func actionEvent(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	log.Printf("action id %s\n", payload.ActionCallback.BlockActions[0].ActionID)
	switch payload.ActionCallback.BlockActions[0].ActionID {
	case view.EditEventsActionId:
		return pushEditEventModal(payload)
	case view.RemoveEventActionId:
		return removeEvent(payload)
	}
	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       "Unknown action type",
		StatusCode: 400,
	}, nil
}

func removeEvent(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	log.Printf("remove starteddddddddddddddd	sss")
	blockIdToDelete := payload.ActionCallback.BlockActions[0].BlockID

	todaysEvents, tomorrowsEvents := view.ExtractEvents(payload.View.Blocks.BlockSet)

	todaysEvents = model.RemoveEvent(blockIdToDelete, todaysEvents)
	tomorrowsEvents = model.RemoveEvent(blockIdToDelete, tomorrowsEvents)
	todaysSectionBlocks, tomorrowsSectionBlocks := view.ConvertToEventsWithRemoveButton(todaysEvents, tomorrowsEvents)

	index := view.ExtractInputIndex(payload.View.Blocks.BlockSet)

	modalRequest := view.NewEditEventsModal(index, todaysSectionBlocks, tomorrowsSectionBlocks)

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.UpdateView(modalRequest, payload.View.ExternalID, payload.Hash, payload.View.ID)
	if err != nil {
		log.Printf("Err opening view: %v\n", err)
	} else {
		log.Printf("Success open view: %v\n", viewResponse)
		indent, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success open modal2: %v", string(indent))
	}
	update := slack.NewUpdateViewSubmissionResponse(&modalRequest)
	jsonBytes, err := json.Marshal(update)
	log.Printf("Json bytes: %v\n", string(jsonBytes))

	return events.APIGatewayProxyResponse{
		Headers: JsonHeaders(),
		StatusCode: 200,
	}, nil
}

func pushModalWithAddedEvent(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	action := payload.View.State.Values[view.AddEventDayInputBlock][view.AddEventDayActionId]
	marshal, _ := json.Marshal(action)
	fmt.Printf("aAdd Event button pressed by user %s with value %v\n", payload.User.Name, string(marshal))

	index := view.ExtractInputIndex(payload.View.Blocks.BlockSet)
	todaysEvents, tomorrowsEvents := view.ExtractEvents(payload.View.Blocks.BlockSet)

	newEvent := view.BuildNewEventSectionBlock(index, payload.View.State.Values)
	switch newEvent.Day {
	case view.TodayOptionValue:
		todaysEvents = model.AddEventInOrder(newEvent, todaysEvents)
	case view.TomorrowOptionValue:
		tomorrowsEvents = model.AddEventInOrder(newEvent, tomorrowsEvents)
	}
	todaysSectionBlocks, tomorrowsSectionBlocks := view.ConvertToEventsWithRemoveButton(todaysEvents, tomorrowsEvents)

	modalRequest := view.NewEditEventsModal(index + 1, todaysSectionBlocks, tomorrowsSectionBlocks)
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

func pushEditEventModal(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Message button pressed by user %s with value %v\n", payload.User.Name, payload)

	todaysEvents, tomorrowsEvents := view.ExtractEvents(payload.View.Blocks.BlockSet)
	todaysSectionBlocks, tomorrowsSectionEvents := view.ConvertToEventsWithRemoveButton(todaysEvents, tomorrowsEvents)
	index := view.ExtractInputIndex(payload.View.Blocks.BlockSet)

	modalRequest := view.NewEditEventsModal(index +1, todaysSectionBlocks, tomorrowsSectionEvents)
	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.PushView(payload.TriggerID, modalRequest)
	if err != nil {
		log.Printf("Err opening modalss: %v\n", err)
	} else {
		log.Printf("Success open view: %v\n", viewResponse)
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

func viewClosed(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error){
	if payload.View.Title.Text == "bZapp - Edit Events" && !payload.IsCleared  {
		return returnToSummaryModal(payload)
	}
	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       fmt.Sprintf("Unimplemented Event Type: %v", payload.Type),
		StatusCode: 400,
	}, nil
}


func returnToSummaryModal(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	todaysEvents, tomorrowsEvents := view.ExtractEvents(payload.View.Blocks.BlockSet)
	todaysSectionBlocks, tomorrowsSectionBlocks := view.ConvertToEventsWithoutRemoveButton(todaysEvents, tomorrowsEvents)

	modalRequest := view.NewSummaryModal(todaysSectionBlocks, tomorrowsSectionBlocks)
	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.UpdateView(modalRequest, payload.View.ExternalID, payload.Hash, payload.View.RootViewID)
	if err != nil {
		log.Printf("Err opening view: %v\n", err)
	} else {
		indent, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success return to summary view: %v", string(indent))
	}
	update := slack.NewUpdateViewSubmissionResponse(&modalRequest)
	jsonBytes, err := json.Marshal(update)
	log.Printf("json return to summary view: %v\n", format.PrettyJsonNoError(string(jsonBytes)))

	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		StatusCode: 200,
	}, nil
}