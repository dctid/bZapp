package bZapp

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/slack-go/slack"
	"log"
	"net/url"
	"os"
)

func Slash(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//var headers = map[string]string{
	//	"Content-Type": "application/json",
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
	}

	//var bodyMap map[string]interface{}
	//err = json.Unmarshal([]byte(event.Body), &bodyMap)
	//if err != nil {
	//	log.Printf("Err parsing body: %v\n", err)
	//	return events.APIGatewayProxyResponse{
	//		StatusCode: 500,
	//	}, err
	//}

	triggerId := m["trigger_id"][0]// fmt.Sprintf("%v", bodyMap["trigger_id"])
	modalRequest := NewModal(NoEventYetSection, NoEventYetSection)
	//modalRequest.ExternalID = "adsbadfbadf"

	api := slack.New(os.Getenv("SLACK_TOKEN"), slack.OptionDebug(true), slack.OptionHTTPClient(Client))
	viewResponse, err := api.OpenView(triggerId, modalRequest)
	if err != nil {
		log.Printf("Err opening modal: %v", err)
	} else {
		indent, _ := json.MarshalIndent(viewResponse, "", "\t")
		log.Printf("Success open modal: %v", string(indent))
	}

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

		StatusCode: 200,
	}, nil
}
