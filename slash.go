package bZapp

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func Slash(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//var headers = map[string]string{
	//	"Content-Type": "application/json",
	//}
	//var block = slack.NewTextBlockObject("plain_text", "HIII", false, false)

	modalRequest := NewModal(NoEventYetSection, NoEventYetSection)

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

