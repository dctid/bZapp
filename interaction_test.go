package bZapp

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/mocks"
	"github.com/dctid/bZapp/model"
	"github.com/dctid/bZapp/test"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"
)

func TestInteraction(t *testing.T) {
	Client = &mocks.MockClient{}

	os.Setenv("SLACK_TOKEN", "token_token")

	type args struct {
		ctx   context.Context
		event events.APIGatewayProxyRequest
	}
	type do struct {
		url  *url.URL
		body string
		headers http.Header
	}
	var gotDo do

	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
		wantDo  do
	}{
		{
			name: "open edit events",
			args: args{event: events.APIGatewayProxyRequest{Body: test.EditEventsActionButton}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url: getUrl("https://slack.com/api/views.push"),
				headers: http.Header{"Authorization":[]string{"Bearer token_token"}, "Content-Type":[]string{"application/json"}},
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"trigger_id": "1288231154914.260884079521.ba1595ee20fab577e5ac042a518713fd",
    							"view": %s
							}`, test.EditEventsModal)),
			},
		},
		{
			name: "remove event",
			args: args{event: events.APIGatewayProxyRequest{Body: test.RemoveEventAction}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url: getUrl("https://slack.com/api/views.update"),
				headers: http.Header{"Authorization":[]string{"Bearer token_token"}, "Content-Type":[]string{"application/json"}},
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"view_id": "V01CMKMUWUS",
    							"view": %s
							}`, test.RemoveEventsModal)),
			},
		},
		{
			name: "add event submission",
			args: args{event: events.APIGatewayProxyRequest{Body: test.AddEventSubmission}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
				Body:       format.PrettyJsonNoError(test.AddEventSubmissionResponse),
			},
			wantErr: false,
			wantDo:  do{},
		},
		{
			name: "submit and send message to channel",
			args: args{event: events.APIGatewayProxyRequest{Body: test.SubmitPayload}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
			},
			wantErr: false,
			wantDo: do{
				url:  getUrl("https://slack.com/api/chat.postMessage"),
				body: format.PrettyJsonNoError(test.SubmissionJson),
				headers: http.Header{"Authorization":[]string{"Bearer token_token"}, "Content-type":[]string{"application/json"}},
			},
		},
		{
			name: "close edit events",
			args: args{event: events.APIGatewayProxyRequest{Body: test.CloseEditEvents}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:  getUrl("https://slack.com/api/views.update"),
				body: format.PrettyJsonNoError(test.SummaryModal),
				headers: http.Header{"Authorization":[]string{"Bearer token_token"}, "Content-Type":[]string{"application/json"}},

			},
		},
		{
			name: "open edit goals actions",
			args: args{event: events.APIGatewayProxyRequest{Body: test.EditGoalsActionButton}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url: getUrl("https://slack.com/api/views.push"),
				headers: http.Header{"Authorization":[]string{"Bearer token_token"}, "Content-Type":[]string{"application/json"}},
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"trigger_id": "1411346195543.260884079521.14fdd4f0ec90fe20a07ea8dc9429d891",
    							"view": %s
							}`, test.EditGoalsModal)),
			},
		},
		{
			name: "add goal submission",
			args: args{event: events.APIGatewayProxyRequest{Body: test.AddGoalSubmission}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
				Body:       format.PrettyJsonNoError(test.AddGoalSubmissionResponse),
			},
			wantErr: false,
			wantDo:  do{},
		},
		{
			name: "add goal 2nd submission",
			args: args{event: events.APIGatewayProxyRequest{Body: test.Add2ndGoalSubmission}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
				Body:       format.PrettyJsonNoError(test.Add2ndGoalSubmissionResponse),
			},
			wantErr: false,
			wantDo:  do{},
		},
		{
			name: "remove goal actions",
			args: args{event: events.APIGatewayProxyRequest{Body: test.RemoveGoalAction}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url: getUrl("https://slack.com/api/views.update"),
				headers: http.Header{"Authorization":[]string{"Bearer token_token"}, "Content-Type":[]string{"application/json"}},
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"view_id": "V01DBFTR588",
    							"view": %s
							}`, test.RemoveGoalsModal)),
			},
		},
		{
			name: "close edit goals",
			args: args{event: events.APIGatewayProxyRequest{Body: test.CloseEditGoals}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:  getUrl("https://slack.com/api/views.update"),
				headers: http.Header{"Authorization":[]string{"Bearer token_token"}, "Content-Type":[]string{"application/json"}},
				body: format.PrettyJsonNoError(test.SummaryModalWithGoals),
			},
		},
	}
	for _, tt := range tests {
		model.Hash = func() string {
			return "Fake hash"
		}
		gotDo = do{}
		mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
			log.Printf("url %s ", req.URL)
			body, _ := ioutil.ReadAll(req.Body)
			gotDo = do{
				url:  req.URL,
				body: format.PrettyJsonNoError(string(body)),
				headers: req.Header,
			}

			return &http.Response{
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(response))),
				StatusCode: 200,
			}, nil
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := Interaction(tt.args.ctx, tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("Interaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				if !assert.EqualValues(t, tt.want.Body, format.PrettyJsonNoError(got.Body)) {
					t.Errorf("Interaction() got = %v, want %v", got, tt.want)
				}
			}
			if !reflect.DeepEqual(tt.wantDo, do{}) {
				assert.EqualValues(t, tt.wantDo.url, gotDo.url)
				assert.EqualValues(t, tt.wantDo.body, gotDo.body)
				assert.EqualValues(t, tt.wantDo.headers, gotDo.headers)
			} else {
				assert.EqualValues(t, do{}, gotDo)
			}
		})
	}
}

func getUrl(urlString string) *url.URL {
	result, _ := url.Parse(urlString)
	return result
}

