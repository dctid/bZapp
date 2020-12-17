package bZapp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/model"
	"github.com/dctid/bZapp/view"
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
	var payload slack.InteractionCallback
	err = json.Unmarshal([]byte(m["payload"][0]), &payload)
	if err != nil {
		fmt.Printf("Could not parse action response JSON: %v\n", err)
		return events.APIGatewayProxyResponse{Headers: JsonHeaders(),
			Body:       "Bad payload",
			StatusCode: 400,
		}, err
	}
	var metadata *model.Metadata
	err = json.Unmarshal([]byte(payload.View.PrivateMetadata), &metadata)
	if err != nil {
		log.Printf("Couldn't parse metadata %s", err)
	} else {
		log.Printf("Metadata: %v", metadata)
	}

	currentModel, err := GetModelFromDb(ctx, metadata.ChannelId)

	switch payload.Type {
	case slack.InteractionTypeViewSubmission:
		return viewSubmission(ctx, &payload, currentModel, metadata)
	case slack.InteractionTypeBlockActions:
		return actionEvent(ctx, &payload, currentModel, metadata)
	case slack.InteractionTypeViewClosed:
		return viewClosed(&payload, currentModel, metadata)
	}

	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       fmt.Sprintf("Unimplemented Event Type: %v", payload.Type),
		StatusCode: 400,
	}, nil

}

func viewSubmission(ctx context.Context, payload *slack.InteractionCallback, currentModel *model.Model, currentMetadata *model.Metadata) (events.APIGatewayProxyResponse, error) {
	if payload.View.Title.Text == view.EditGoalsTitle {
		return pushModalWithAddedGoal(ctx, payload, currentModel, currentMetadata)
	} else if payload.View.Title.Text == view.EditEventsTitle {
		return pushModalWithAddedEvent(ctx, payload, currentModel, currentMetadata)
	} else {
		return publishbZapp(currentModel, currentMetadata)
	}
}

func publishbZapp(currentModel *model.Model, metadata *model.Metadata) (events.APIGatewayProxyResponse, error) {

	message := view.DailySummaryMessage(currentModel)

	post, err := Post(
		metadata.ResponseUrl,
		http.Header{
			"Content-type":  []string{"application/json"},
			"Authorization": []string{fmt.Sprintf("Bearer %s", os.Getenv("SLACK_TOKEN"))},
		},
		message,
	)

	if err != nil {
		log.Printf("Error: %s", err)
	} else {
		bytes, _ := ioutil.ReadAll(post.Body)
		bodyString := string(bytes)
		log.Printf("Success: %s", string(bytes))
		if bodyString != "ok" {
			log.Printf("Error!!!!: %s", bodyString)
			return handlePostMessageError(string(bytes))
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func handlePostMessageError(errorMessage string) (events.APIGatewayProxyResponse, error) {

	errorMsg := fmt.Sprintf("Unknown error: %s", errorMessage)
	if errorMessage == "channel_not_found" || errorMessage == "not_in_channel" {
		errorMsg = "It looks like bZapp is not in your private channel :Shrug:. A simple @bzapp mention is you need to do!"
	} else if errorMessage == "expired_url" {
		errorMsg = "It looks like bZapp's comand expired before you could submit :timer_clock:. \nPlease close this modal, then reopen it with `/bzapp` and click submit again."
	}
	modalUpdatedWithNewEvent := view.NewErrorModal(errorMsg)

	jsonBytes, err := json.Marshal(slack.NewUpdateViewSubmissionResponse(modalUpdatedWithNewEvent))
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers:    JsonHeaders(),
			Body:       "Error processing request",
			StatusCode: 500,
		}, err
	}
	log.Printf("body sent to slack: %v", string(jsonBytes))

	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil

}

func pushModalWithAddedEvent(ctx context.Context, payload *slack.InteractionCallback, currentModel *model.Model, currentMetadata *model.Metadata) (events.APIGatewayProxyResponse, error) {

	newEvent := view.BuildNewEvent(currentModel.Index, payload.View.State.Values)
	currentModel.Events = currentModel.Events.AddEvent(newEvent)
	currentModel.Index++

	_ = SaveModel(ctx, currentMetadata.ChannelId, currentModel)

	modalRequest := view.NewEditEventsModal(currentModel, currentMetadata)
	modalUpdatedWithNewEvent := slack.NewUpdateViewSubmissionResponse(modalRequest)
	jsonBytes, err := json.Marshal(modalUpdatedWithNewEvent)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers:    JsonHeaders(),
			Body:       "Error processing request",
			StatusCode: 500,
		}, err
	}
	log.Printf("body sent to slack: %v", string(jsonBytes))

	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func actionEvent(ctx context.Context, payload *slack.InteractionCallback, currentModel *model.Model, currentMetadata *model.Metadata) (events.APIGatewayProxyResponse, error) {
	log.Printf("action id %s\n", payload.ActionCallback.BlockActions[0].ActionID)
	switch payload.ActionCallback.BlockActions[0].ActionID {
	case view.EditEventsActionId:
		return pushEditEventModal(ctx, payload, currentModel, currentMetadata)
	case view.EditGoalsActionId:
		return pushEditGoalsModal(ctx, payload, currentModel, currentMetadata)
	case view.RemoveEventActionId:
		return removeEvent(ctx, payload, currentModel, currentMetadata)
	case view.RemoveGoalActionId:
		return removeGoal(ctx, payload, currentModel, currentMetadata)
	}
	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       "Unknown action type",
		StatusCode: 400,
	}, nil
}

func pushEditEventModal(ctx context.Context, payload *slack.InteractionCallback, currentModel *model.Model, currentMetadata *model.Metadata) (events.APIGatewayProxyResponse, error) {

	currentModel.Index++
	_ = SaveModel(ctx, currentMetadata.ChannelId, currentModel)
	modalRequest := view.NewEditEventsModal(currentModel, currentMetadata)

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.PushView(payload.TriggerID, *modalRequest)
	if err != nil {
		log.Printf("Err opening modalss: %v\n", err)
	} else {
		responseFromSlack, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success pushing edit modal: %v", string(responseFromSlack))
	}
	update := slack.NewUpdateViewSubmissionResponse(modalRequest)
	jsonBytes, err := json.Marshal(update)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers:    JsonHeaders(),
			Body:       "Error processing request",
			StatusCode: 500,
		}, err
	}
	log.Printf("Json bytes: %v\n", format.PrettyJsonNoError(string(jsonBytes)))

	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		StatusCode: 200,
	}, nil
}

func pushEditGoalsModal(ctx context.Context, payload *slack.InteractionCallback, currentModel *model.Model, currentMetadata *model.Metadata) (events.APIGatewayProxyResponse, error) {

	currentModel.Index++
	_ = SaveModel(ctx, currentMetadata.ChannelId, currentModel)
	modalRequest := view.NewEditGoalsModal(currentModel, currentMetadata)

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.PushView(payload.TriggerID, *modalRequest)
	if err != nil {
		log.Printf("Err opening modalss: %v\n", err)
	} else {
		responseFromSlack, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success pushing edit modal: %v", string(responseFromSlack))
	}
	update := slack.NewUpdateViewSubmissionResponse(modalRequest)
	jsonBytes, err := json.Marshal(update)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers:    JsonHeaders(),
			Body:       "Error processing request",
			StatusCode: 500,
		}, err
	}
	log.Printf("Json bytes: %v\n", format.PrettyJsonNoError(string(jsonBytes)))

	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		StatusCode: 200,
	}, nil
}

func pushModalWithAddedGoal(ctx context.Context, payload *slack.InteractionCallback, currentModel *model.Model, currentMetadata *model.Metadata) (events.APIGatewayProxyResponse, error) {

	category, goal := view.BuildNewGoalSectionBlock(currentModel.Index, payload.View.State.Values)
	currentModel.Goals = currentModel.Goals.AddGoal(category, goal)
	currentModel.Index++
	_ = SaveModel(ctx, currentMetadata.ChannelId, currentModel)

	modalRequest := view.NewEditGoalsModal(currentModel, currentMetadata)
	modalUpdatedWithNewEvent := slack.NewUpdateViewSubmissionResponse(modalRequest)

	jsonBytes, err := json.Marshal(modalUpdatedWithNewEvent)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Headers:    JsonHeaders(),
			Body:       "Error processing request",
			StatusCode: 500,
		}, err
	}
	log.Printf("body sent to slack: %v", string(jsonBytes))

	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func removeEvent(ctx context.Context, payload *slack.InteractionCallback, currentModel *model.Model, currentMetadata *model.Metadata) (events.APIGatewayProxyResponse, error) {

	currentModel.Events = currentModel.Events.RemoveEvent(payload.ActionCallback.BlockActions[0].BlockID)
	err := SaveModel(ctx, currentMetadata.ChannelId, currentModel)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}
	modalRequest := view.NewEditEventsModal(currentModel, currentMetadata)
	requestAsJson, _ := json.MarshalIndent(modalRequest, "", "\t")
	log.Printf("Body sent to slack after removing event: %v", string(requestAsJson))

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.UpdateView(*modalRequest, payload.View.ExternalID, payload.Hash, payload.View.ID)
	if err != nil {
		log.Printf("Err removing event from modal: %v\n", err)
	} else {
		responseFromSlack, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success event from modal: %v", string(responseFromSlack))
	}
	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		StatusCode: 200,
	}, nil
}

func removeGoal(ctx context.Context, payload *slack.InteractionCallback, currentModel *model.Model, currentMetadata *model.Metadata) (events.APIGatewayProxyResponse, error) {

	currentModel.Goals = currentModel.Goals.RemoveGoal(payload.ActionCallback.BlockActions[0].BlockID)
	_ = SaveModel(ctx, currentMetadata.ChannelId, currentModel)
	modalRequest := view.NewEditGoalsModal(currentModel, currentMetadata)
	requestAsJson, _ := json.MarshalIndent(modalRequest, "", "\t")
	log.Printf("Body sent to slack after removing goal: %v", string(requestAsJson))

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.UpdateView(*modalRequest, payload.View.ExternalID, payload.Hash, payload.View.ID)
	if err != nil {
		log.Printf("Err removing event from modal: %v\n", err)
	} else {
		responseFromSlack, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success event from modal: %v", string(responseFromSlack))
	}
	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		StatusCode: 200,
	}, nil
}

func viewClosed(payload *slack.InteractionCallback, currentModel *model.Model, currentMetadata *model.Metadata) (events.APIGatewayProxyResponse, error) {
	if (payload.View.Title.Text == view.EditEventsTitle || payload.View.Title.Text == view.EditGoalsTitle) && !payload.IsCleared {
		return returnToSummaryModal(payload, currentModel, currentMetadata)
	}
	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       fmt.Sprintf("Unimplemented Event Type: %v", payload.Type),
		StatusCode: 400,
	}, nil
}

func returnToSummaryModal(payload *slack.InteractionCallback, currentModel *model.Model, currentMetadata *model.Metadata) (events.APIGatewayProxyResponse, error) {
	modalRequest := view.NewSummaryModal(currentModel, currentMetadata)

	update := slack.NewUpdateViewSubmissionResponse(&modalRequest)
	jsonBytes, err := json.Marshal(update)
	log.Printf("json return to summary view: %v\n", format.PrettyJsonNoError(string(jsonBytes)))

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.UpdateView(modalRequest, payload.View.ExternalID, payload.Hash, payload.View.RootViewID)

	if err != nil {
		log.Printf("Err opening view: %v\n", err)
	} else {
		slackResponse, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success return to summary view: %v", string(slackResponse))
	}
	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		StatusCode: 200,
	}, nil
}
