package view

import (
	"encoding/json"
	"github.com/dctid/bZapp/format"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"testing"
)

var summaryModal = `{
	"title": {
		"type": "plain_text",
		"text": "bZapp",
		"emoji": true
	},
	"submit": {
		"type": "plain_text",
		"text": "Submit",
		"emoji": true
	},
	"type": "modal",
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
				"text": "9:15 Standup"
			}
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "11:30 IPM"
			}
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "3:15 Retro"
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
				"text": "9:15 Standup"
			}
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "1:30 User Interview"
			}
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "3:00 Synthesis"
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
			  "action_id": "edit_events",
			  "text": {
				"type": "plain_text",
				"text": "Edit Events",
				"emoji": true
			  },
			  "value": "edit_events"
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
}`

func TestNewSummaryModal(t *testing.T) {

	//expected := slack.ModalViewRequest{}

	//expectedJson, _ := json.Marshal(editEventsModal)
	standup := slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "9:15 Standup", false, false), nil, nil,
	)
	synthesis := slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "11:30 IPM", false, false), nil, nil,
	)
	retro := slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "3:15 Retro", false, false), nil, nil,
	)

	todaysEvents := []*slack.SectionBlock{standup, synthesis, retro}

	tomorrowStandup := slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "9:15 Standup", false, false), nil, nil,
	)
	userInterview := slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "1:30 User Interview", false, false), nil, nil,
	)
	tomorrowsSynthesis := slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "3:00 Synthesis", false, false), nil, nil,
	)

	 tomorrowsEvents := []*slack.SectionBlock{
		tomorrowStandup, userInterview, tomorrowsSynthesis,
	}

	result := NewSummaryModal(todaysEvents, tomorrowsEvents)
	actualJson, _ := json.Marshal(result)
	//assert.EqualValues(t, expected, result)
	expectedJsonString, _ := format.PrettyJson(summaryModal)
	actualJsonString, _ := format.PrettyJson(string(actualJson))
	assert.EqualValues(t, expectedJsonString, actualJsonString)
}