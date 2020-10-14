package bZapp

import (
	"bytes"
	"context"
	"github.com/dctid/bZapp/format"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/dctid/bZapp/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-lambda-go/events"
)

const expected = `{
	"trigger_id": "1282571347205.260884079521.45166c59ef86cfcf9409d2ec2d4b4a58",
    "view": {
	  "type": "modal",
	  "title": {
		"type": "plain_text",
		"text": "bZapp",
		"emoji": true
	  },
      "private_metadata": "{\"Index\":0,\"Events\":{\"TodaysEvents\":null,\"TomorrowsEvents\":null},\"Goals\":null}",
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
		  "type": "context",
		  "elements": [
			{
			  "type": "mrkdwn",
			  "text": "*Today's Events*"
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
			"text": "_No events yet_"
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
			  "text": "*Tomorrow's Events*"
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
			"text": "_No events yet_"
		  }
		},
		{
			"type": "divider"
		},
		{
			"elements": [
           		{
           			"text": "Goals",
           			"type": "mrkdwn"
           		}
           	],
           	"type": "context"
		},
		{
			"type": "divider"
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
		},
		{
			"block_id": "convo_input_id",
			"element": {
				"action_id": "conversation_select_action_id",
				"default_to_current_conversation": true,
				"response_url_enabled": true,
				"type": "conversations_select"
			},
			"label": {
				"text": "Selectachanneltoposttheresulton",
				"type": "plain_text"
			},
			"type": "input"
		}
	  ]
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

func TestSlash(t *testing.T) {

	var urlCalled *url.URL = nil
	var bodyCalled string
	expectUrl, _ := url.Parse("https://slack.com/api/views.open")
	Client = &mocks.MockClient{}

	prettyJsonExpected, err := format.PrettyJson(expected)
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
}

var encodedBody = `token=8KTh0sVRkeZozlTxrBRqk1NO&team_id=T7NS02BFB&team_domain=ford-community&channel_id=D7P4LC5G9&channel_name=directmessage&user_id=U7QNBA36K&user_name=cdorman1&command=%2Fbzapp&text=&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FT7NS02BFB%2F1307783467168%2FGvz9lFVBwn9xo8TweP2vJHsP&trigger_id=1282571347205.260884079521.45166c59ef86cfcf9409d2ec2d4b4a58`