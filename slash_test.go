package bZapp

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/model"
	"github.com/dctid/bZapp/test"
	"github.com/dctid/bZapp/view"
	"github.com/slack-go/slack"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/dctid/bZapp/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-lambda-go/events"
)

func TestSlash_initialRequestInChannel(t *testing.T) {
	defer mocks.ResetMockDynamoDbCalls()
	var urlCalled *url.URL = nil
	var bodyCalled string
	expectUrl, _ := url.Parse("https://slack.com/api/views.open")
	Client = &mocks.MockClient{}

	DynamoDB = &mocks.MockDynamoDB{
		GetItemOutput: &dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
			},
		},
		PutItemOutput: &dynamodb.PutItemOutput{},
	}

	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		log.Printf("url %s ", req.URL)
		urlCalled = req.URL
		body, _ := ioutil.ReadAll(req.Body)
		bodyCalled = string(body)
		response := test.ReadFile(t, "slash/slash_response.json")
		return &http.Response{
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(response))),
			StatusCode: 200,
		}, nil
	}

	result, err := Slash(context.Background(), events.APIGatewayProxyRequest{
		Body: test.UrlEncode(slackCommand),
	})
	assert.NoError(t, err)
	assert.EqualValues(t, expectUrl, urlCalled)

	prettyJsonExpected := test.ReadFile(t,"slash/slash_init_expected.json")
	prettyJsonActual := format.PrettyJson(t, bodyCalled)
	assert.NoError(t, err)
	assert.EqualValues(t, prettyJsonExpected, prettyJsonActual)
	assert.EqualValues(t, events.APIGatewayProxyResponse{StatusCode: 200}, result)
	actualPutCall := mocks.MockDynamoDbCalls.PutItemWithContext
	currentModel := model.Model{Events: model.Events{}, Goals: model.Goals{}}
	modelBytes, err := json.Marshal(currentModel)
	expectedPutCall := &dynamodb.PutItemInput{
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
	assert.EqualValues(t, expectedPutCall, actualPutCall)
	expectedGetItemInput := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("D7P4LC5G9"),
			},
		},
		TableName: aws.String("bZappTable"),
	}
	actualGetItemInput := mocks.MockDynamoDbCalls.GetItemWithContext
	assert.EqualValues(t, expectedGetItemInput, actualGetItemInput)
}

func TestSlash_appExistsInChannel(t *testing.T) {
	defer mocks.ResetMockDynamoDbCalls()
	var urlCalled *url.URL = nil
	var bodyCalled string
	expectUrl, _ := url.Parse("https://slack.com/api/views.open")
	Client = &mocks.MockClient{}
	model.Clock = mocks.NewMockClock("2020-12-02 08:48:21")
	currentModel := model.Model{
		Events: model.Events{
			"2020-12-02": []model.Event{{
				Id:    "today_event",
				Title: "Let's do something",
				Day:   view.TodayOptionValue,
				Hour:  9,
				Min:   15,
				AmPm:  "AM",
			}},
			"2020-12-03": []model.Event{{
				Id:    "tomorrow_event",
				Title: "Let's do something else",
				Day:   view.TomorrowOptionValue,
				Hour:  3,
				Min:   30,
				AmPm:  "PM",
			}},
		},
		Goals: model.Goals{
			"someGoal": []model.Goal{
				{Id: "goal_id", Value: "the goal"},
			},
		},
	}
	modelBytes, err := json.Marshal(currentModel)

	DynamoDB = &mocks.MockDynamoDB{
		GetItemOutput: &dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String("D7P4LC5G9"),
				},
				"model": {
					S: aws.String(string(modelBytes)),
				},
			},
		},
	}


	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		log.Printf("url %s ", req.URL)
		urlCalled = req.URL
		body, _ := ioutil.ReadAll(req.Body)
		bodyCalled = string(body)
		response := test.ReadFile(t, "slash/slash_response.json")
		return &http.Response{
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(response))),
			StatusCode: 200,
		}, nil
	}

	result, err := Slash(context.Background(), events.APIGatewayProxyRequest{
		Body: test.UrlEncode(slackCommand),
	})
	assert.NoError(t, err)
	assert.EqualValues(t, expectUrl, urlCalled)

	prettyJsonExpected := test.ReadFile(t, "slash/slash_existing_channel.json")
	prettyJsonActual := format.PrettyJson(t, bodyCalled)
	assert.NoError(t, err)
	assert.EqualValues(t, prettyJsonExpected, prettyJsonActual)
	assert.EqualValues(t, events.APIGatewayProxyResponse{StatusCode: 200}, result)
	expectedGetItemInput := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("D7P4LC5G9"),
			},
		},
		TableName: aws.String("bZappTable"),
	}
	actualGetItemInput := mocks.MockDynamoDbCalls.GetItemWithContext
	assert.EqualValues(t, expectedGetItemInput, actualGetItemInput)
}

func TestSlash_appExistsInChannel_onFriday(t *testing.T) {
	defer mocks.ResetMockDynamoDbCalls()
	var urlCalled *url.URL = nil
	var bodyCalled string
	expectUrl, _ := url.Parse("https://slack.com/api/views.open")
	Client = &mocks.MockClient{}
	model.Clock = mocks.NewMockClock("2020-12-04 08:48:21")
	currentModel := model.Model{
		Events: model.Events{
			"2020-12-03": []model.Event{{
				Id:    "yesterday_event",
				Title: "Let's do something",
				Day:   view.TodayOptionValue,
				Hour:  9,
				Min:   15,
				AmPm:  "AM",
			}},
			"2020-12-04": []model.Event{{
				Id:    "today_event",
				Title: "Let's do something else",
				Day:   view.TomorrowOptionValue,
				Hour:  3,
				Min:   30,
				AmPm:  "PM",
			}},
		},
		Goals: model.Goals{
			"someGoal": []model.Goal{
				{Id: "goal_id", Value: "the goal"},
			},
		},
	}
	modelBytes, err := json.Marshal(currentModel)

	DynamoDB = &mocks.MockDynamoDB{
		GetItemOutput: &dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String("D7P4LC5G9"),
				},
				"model": {
					S: aws.String(string(modelBytes)),
				},
			},
		},
	}


	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		log.Printf("url %s ", req.URL)
		urlCalled = req.URL
		body, _ := ioutil.ReadAll(req.Body)
		bodyCalled = string(body)
		response := test.ReadFile(t, "slash/slash_response.json")
		return &http.Response{
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(response))),
			StatusCode: 200,
		}, nil
	}

	result, err := Slash(context.Background(), events.APIGatewayProxyRequest{
		Body: test.UrlEncode(slackCommand),
	})
	assert.NoError(t, err)
	assert.EqualValues(t, expectUrl, urlCalled)

	prettyJsonExpected := test.ReadFile(t, "slash/slash_existing_channel_on_friday.json")
	prettyJsonActual := format.PrettyJson(t, bodyCalled)
	assert.NoError(t, err)
	assert.EqualValues(t, prettyJsonExpected, prettyJsonActual)
	assert.EqualValues(t, events.APIGatewayProxyResponse{StatusCode: 200}, result)
	expectedGetItemInput := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("D7P4LC5G9"),
			},
		},
		TableName: aws.String("bZappTable"),
	}
	actualGetItemInput := mocks.MockDynamoDbCalls.GetItemWithContext
	assert.EqualValues(t, expectedGetItemInput, actualGetItemInput)
}

var slackCommand = slack.SlashCommand{
	Token:       "8KTh0sVRkeZozlTxrBRqk1NO",
	TeamID:      "T7NS02BFB",
	TeamDomain:  "ford-community",
	ChannelID:   "D7P4LC5G9",
	ChannelName: "directmessage",
	UserID:      "U7QNBA36K",
	UserName:    "cdorman1",
	Command:     "/bzapp",
	ResponseURL: "https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP",
	TriggerID:   "1282571347205.260884079521.45166c59ef86cfcf9409d2ec2d4b4a58",
}

