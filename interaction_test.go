package bZapp

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
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
			args:     args{event: events.APIGatewayProxyRequest{Body: test.MakePayload(test.EditEventsActionButtonPayload)}},
			response: test.SuccessResponse(test.ReadFile(t, "slash/slash_response.json")),
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:     test.ParseUrl("https://slack.com/api/views.push"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-Type": []string{"application/json"}},
				body:    test.ReadFile(t, "interaction/edit_events_modal_response.json"),
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: test.GetItemOutput("open_edit_events_channel_id", &model.Model{}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: test.GetItemInput("open_edit_events_channel_id"),
				PutItemWithContext: test.PutItemInput("open_edit_events_channel_id", &model.Model{
					Index: 1,
				}),
			},
			date: "2020-12-02 08:48:21",
		},
		{
			name:     "remove event",
			args:     args{event: events.APIGatewayProxyRequest{Body: test.MakePayload(test.RemoveEventActionPayload)}},
			response: test.SuccessResponse(test.ReadFile(t, "slash/slash_response.json")),
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:     test.ParseUrl("https://slack.com/api/views.update"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-Type": []string{"application/json"}},
				body:    test.ReadFile(t, "interaction/add_event_modal_with_event_removed.json"),
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: test.GetItemOutput("remove_event_channel_id", &model.Model{
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
				}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: test.GetItemInput("remove_event_channel_id"),
				PutItemWithContext: test.PutItemInput("remove_event_channel_id", &model.Model{
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
			args: args{event: events.APIGatewayProxyRequest{Body: test.MakePayload(test.AddEventSubmissionPayload)}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
				Body:       test.ReadFile(t, "interaction/add_event_response.json"),
			},
			wantErr: false,
			wantDo:  do{},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: test.GetItemOutput("add_event_channel_id", &model.Model{
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
				}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: test.GetItemInput("add_event_channel_id"),
				PutItemWithContext: test.PutItemInput("add_event_channel_id", &model.Model{
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
			args:     args{event: events.APIGatewayProxyRequest{Body: test.MakePayload(test.SubmitPayload)}},
			response: test.SuccessResponse("ok"),
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
			},
			wantErr: false,
			wantDo: do{
				url:     test.ParseUrl("https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"),
				body:    test.ReadFile(t, "interaction/post_message_to_channel.json"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-type": []string{"application/json"}},
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: test.GetItemOutput("send_message_channel_id", &model.Model{
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
				}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: test.GetItemInput("send_message_channel_id"),
			},
			date: "2020-12-02 08:48:21",
		},
		{
			name:     "submit and send message when response url is expired",
			args:     args{event: events.APIGatewayProxyRequest{Body: test.MakePayload(test.SubmitPayload)}},
			response: test.SuccessResponse("expired_url"),
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body: test.ReadFile(t, "interaction/post_message_url_expired.json"),
			},
			wantErr: false,
			wantDo: do{
				url:     test.ParseUrl("https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"),
				body:    test.ReadFile(t, "interaction/post_message_to_channel.json"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-type": []string{"application/json"}},
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: test.GetItemOutput("send_message_channel_id", &model.Model{
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
				}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: test.GetItemInput("send_message_channel_id"),
			},
			date: "2020-12-02 08:48:21",
		},
		{
			name:     "close edit events",
			args:     args{event: events.APIGatewayProxyRequest{Body: test.MakePayload(test.CloseEditEventsPayload)}},
			response: test.SuccessResponse(test.ReadFile(t, "slash/slash_response.json")),
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:     test.ParseUrl("https://slack.com/api/views.update"),
				body:    test.ReadFile(t, "interaction/summary_modal_with_added_events_body.json"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-Type": []string{"application/json"}},
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: test.GetItemOutput("close_edit_events_channel_id", &model.Model{
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
				}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: test.GetItemInput("close_edit_events_channel_id"),
			},
			date: "2020-12-02 08:48:21",
		},
		{
			name:     "open edit goals actions",
			args:     args{event: events.APIGatewayProxyRequest{Body: test.MakePayload(test.EditGoalsActionButtonPayload)}},
			response: test.SuccessResponse(test.ReadFile(t, "slash/slash_response.json")),
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:     test.ParseUrl("https://slack.com/api/views.push"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-Type": []string{"application/json"}},
				body:    test.ReadFile(t, "interaction/edit_goals_modal_response.json"),
			},

			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: test.GetItemOutput("open_edit_events_channel_id", &model.Model{}),
			},
			wantDynamoCalls: &mocks.MockDynamoDbInputs{
				GetItemWithContext: test.GetItemInput("open_edit_events_channel_id"),
				PutItemWithContext: test.PutItemInput("open_edit_events_channel_id", &model.Model{
					Index: 1,
				}),
			},
			date: "2020-12-02 08:48:21",
		},
		{
			name: "add goal submission",
			args: args{event: events.APIGatewayProxyRequest{Body: test.MakePayload(test.AddGoalSubmissionPayload)}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
				Body:        test.ReadFile(t, "interaction/add_goal_response.json"),
			},
			wantErr: false,
			wantDo:  do{},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: test.GetItemOutput("add_goal_channel_id", &model.Model{
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
				GetItemWithContext: test.GetItemInput("add_goal_channel_id"),
				PutItemWithContext: test.PutItemInput("add_goal_channel_id", &model.Model{
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
			args: args{event: events.APIGatewayProxyRequest{Body: test.MakePayload(test.Add2ndGoalSubmissionPayload)}},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
				Body:       test.ReadFile(t, "interaction/add_2nd_goal_response.json"),
			},
			wantErr: false,
			wantDo:  do{},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: test.GetItemOutput("add_2nd_goal_channel_id", &model.Model{
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
				GetItemWithContext: test.GetItemInput("add_2nd_goal_channel_id"),
				PutItemWithContext: test.PutItemInput("add_2nd_goal_channel_id", &model.Model{
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
			args:     args{event: events.APIGatewayProxyRequest{Body: test.MakePayload(test.RemoveGoalActionPayload)}},
			response: test.SuccessResponse(test.ReadFile(t, "slash/slash_response.json")),
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:     test.ParseUrl("https://slack.com/api/views.update"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-Type": []string{"application/json"}},
				body:    test.ReadFile(t, "interaction/edit_goals_modal_with_goal_removed.json"),
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: test.GetItemOutput("remove_goal_channel_id", &model.Model{
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
				GetItemWithContext: test.GetItemInput("remove_goal_channel_id"),
				PutItemWithContext: test.PutItemInput("remove_goal_channel_id", &model.Model{
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
			args:     args{event: events.APIGatewayProxyRequest{Body: test.MakePayload(test.CloseEditGoalsPayload)}},
			response: test.SuccessResponse(test.ReadFile(t, "slash/slash_response.json")),
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers:    JsonHeaders(),
			},
			wantErr: false,
			wantDo: do{
				url:     test.ParseUrl("https://slack.com/api/views.update"),
				headers: http.Header{"Authorization": []string{"Bearer token_token"}, "Content-Type": []string{"application/json"}},
				body:    test.ReadFile(t, "interaction/summary_modal_with_added_goals_body.json"),
			},
			dynamoResponses: &mocks.MockDynamoDB{
				GetItemOutput: test.GetItemOutput("close_edit_goals_channel_id", &model.Model{
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
				GetItemWithContext: test.GetItemInput("close_edit_goals_channel_id"),
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
			mocks.ResetMockDynamoDbCalls()
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

	}
}