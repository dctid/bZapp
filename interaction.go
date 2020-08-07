package bZapp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/slack-go/slack"
	"log"
	"net/url"
	"os"
)

func Interaction(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var headers = map[string]string{
		"content-Type": "application/json",
	}

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
		return events.APIGatewayProxyResponse{Headers: headers,
			Body:       "Bad payload",
			StatusCode: 400,
		}, err
	}
	switch payload.Type {
	case slack.InteractionTypeViewSubmission:
		return pushModalWithAddedEvent(payload, err, headers)
	case slack.InteractionTypeBlockActions:
		return pushEditEventModal(payload, err, headers)
	}

	return events.APIGatewayProxyResponse{
		Headers:    headers,
		Body:       fmt.Sprintf("Unimplemented Event Type: %v", payload.Type),
		StatusCode: 400,
	}, nil

}

func pushModalWithAddedEvent(payload slack.InteractionCallback, err error, headers map[string]string) (events.APIGatewayProxyResponse, error) {
	action := payload.View.State.Values[AddEventDayInputBlock][AddEventDayActionId]
	marshal, err := json.Marshal(action)
	fmt.Printf("Add Event button pressed by user %s with value %v\n", payload.User.Name, string(marshal))

	eventTitle := payload.View.State.Values[AddEventTitleInputBlock][AddEventTitleActionId].Value
	eventDay := payload.View.State.Values[AddEventDayInputBlock][AddEventDayActionId].SelectedOption.Text.Text
	eventHours := payload.View.State.Values[AddEventHoursInputBlock][AddEventHoursActionId].SelectedOption.Text.Text
	eventMins := payload.View.State.Values[AddEventMinsInputBlock][AddEventMinsActionId].SelectedOption.Text.Text

	fmt.Printf("Add Event title: %s, day: %s, hour: %s, mins: %s\n", eventTitle, eventDay, eventHours, eventMins)
	modalRequest := NewEditEventsModal([]*slack.SectionBlock{TodaySection()}, []*slack.SectionBlock{NoEventYetSection})
	modalRequest.PrivateMetadata = "test metadata"

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true))
	interaction, err := api.UpdateView(modalRequest, payload.View.ExternalID, payload.Hash, payload.View.ID)
	if err != nil {
		log.Printf("Err opening modal: %v\n", err)
	} else {
		log.Printf("Success open modal: %v\n", interaction)
	}
	update := slack.NewUpdateViewSubmissionResponse(&modalRequest)
	jsonBytes, err := json.Marshal(update)
	log.Printf("Json bytes: %v\n", jsonBytes)

	return events.APIGatewayProxyResponse{
		Headers:    headers,
		StatusCode: 200,
	}, nil
}

func pushEditEventModal(payload slack.InteractionCallback, err error, headers map[string]string) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Message button pressed by user %s with value %v\n", payload.User.Name, payload)
	modalRequest := NewEditEventsModal([]*slack.SectionBlock{TodaySection()}, []*slack.SectionBlock{NoEventYetSection})
	modalRequest.PrivateMetadata = "test metadata"

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true))
	interaction, err := api.UpdateView(modalRequest, payload.View.ExternalID, payload.Hash, payload.View.ID)
	if err != nil {
		log.Printf("Err opening modal: %v\n", err)
	} else {
		log.Printf("Success open modal: %v\n", interaction)
	}
	update := slack.NewUpdateViewSubmissionResponse(&modalRequest)
	jsonBytes, err := json.Marshal(update)
	log.Printf("Json bytes: %v\n", jsonBytes)
	//var bodyMap map[string]interface{}
	//err = json.Unmarshal([]byte(event.Body), &bodyMap)
	//if err != nil {
	//	log.Printf("Err parsing body: %v\n", err)
	//	return events.APIGatewayProxyResponse{
	//		StatusCode: 500,
	//	}, err
	//}

	//triggerId := m["trigger_id"][0]// fmt.Sprintf("%v", bodyMap["trigger_id"])
	//modalRequest := NewModal(NoEventYetSection, NoEventYetSection)

	//jsonBytes, err := json.Marshal(modalRequest)
	//log.Printf("json %s", jsonBytes)

	//postHeaders := http.Header{"Content-Type": {"application/json"},
	//	"accept": {"application/json"},
	//	"Authorization": {"Bearer [add token]"}}
	//
	//_, err = Post("https://slack.com/api/views.open", postHeaders, modalRequest)
	//if err != nil {
	//	return events.APIGatewayProxyResponse{
	//		StatusCode: 500,
	//	}, err
	//}

	//defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//println(string(body))

	return events.APIGatewayProxyResponse{
		Headers: headers,
		//Body: string(jsonBytes),
		StatusCode: 200,
	}, nil
}
