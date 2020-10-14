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
	var payload view.InteractionPayload
	err = json.Unmarshal([]byte(m["payload"][0]), &payload)
	if err != nil {
		fmt.Printf("Could not parse action response JSON: %v\n", err)
		return events.APIGatewayProxyResponse{Headers: JsonHeaders(),
			Body:       "Bad payload",
			StatusCode: 400,
		}, err
	}
	var currentModel model.Model
	err = json.Unmarshal([]byte(payload.View.PrivateMetadata), &currentModel)
	if err != nil {
		log.Printf("Couldn't parse metadata %s", err)
	} else {
		log.Printf("Metadata: %v", currentModel)
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
	} else if payload.View.Title.Text == view.EditGoalsTitle {
		return pushModalWithAddedGoal(payload)
	} else {
		return pushModalWithAddedEvent(payload)
	}
}

func publishbZapp(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {

	url, message := view.DailySummaryMessage(payload)

	post, err := Post(url, http.Header{"Content-type": []string{"application/json"}}, message)
	if err != nil {
		log.Printf("Error: %s", err)
	} else {
		bytes, _ := ioutil.ReadAll(post.Body)
		log.Printf("Success: %s", string(bytes))
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func pushModalWithAddedEvent(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	modalUpdatedWithNewEvent := view.AddEventToEditModal(payload)
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

func actionEvent(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	log.Printf("action id %s\n", payload.ActionCallback.BlockActions[0].ActionID)
	switch payload.ActionCallback.BlockActions[0].ActionID {
	case view.EditEventsActionId:
		return pushEditEventModal(payload)
	case view.EditGoalsActionId:
		return pushEditGoalsModal(payload)
	case view.RemoveEventActionId:
		return removeEvent(payload)
	case view.RemoveGoalActionId:
		return removeGoal(payload)
	}
	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       "Unknown action type",
		StatusCode: 400,
	}, nil
}

func pushEditEventModal(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {

	modalRequest := view.OpenEditEventModalFromSummaryModal(payload)

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.PushView(payload.TriggerID, modalRequest)
	if err != nil {
		log.Printf("Err opening modalss: %v\n", err)
	} else {
		responseFromSlack, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success pushing edit modal: %v", string(responseFromSlack))
	}
	update := slack.NewUpdateViewSubmissionResponse(&modalRequest)
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

func pushEditGoalsModal(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {

	modalRequest := view.OpenEditGoalsModalFromSummaryModal(payload)

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.PushView(payload.TriggerID, modalRequest)
	if err != nil {
		log.Printf("Err opening modalss: %v\n", err)
	} else {
		responseFromSlack, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success pushing edit modal: %v", string(responseFromSlack))
	}
	update := slack.NewUpdateViewSubmissionResponse(&modalRequest)
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

func pushModalWithAddedGoal(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	modalUpdatedWithNewEvent := view.AddGoalToEditModal(payload)
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


func removeEvent(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	modalRequest := view.RemoveEventFromEditModal(payload)
	requestAsJson, _ := json.MarshalIndent(modalRequest, "", "\t")
	log.Printf("Body sent to slack after removing event: %v", string(requestAsJson))

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.UpdateView(modalRequest, payload.View.ExternalID, payload.Hash, payload.View.ID)
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

func removeGoal(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	modalRequest := view.RemoveGoalFromEditModal(payload)
	requestAsJson, _ := json.MarshalIndent(modalRequest, "", "\t")
	log.Printf("Body sent to slack after removing goal: %v", string(requestAsJson))

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.UpdateView(modalRequest, payload.View.ExternalID, payload.Hash, payload.View.ID)
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

func viewClosed(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	if payload.View.Title.Text == "bZapp - Edit Events" && !payload.IsCleared {
		return returnToSummaryModal(payload)
	}
	return events.APIGatewayProxyResponse{
		Headers:    JsonHeaders(),
		Body:       fmt.Sprintf("Unimplemented Event Type: %v", payload.Type),
		StatusCode: 400,
	}, nil
}

func returnToSummaryModal(payload view.InteractionPayload) (events.APIGatewayProxyResponse, error) {
	modalRequest := view.SummaryModalWithEventsAddedInEditModal(payload)

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
