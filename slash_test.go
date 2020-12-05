package bZapp

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/model"
	"github.com/dctid/bZapp/view"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/dctid/bZapp/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-lambda-go/events"
)

const initExpected = `{
  "trigger_id": "1282571347205.260884079521.45166c59ef86cfcf9409d2ec2d4b4a58",
  "view": {
    "type": "modal",
    "title": {
      "type": "plain_text",
      "text": "bZapp",
      "emoji": true
    },
    "private_metadata": "{\"channel_id\":\"D7P4LC5G9\"}",
    "submit": {
      "type": "plain_text",
      "text": "Submit",
      "emoji": true
    },
    "close": {
      "type": "plain_text",
      "text": "Cancel",
      "emoji": true
    },
    "blocks": [
      {
        "type": "divider"
      },
      {
		"text": {
			"text": "Events",
			"type": "plain_text"
		},
		"type": "header"
      },
      {
        "type": "section",
        "text": {
          "type": "mrkdwn",
          "text": "_No events yet_"
        }
      },
      {
        "type": "divider"
      },
      {
		"text": {
			"text": "Goals",
			"type": "plain_text"
		},
		"type": "header"
      },
      {
        "text": {
          "text": "_Nogoalsyet_",
          "type": "mrkdwn"
        },
        "type": "section"
      },
      {
        "type": "divider"
      },
      {
        "type": "actions",
        "block_id": "actions_block",
        "elements": [
          {
            "type": "button",
            "action_id": "edit_events",
            "text": {
              "type": "plain_text",
              "text": "Edit Events",
              "emoji": true
            },
            "value": "edit_events"
          },
          {
            "action_id": "edit_goals",
            "text": {
              "emoji": true,
              "text": "EditGoals",
              "type": "plain_text"
            },
            "type": "button",
            "value": "edit_goals"
          }
        ]
      }
    ]
  }
}`

const existingExpected = `{
  "trigger_id": "1282571347205.260884079521.45166c59ef86cfcf9409d2ec2d4b4a58",
  "view": {
	"type": "modal",
	"title": {
		"type": "plain_text",
		"text": "bZapp",
		"emoji": true
	},
	"blocks": [
		{
			"type": "divider"
		},
		{
			"type": "header",
			"text": {
				"type": "plain_text",
				"text": "Events"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Today*"
				}
			]
		},
		{
			"type": "divider"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": ":small_orange_diamond: 9:15 Let's do something"
			},
			"block_id": "today_event"
		},
		{
			"type": "divider"
		},
		{
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Tomorrow*"
				}
			]
		},
		{
			"type": "divider"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": ":small_orange_diamond: 3:30 Let's do something else"
			},
			"block_id": "tomorrow_event"
		},
		{
			"type": "divider"
		},
		{
			"type": "header",
			"text": {
				"type": "plain_text",
				"text": "Goals"
			}
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "_No goals yet_"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "actions",
			"block_id": "actions_block",
			"elements": [
				{
					"type": "button",
					"text": {
						"type": "plain_text",
						"text": "Edit Events",
						"emoji": true
					},
					"action_id": "edit_events",
					"value": "edit_events"
				},
				{
					"type": "button",
					"text": {
						"type": "plain_text",
						"text": "Edit Goals",
						"emoji": true
					},
					"action_id": "edit_goals",
					"value": "edit_goals"
				}
			]
		}
	],
	"close": {
		"type": "plain_text",
		"text": "Cancel",
		"emoji": true
	},
	"submit": {
		"type": "plain_text",
		"text": "Submit",
		"emoji": true
	},
	"private_metadata": "{\"channel_id\":\"D7P4LC5G9\"}"
}
}`
const existingExpectedOnFriday = `{
  "trigger_id": "1282571347205.260884079521.45166c59ef86cfcf9409d2ec2d4b4a58",
  "view": {
	"type": "modal",
	"title": {
		"type": "plain_text",
		"text": "bZapp",
		"emoji": true
	},
	"blocks": [
		{
			"type": "divider"
		},
		{
			"type": "header",
			"text": {
				"type": "plain_text",
				"text": "Events"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "context",
			"elements": [
				{
					"type": "mrkdwn",
					"text": "*Today*"
				}
			]
		},
		{
			"type": "divider"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": ":small_orange_diamond: 3:30 Let's do something else"
			},
			"block_id": "today_event"
		},
		{
			"type": "divider"
		},
		{
			"type": "header",
			"text": {
				"type": "plain_text",
				"text": "Goals"
			}
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "_No goals yet_"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "actions",
			"block_id": "actions_block",
			"elements": [
				{
					"type": "button",
					"text": {
						"type": "plain_text",
						"text": "Edit Events",
						"emoji": true
					},
					"action_id": "edit_events",
					"value": "edit_events"
				},
				{
					"type": "button",
					"text": {
						"type": "plain_text",
						"text": "Edit Goals",
						"emoji": true
					},
					"action_id": "edit_goals",
					"value": "edit_goals"
				}
			]
		}
	],
	"close": {
		"type": "plain_text",
		"text": "Cancel",
		"emoji": true
	},
	"submit": {
		"type": "plain_text",
		"text": "Submit",
		"emoji": true
	},
	"private_metadata": "{\"channel_id\":\"D7P4LC5G9\"}"
}
}`


 const response = `{
	"ok": true,
	"error": "",
	"view": {
		"ok": false,
		"error": "",
		"id": "V018C5094DQ",
		"team_id": "T7NS02BFB",
		"type": "view",
		"title": {
			"type": "plain_text",
			"text": "bZapp",
			"emoji": true
		},
		"close": {
			"type": "plain_text",
			"text": "Cancel",
			"emoji": true
		},
		"submit": {
			"type": "plain_text",
			"text": "Submit",
			"emoji": true
		},
		"blocks": [
			{
				"type": "divider",
				"block_id": "kTDEm"
			},
			{
				"type": "context",
				"block_id": "zzEF",
				"elements": [
					{
						"type": "mrkdwn",
						"text": "*Today's Events*"
					}
				]
			},
			{
				"type": "divider",
				"block_id": "X=862"
			},
			{
				"type": "section",
				"text": {
					"type": "mrkdwn",
					"text": "No events yet"
				},
				"block_id": "F+41T"
			},
			{
				"type": "divider",
				"block_id": "s6+nf"
			},
			{
				"type": "context",
				"block_id": "jD+",
				"elements": [
					{
						"type": "mrkdwn",
						"text": "*Tomorrow's Events*"
					}
				]
			},
			{
				"type": "divider",
				"block_id": "e5z8"
			},
			{
				"type": "section",
				"text": {
					"type": "mrkdwn",
					"text": "No events yet"
				},
				"block_id": "IVdv"
			},
			{
				"type": "divider",
				"block_id": "X6E"
			},
			{
				"type": "input",
				"block_id": "OM=8",
				"label": {
					"type": "plain_text",
					"text": "Add Event",
					"emoji": true
				},
				"element": {
					"type": "plain_text_input",
					"action_id": "add_event",
					"placeholder": {
						"type": "plain_text",
						"text": "Title",
						"emoji": true
					}
				}
			},
			{
				"type": "actions",
				"block_id": "pR5OI",
				"elements": [
					{
						"type": "static_select",
						"placeholder": {
							"type": "plain_text",
							"text": "Select hour",
							"emoji": true
						},
						"action_id": "hours_select",
						"options": [
							{
								"text": {
									"type": "plain_text",
									"text": "9 AM",
									"emoji": true
								},
								"value": "hour-9"
							},
							{
								"text": {
									"type": "plain_text",
									"text": "10 AM",
									"emoji": true
								},
								"value": "hour-10"
							},
							{
								"text": {
									"type": "plain_text",
									"text": "11 AM",
									"emoji": true
								},
								"value": "hour-11"
							},
							{
								"text": {
									"type": "plain_text",
									"text": "12 PM",
									"emoji": true
								},
								"value": "hour-12"
							},
							{
								"text": {
									"type": "plain_text",
									"text": "1 PM",
									"emoji": true
								},
								"value": "hour-1"
							},
							{
								"text": {
									"type": "plain_text",
									"text": "2 PM",
									"emoji": true
								},
								"value": "hour-2"
							},
							{
								"text": {
									"type": "plain_text",
									"text": "3 PM",
									"emoji": true
								},
								"value": "hour-3"
							},
							{
								"text": {
									"type": "plain_text",
									"text": "4 PM",
									"emoji": true
								},
								"value": "hour-4"
							}
						]
					},
					{
						"type": "static_select",
						"placeholder": {
							"type": "plain_text",
							"text": "Select minutes",
							"emoji": true
						},
						"action_id": "mins_select",
						"options": [
							{
								"text": {
									"type": "plain_text",
									"text": "00",
									"emoji": true
								},
								"value": "min-0"
							},
							{
								"text": {
									"type": "plain_text",
									"text": "15",
									"emoji": true
								},
								"value": "min-15"
							},
							{
								"text": {
									"type": "plain_text",
									"text": "30",
									"emoji": true
								},
								"value": "min-30"
							},
							{
								"text": {
									"type": "plain_text",
									"text": "45",
									"emoji": true
								},
								"value": "min-45"
							}
						]
					},
					{
						"type": "datepicker",
						"action_id": "datepicker",
						"placeholder": {
							"type": "plain_text",
							"text": "Select a date",
							"emoji": true
						}
					},
					{
						"type": "button",
						"text": {
							"type": "plain_text",
							"text": "Add",
							"emoji": true
						},
						"action_id": "lk/mR",
						"value": "add_event"
					}
				]
			}
		],
		"private_metadata": "",
		"callback_id": "",
		"state": {
			"values": {}
		},
		"hash": "1596660783.uygR3WMh",
		"clear_on_close": false,
		"notify_on_close": false,
		"root_view_id": "V018C5094DQ",
		"previous_view_id": "",
		"app_id": "A0131JT7VPF",
		"external_id": "",
		"bot_id": "B0133F8RE11"
	}
}`


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

	prettyJsonExpected, err := format.PrettyJson(initExpected)
	assert.NoError(t, err)

	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		log.Printf("url %s ", req.URL)
		urlCalled = req.URL
		body, _ := ioutil.ReadAll(req.Body)
		bodyCalled = string(body)
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewReader([]byte(response))),
			StatusCode: 200,
		}, nil
	}

	result, err := Slash(context.Background(), events.APIGatewayProxyRequest{
		Body: encodedBody,
	})
	assert.NoError(t, err)
	assert.EqualValues(t, expectUrl, urlCalled)

	prettyJsonActual, err := format.PrettyJson(bodyCalled)
	assert.NoError(t, err)
	assert.EqualValues(t, prettyJsonExpected, prettyJsonActual)
	assert.EqualValues(t, events.APIGatewayProxyResponse{StatusCode: 200}, result)
	actualPutCall := mocks.MockDynamoDbCalls.PutItemWithContext
	currentModel := model.Model{ChannelId: "D7P4LC5G9", Events: model.Events{}, Goals: model.Goals{}}
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
		ChannelId: "D7P4LC5G9",
		Events: model.Events{
			"2020-12-02":    []model.Event{{
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
				"model" : {
					S: aws.String(string(modelBytes)),
				},
			},
		},
	}

	prettyJsonExpected, err := format.PrettyJson(existingExpected)
	assert.NoError(t, err)

	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		log.Printf("url %s ", req.URL)
		urlCalled = req.URL
		body, _ := ioutil.ReadAll(req.Body)
		bodyCalled = string(body)
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewReader([]byte(response))),
			StatusCode: 200,
		}, nil
	}

	result, err := Slash(context.Background(), events.APIGatewayProxyRequest{
		Body: encodedBody,
	})
	assert.NoError(t, err)
	assert.EqualValues(t, expectUrl, urlCalled)

	prettyJsonActual, err := format.PrettyJson(bodyCalled)
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
		ChannelId: "D7P4LC5G9",
		Events: model.Events{
			"2020-12-03":    []model.Event{{
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
				"model" : {
					S: aws.String(string(modelBytes)),
				},
			},
		},
	}

	prettyJsonExpected, err := format.PrettyJson(existingExpectedOnFriday)
	assert.NoError(t, err)

	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		log.Printf("url %s ", req.URL)
		urlCalled = req.URL
		body, _ := ioutil.ReadAll(req.Body)
		bodyCalled = string(body)
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewReader([]byte(response))),
			StatusCode: 200,
		}, nil
	}

	result, err := Slash(context.Background(), events.APIGatewayProxyRequest{
		Body: encodedBody,
	})
	assert.NoError(t, err)
	assert.EqualValues(t, expectUrl, urlCalled)

	prettyJsonActual, err := format.PrettyJson(bodyCalled)
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

var encodedBody = `token=8KTh0sVRkeZozlTxrBRqk1NO&team_id=T7NS02BFB&team_domain=ford-community&channel_id=D7P4LC5G9&channel_name=directmessage&user_id=U7QNBA36K&user_name=cdorman1&command=%2Fbzapp&text=&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FT7NS02BFB%2F1307783467168%2FGvz9lFVBwn9xo8TweP2vJHsP&trigger_id=1282571347205.260884079521.45166c59ef86cfcf9409d2ec2d4b4a58`