package bZapp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/mocks"
	"github.com/dctid/bZapp/model"
	"github.com/dctid/bZapp/test"
	"github.com/dctid/bZapp/view"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"testing"
)

var successResponse = &http.Response{
	Body:       ioutil.NopCloser(bytes.NewReader([]byte(response))),
	StatusCode: 200,
}

func getItemOutput(modelToReturn *model.Model) *dynamodb.GetItemOutput {
	modelBytes, _ := json.Marshal(modelToReturn)
	return &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("D7P4LC5G9"),
			},
			"model": {
				S: aws.String(string(modelBytes)),
			},
		},
	}
}

func putItemInput(modelToSave *model.Model) *dynamodb.PutItemInput {
	modelBytes, _ := json.Marshal(modelToSave)
	return &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("D7P4LC5G9"),
			},
			"model": {
				S: aws.String(string(modelBytes)),
			},
		},
		TableName: aws.String("bZappTable"),
	}
}

var getItemInput = &dynamodb.GetItemInput{
	Key: map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String("D7P4LC5G9"),
		},
	},
	TableName: aws.String("bZappTable"),
}

func TestInteraction(t *testing.T) {
	Client = &mocks.MockClient{}

	os.Setenv("SLACK_TOKEN", "token_token")

	type args struct {
		ctx   context.Context
		event events.APIGatewayProxyRequest
	}
	type do struct {
		url     *url.URL
		body    string
		headers http.Header
	}
	var gotDo do

	tests := []struct {
		name            string
		args            args
		response        *http.Response
		want            events.APIGatewayProxyResponse
		wantErr         bool
		wantDo          do
		dynamoResponses *mocks.MockDynamoDB
		wantDynamoCalls *mocks.MockDynamoDbInputs
		date            string
	}{
		{
			name:     "open edit events",
			args:     args{event: events.APIGatewayProxyRequest{Body: test.EditEventsActionButton}},
			response: successResponse,
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:     getUrl("https://slack.com/api/views.push"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-Type": []string{"application/json"}},
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"trigger_id": "1288231154914.260884079521.ba1595ee20fab577e5ac042a518713fd",
								"view": %s
							}`, test.EditEventsModal)),
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: getItemOutput(&model.Model{}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: getItemInput,
				PutItemWithContext: putItemInput(&model.Model{
					Index: 1,
				}),
			},
			date: "2020-12-02 08:48:21",
		},
		{
			name:     "remove event",
			args:     args{event: events.APIGatewayProxyRequest{Body: test.RemoveEventAction}},
			response: successResponse,
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:     getUrl("https://slack.com/api/views.update"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-Type": []string{"application/json"}},
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"view_id": "V01CMKMUWUS",
								"view": %s
							}`, test.RemoveEventsModal)),
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: getItemOutput(&model.Model{
					Index: 4,
					Events: model.Events{
						"2020-11-28": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						}, {
							Id:    "YUBFMb",
							Title: "wer",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   30,
							AmPm:  "AM",
						},
						},
						"2020-12-02": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						}, {
							Id:    "YUBFMb",
							Title: "wer",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   30,
							AmPm:  "AM",
						},
						},
						"2020-12-03": []model.Event{{
							Id:    "PTjSgI",
							Title: "dfs",
							Day:   view.TomorrowOptionValue,
							Hour:  11,
							Min:   15,
							AmPm:  "AM",
						}, {
							Id:    "YUBFMb",
							Title: "wer",
							Day:   view.TomorrowOptionValue,
							Hour:  10,
							Min:   30,
							AmPm:  "AM",
						},
						},
					},
				},
				),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: getItemInput,
				PutItemWithContext: putItemInput(&model.Model{
					Index: 4,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						},
						},
						"2020-12-03": []model.Event{{
							Id:    "PTjSgI",
							Title: "dfs",
							Day:   view.TomorrowOptionValue,
							Hour:  11,
							Min:   15,
							AmPm:  "AM",
						},
						},
					},
				}),
			},
			date: "2020-12-02 08:48:21",
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
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: getItemOutput(&model.Model{
					Index: 4,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						}},
						"2020-12-03": []model.Event{{
							Id:    "PTjSgI",
							Title: "dfs",
							Day:   view.TomorrowOptionValue,
							Hour:  11,
							Min:   15,
							AmPm:  "AM",
						}},
					},
				},
				),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: getItemInput,
				PutItemWithContext: putItemInput(&model.Model{
					Index: 5,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						}, {
							Id:    "Fake hash",
							Title: "sd",
							Day:   view.TodayOptionValue,
							Hour:  1,
							Min:   15,
							AmPm:  "PM",
						},
						},
						"2020-12-03": []model.Event{{
							Id:    "PTjSgI",
							Title: "dfs",
							Day:   view.TomorrowOptionValue,
							Hour:  11,
							Min:   15,
							AmPm:  "AM",
						},
						},
					},
				}),
			},
			date: "2020-12-02 08:48:21",
		},
		{
			name:     "submit and send message to channel",
			args:     args{event: events.APIGatewayProxyRequest{Body: test.SubmitPayload}},
			response: successResponse,
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
			},
			wantErr: false,
			wantDo: do{
				url:     getUrl("https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"),
				body:    format.PrettyJsonNoError(test.SubmissionJson),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-type": []string{"application/json"}},
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: getItemOutput(&model.Model{
					Index: 6,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "coEbHc",
							Title: "asdf",
							Day:   view.TodayOptionValue,
							Hour:  9,
							Min:   0,
							AmPm:  "AM",
						}},
						"2020-12-03": []model.Event{{
							Id:    "DZosTr",
							Title: "qewr",
							Day:   view.TomorrowOptionValue,
							Hour:  1,
							Min:   15,
							AmPm:  "PM",
						}},
					},
					Goals: model.Goals{
						"Questions?": []model.Goal{
							{Id: "RrMdIA", Value: "afsasdf"},
						},
						"Team Needs": []model.Goal{
							{Id: "bSAnHN", Value: "sfd"},
						},
					},
				},
				),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: getItemInput,
			},
			date: "2020-12-02 08:48:21",
		},
		{
			name: "submit and send message to private channel bzapp is not a member",
			args: args{event: events.APIGatewayProxyRequest{Body: test.SubmitPayload}},
			response: &http.Response{
				Body: ioutil.NopCloser(bytes.NewReader([]byte(
					`{
					  "ok": false,
					  "error": "channel_not_found",
					  "warning": "missing_charset",
					  "response_metadata": {
						"warnings": [
						  "missing_charset"
						]
					  }
					}`,
				))),
				StatusCode: 200,
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body: format.PrettyJsonNoError(`{
					"response_action": "update",
					"view": {
						"type": "modal",
						"title": {
							"type": "plain_text",
							"text": "bZapp",
							"emoji": true
						},
						"blocks": [
						{
							"type": "section",
							"text": {
								"type": "plain_text",
								"text": "It looks like bZapp is not in your private channel :Shrug:. A simple @bzapp mention is you need to do!"
							}
						}
					],
						"close": {
						"type": "plain_text",
						"text": "Close",
						"emoji": true
					}
					}
				}`),
			},
			wantErr: false,
			wantDo: do{
				url:     getUrl("https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"),
				body:    format.PrettyJsonNoError(test.SubmissionJson),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-type": []string{"application/json"}},
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: getItemOutput(&model.Model{
					Index: 6,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "coEbHc",
							Title: "asdf",
							Day:   view.TodayOptionValue,
							Hour:  9,
							Min:   0,
							AmPm:  "AM",
						}},
						"2020-12-03": []model.Event{{
							Id:    "DZosTr",
							Title: "qewr",
							Day:   view.TomorrowOptionValue,
							Hour:  1,
							Min:   15,
							AmPm:  "PM",
						}},
					},
					Goals: model.Goals{
						"Questions?": []model.Goal{
							{Id: "RrMdIA", Value: "afsasdf"},
						},
						"Team Needs": []model.Goal{
							{Id: "bSAnHN", Value: "sfd"},
						},
					},
				},
				),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: getItemInput,
			},
			date: "2020-12-02 08:48:21",
		},
		{
			name:     "close edit events",
			args:     args{event: events.APIGatewayProxyRequest{Body: test.CloseEditEvents}},
			response: successResponse,
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:     getUrl("https://slack.com/api/views.update"),
				body:    format.PrettyJsonNoError(test.SummaryModal),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-Type": []string{"application/json"}},
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: getItemOutput(&model.Model{
					Index: 5,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						}},
						"2020-12-03": []model.Event{{
							Id:    "PTjSgI",
							Title: "dfs",
							Day:   view.TomorrowOptionValue,
							Hour:  11,
							Min:   15,
							AmPm:  "AM",
						}},
					},
				},
				),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: getItemInput,
			},
			date: "2020-12-02 08:48:21",
		},
		{
			name:     "open edit goals actions",
			args:     args{event: events.APIGatewayProxyRequest{Body: test.EditGoalsActionButton}},
			response: successResponse,
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:     getUrl("https://slack.com/api/views.push"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-Type": []string{"application/json"}},
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"trigger_id": "1411346195543.260884079521.14fdd4f0ec90fe20a07ea8dc9429d891",
								"view": %s
							}`, test.EditGoalsModal)),
			},

			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: getItemOutput(&model.Model{}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: getItemInput,
				PutItemWithContext: putItemInput(&model.Model{
					Index: 1,
				}),
			},
			date: "2020-12-02 08:48:21",
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
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: getItemOutput(&model.Model{
					Index: 6,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						}},
						"2020-12-03": []model.Event{{
							Id:    "PTjSgI",
							Title: "dfs",
							Day:   view.TomorrowOptionValue,
							Hour:  11,
							Min:   15,
							AmPm:  "AM",
						}},
					},
				}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: getItemInput,
				PutItemWithContext: putItemInput(&model.Model{
					Index: 7,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						},
						},
						"2020-12-03": []model.Event{{
							Id:    "PTjSgI",
							Title: "dfs",
							Day:   view.TomorrowOptionValue,
							Hour:  11,
							Min:   15,
							AmPm:  "AM",
						},
						},
					},
					Goals: model.Goals{
						"Team Needs": []model.Goal{
							{Id: "Fake hash", Value: "lskfd"},
						},
					},
				}),
			},
			date: "2020-12-02 08:48:21",
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
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: getItemOutput(&model.Model{
					Index: 7,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						}},
						"2020-12-03": []model.Event{{
							Id:    "PTjSgI",
							Title: "dfs",
							Day:   view.TomorrowOptionValue,
							Hour:  11,
							Min:   15,
							AmPm:  "AM",
						}},
					},
					Goals: model.Goals{
						"Team Needs": []model.Goal{
							{Id: "YbiWhf", Value: "lskfd"},
						},
					},
				}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: getItemInput,
				PutItemWithContext: putItemInput(&model.Model{
					Index: 8,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						},
						},
						"2020-12-03": []model.Event{{
							Id:    "PTjSgI",
							Title: "dfs",
							Day:   view.TomorrowOptionValue,
							Hour:  11,
							Min:   15,
							AmPm:  "AM",
						},
						},
					},
					Goals: model.Goals{
						"Questions?": []model.Goal{
							{Id: "Fake hash", Value: "adsf"},
						},
						"Team Needs": []model.Goal{
							{Id: "YbiWhf", Value: "lskfd"},
						},
					},
				}),
			},
			date: "2020-12-02 08:48:21",
		},
		{
			name:     "remove goal actions",
			args:     args{event: events.APIGatewayProxyRequest{Body: test.RemoveGoalAction}},
			response: successResponse,
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:     getUrl("https://slack.com/api/views.update"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-Type": []string{"application/json"}},
				body: format.PrettyJsonNoError(fmt.Sprintf(
					`{
								"view_id": "V01DBFTR588",
								"view": %s
							}`, test.RemoveGoalsModal)),
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: getItemOutput(&model.Model{
					Index: 9,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						}},
						"2020-12-03": []model.Event{{
							Id:    "PTjSgI",
							Title: "dfs",
							Day:   view.TomorrowOptionValue,
							Hour:  11,
							Min:   15,
							AmPm:  "AM",
						}},
					},
					Goals: model.Goals{
						"Learnings": []model.Goal{
							{Id: "mopNVQ", Value: "sdfg"},
						},
						"Questions?": []model.Goal{
							{Id: "GqrOzx", Value: "adsf"},
						},
						"Team Needs": []model.Goal{
							{Id: "YbiWhf", Value: "lskfd"},
						},
					},
				}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: getItemInput,
				PutItemWithContext: putItemInput(&model.Model{
					Index: 9,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						},
						},
						"2020-12-03": []model.Event{{
							Id:    "PTjSgI",
							Title: "dfs",
							Day:   view.TomorrowOptionValue,
							Hour:  11,
							Min:   15,
							AmPm:  "AM",
						},
						},
					},
					Goals: model.Goals{
						"Learnings": []model.Goal{
							{Id: "mopNVQ", Value: "sdfg"},
						},
						"Questions?": []model.Goal{
						},
						"Team Needs": []model.Goal{
							{Id: "YbiWhf", Value: "lskfd"},
						},
					},
				}),
			},
			date: "2020-12-02 08:48:21",
		},
		{
			name:     "close edit goals",
			args:     args{event: events.APIGatewayProxyRequest{Body: test.CloseEditGoals}},
			response: successResponse,
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:     getUrl("https://slack.com/api/views.update"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-Type": []string{"application/json"}},
				body:    format.PrettyJsonNoError(test.SummaryModalWithGoals),
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: getItemOutput(&model.Model{
					Index: 9,
					Events: model.Events{
						"2020-12-02": []model.Event{{
							Id:    "WyKVYV",
							Title: "ads",
							Day:   view.TodayOptionValue,
							Hour:  10,
							Min:   0,
							AmPm:  "AM",
						}},
						"2020-12-03": []model.Event{{
							Id:    "PTjSgI",
							Title: "dfs",
							Day:   view.TomorrowOptionValue,
							Hour:  11,
							Min:   15,
							AmPm:  "AM",
						}},
					},
					Goals: model.Goals{
						"Learnings": []model.Goal{
							{Id: "mopNVQ", Value: "sdfg"},
						},
						"Questions?": []model.Goal{
						},
						"Team Needs": []model.Goal{
							{Id: "YbiWhf", Value: "lskfd"},
						},
					},
				}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: getItemInput,
			},
			date: "2020-12-02 08:48:21",
		},
	}
	for _, tt := range tests {
		model.Clock = mocks.NewMockClock(tt.date)
		model.Hash = func() string {
			return "Fake hash"
		}
		gotDo = do{}
		DynamoDB = tt.dynamoResponses
		mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
			log.Printf("url %s ", req.URL)
			body, _ := ioutil.ReadAll(req.Body)
			gotDo = do{
				url:     req.URL,
				body:    format.PrettyJsonNoError(string(body)),
				headers: req.Header,
			}

			return tt.response, nil
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := Interaction(context.Background(), tt.args.event)
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
			if tt.wantDynamoCalls != nil && !reflect.DeepEqual(*tt.wantDynamoCalls, mocks.MockDynamoDbCalls) {
				var wantModel model.Model
				var calledModel model.Model
				if tt.wantDynamoCalls.PutItemWithContext != nil {
					_ = json.Unmarshal(tt.wantDynamoCalls.PutItemWithContext.Item["model"].B, &wantModel)
				}
				if mocks.MockDynamoDbCalls.PutItemWithContext != nil {
					_ = json.Unmarshal(mocks.MockDynamoDbCalls.PutItemWithContext.Item["model"].B, &calledModel)
				}
				assert.EqualValues(t, wantModel, calledModel)

				assert.EqualValues(t, tt.wantDynamoCalls.GetItemWithContext, mocks.MockDynamoDbCalls.GetItemWithContext)
				assert.EqualValues(t, tt.wantDynamoCalls.PutItemWithContext, mocks.MockDynamoDbCalls.PutItemWithContext)
				assert.EqualValues(t, tt.wantDynamoCalls.DeleteItemWithContext, mocks.MockDynamoDbCalls.DeleteItemWithContext)
			}
		})
		mocks.ResetMockDynamoDbCalls()
	}
}

func getUrl(urlString string) *url.URL {
	result, _ := url.Parse(urlString)
	return result
}
