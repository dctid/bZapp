package view

import (
	"encoding/json"
	"github.com/dctid/bZapp/format"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"testing"
)

var editEventsModal = `{
	"title": {
		"type": "plain_text",
		"text": "bZapp - Edit Events",
		"emoji": true
	},
	"notify_on_close": true,
	"submit": {
		"type": "plain_text",
		"text": "Add",
		"emoji": true
	},
	"type": "modal",
	"close": {
		"type": "plain_text",
		"text": "Back",
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
				"text": "9:15 Standup"
			},
			"accessory": {
				"action_id": "remove_today_1_action_id",
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"value": "remove_today_1"
			},
			"block_id": "today_1"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "11:30 IPM"
			},
			"accessory": {
				"action_id": "remove_today_2_action_id",
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"value": "remove_today_2"
			},
			"block_id": "today_2"
		},
		{
			"type": "section",
			"block_id": "today_3",
			"text": {
				"type": "mrkdwn",
				"text": "3:15 Retro"
			},
			"accessory": {
				"action_id": "remove_today_3_action_id",
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"value": "remove_today_3"
			},
			"block_id": "today_3"
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
				"text": "9:15 Standup"
			},
			"accessory": {
				"action_id": "remove_tomorrow_1_action_id",
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"value": "remove_tomorrow_1"
			},
			"block_id": "tomorrow_1"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "1:30 User Interview"
			},
			"accessory": {
				"action_id": "remove_tomorrow_2_action_id",
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"value": "remove_tomorrow_2"
			},
			"block_id": "tomorrow_2"
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "3:00 Synthesis"
			},
			"accessory": {
				"action_id": "remove_tomorrow_3_action_id",
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Remove",
					"emoji": true
				},
				"value": "remove_tomorrow_3"
			},
			"block_id": "tomorrow_3"
		},
		{
			"type": "divider"
		},
		{
			"type": "input",
			"block_id": "add_event_title_input_block-666",
			"element": {
				"action_id": "add_event_title",
				"type": "plain_text_input",
				"placeholder": {
					"type": "plain_text",
					"text": "Title"
				}
			},
			"label": {
				"type": "plain_text",
				"text": "Add Event"
			}
		},
		{
			"type": "input",
			"block_id": "add_event_day_input_block-666",
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
			},
			"label": {
				"type": "plain_text",
				"text": "Day",
				"emoji": true
			}
		},
		{
			"type": "input",
			"block_id": "add_event_hours_input_block-666",
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
			},
			"label": {
				"type": "plain_text",
				"text": "Hour",
				"emoji": true
			}
		},
		{
			"type": "input",
			"block_id": "add_event_mins_input_block-666",
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
			},
			"label": {
				"type": "plain_text",
				"text": "Minutes",
				"emoji": true
			}
		}
	]
}`

func TestNewEditEventsModal(t *testing.T) {

	//expected := slack.ModalViewRequest{}

	//expectedJson, _ := json.Marshal(editEventsModal)
	standup := slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "9:15 Standup", false, false), nil, slack.NewAccessory(slack.NewButtonBlockElement("remove_today_1_action_id", "remove_today_1", slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false))), slack.SectionBlockOptionBlockID("today_1"),
	)
	synthesis := slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "11:30 IPM", false, false), nil, slack.NewAccessory(slack.NewButtonBlockElement("remove_today_2_action_id", "remove_today_2", slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false))), slack.SectionBlockOptionBlockID("today_2"),
	)
	retro := slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "3:15 Retro", false, false), nil, slack.NewAccessory(slack.NewButtonBlockElement("remove_today_3_action_id", "remove_today_3", slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false))), slack.SectionBlockOptionBlockID("today_3"),
	)

	todaysEvents := []slack.Block{standup, synthesis, retro}

	tomorrowStandup := slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "9:15 Standup", false, false), nil, slack.NewAccessory(slack.NewButtonBlockElement("remove_tomorrow_1_action_id", "remove_tomorrow_1", slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false))), slack.SectionBlockOptionBlockID("tomorrow_1"),
	)
	userInterview := slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "1:30 User Interview", false, false), nil, slack.NewAccessory(slack.NewButtonBlockElement("remove_tomorrow_2_action_id", "remove_tomorrow_2", slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false))), slack.SectionBlockOptionBlockID("tomorrow_2"),
	)
	tomorrowsSynthesis := slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "3:00 Synthesis", false, false), nil, slack.NewAccessory(slack.NewButtonBlockElement("remove_tomorrow_3_action_id", "remove_tomorrow_3", slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false))), slack.SectionBlockOptionBlockID("tomorrow_3"),
	)

	 tomorrowsEvents := []slack.Block{
		tomorrowStandup, userInterview, tomorrowsSynthesis,
	}

	result := NewEditEventsModal(666, todaysEvents, tomorrowsEvents)
	actualJson, _ := json.Marshal(result)
	//assert.EqualValues(t, expected, result)
	expectedJsonString, _ := format.PrettyJson(editEventsModal)
	actualJsonString, _ := format.PrettyJson(string(actualJson))

	assert.EqualValues(t, expectedJsonString, actualJsonString)
}
