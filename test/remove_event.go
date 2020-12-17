package test

import "github.com/slack-go/slack"

var RemoveEventActionPayload = slack.InteractionCallback{
	Type:      "block_actions",
	ActionCallback: slack.ActionCallbacks{
		BlockActions: []*slack.BlockAction{{
			ActionID: "remove_event",
			Value: "remove_today_YUBFMb",
			BlockID: "YUBFMb",
		}},

	},
	View: slack.View{
		ID: "V01CMKMUWUS",
		ExternalID: "outsideId",
		PrivateMetadata: `{"channel_id":"D7P4LC5G9","response_url":"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"}`,
	},
	ViewSubmissionCallback: slack.ViewSubmissionCallback{Hash: "cornbeef"},
}

const RemoveEventsModal = `{
	"type": "modal",
	"title": {
		"type": "plain_text",
		"text": "bZapp - Edit Events",
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
				"text": ":small_orange_diamond: 10:00 ads"
			},
			"block_id": "WyKVYV",
			"accessory": {
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"action_id": "remove_event",
				"value": "remove_today_WyKVYV"
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
				"text": ":small_orange_diamond: 11:15 dfs"
			},
			"block_id": "PTjSgI",
			"accessory": {
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"action_id": "remove_event",
				"value": "remove_tomorrow_PTjSgI"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "input",
			"block_id": "add_event_title_input_block-4",
			"label": {
				"type": "plain_text",
				"text": "Add Event"
			},
			"element": {
				"type": "plain_text_input",
				"action_id": "add_event_title",
				"placeholder": {
					"type": "plain_text",
					"text": "Title"
				}
			}
		},
		{
			"type": "input",
			"block_id": "add_event_day_input_block-4",
			"label": {
				"type": "plain_text",
				"text": "Day",
				"emoji": true
			},
			"element": {
				"type": "radio_buttons",
				"action_id": "add_event_day",
				"options": [
					{
						"text": {
							"type": "plain_text",
							"text": "Today",
							"emoji": true
						},
						"value": "today"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "Tomorrow",
							"emoji": true
						},
						"value": "tomorrow"
					}
				]
			}
		},
		{
			"type": "input",
			"block_id": "add_event_hours_input_block-4",
			"label": {
				"type": "plain_text",
				"text": "Hour",
				"emoji": true
			},
			"element": {
				"type": "static_select",
				"placeholder": {
					"type": "plain_text",
					"text": "Select hour",
					"emoji": true
				},
				"action_id": "add_event_hour",
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
			}
		},
		{
			"type": "input",
			"block_id": "add_event_mins_input_block-4",
			"label": {
				"type": "plain_text",
				"text": "Minutes",
				"emoji": true
			},
			"element": {
				"type": "static_select",
				"placeholder": {
					"type": "plain_text",
					"text": "Select Minutes",
					"emoji": true
				},
				"action_id": "add_event_mins",
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
			}
		}
	],
	"close": {
		"type": "plain_text",
		"text": "Back",
		"emoji": true
	},
	"submit": {
		"type": "plain_text",
		"text": "Add",
		"emoji": true
	},
	"private_metadata": "{\"channel_id\":\"D7P4LC5G9\",\"response_url\":\"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP\"}",
	"notify_on_close": true
}
`
