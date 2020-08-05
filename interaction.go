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
	//	"accept": "application/json",
	//	"Authorization": "Bearer [add token]",
	//}
	//var block = slack.NewTextBlockObject("plain_text", "HIII", false, false)

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
	}
	fmt.Printf("Message button pressed by user %s with value %v\n", payload.User.Name, payload)
	modalRequest := NewModal(TodaySection(), NoEventYetSection)
	modalRequest.PrivateMetadata = "test metadata"

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true))
	interaction, err := api.UpdateView(modalRequest,  payload.View.ExternalID, payload.Hash ,payload.View.ID)
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
		Headers:headers,
		//Body: string(jsonBytes),
		StatusCode: 200,
	}, nil
}
