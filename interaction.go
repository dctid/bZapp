package bZapp

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/url"
)

func Interaction(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//var headers = map[string]string{
	//	"Content-Type": "application/json",
	//	"accept": "application/json",
	//	"Authorization": "Bearer xoxb-260884079521-1098577169670-mQjaNJ7Sx6OTEycmjCvKJyTr",
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
		log.Printf("Body: %v", m)
	}

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
	//api := slack.New("xoxb-260884079521-1098577169670-mQjaNJ7Sx6OTEycmjCvKJyTr", slack.OptionDebug(true))
	//interaction, err := api.OpenView(triggerId, modalRequest)
	//if err != nil {
	//	log.Printf("Err opening modal: %v", err)
	//} else {
	//	log.Printf("Success open modal: %v", interaction)
	//}

	//jsonBytes, err := json.Marshal(modalRequest)
	//log.Printf("json %s", jsonBytes)

	//postHeaders := http.Header{"Content-Type": {"application/json"},
	//	"accept": {"application/json"},
	//	"Authorization": {"Bearer xoxb-260884079521-1098577169670-mQjaNJ7Sx6OTEycmjCvKJyTr"}}
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
