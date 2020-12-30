package test

import (
	"github.com/slack-go/slack"
)

var CloseEditGoalsPayload = slack.InteractionCallback{
	Type:      "view_closed",
	View: slack.View{
		ID: "V01DBFTR588",
		RootViewID: "V01DBEZ35GQ",
		ExternalID: "Close_goals_externalId",
		Title: &slack.TextBlockObject{Text: "bZapp - Edit Goals"},
		PrivateMetadata: `{"channel_id":"D7P4LC5G9","response_url":"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"}`,
	},
	ViewSubmissionCallback: slack.ViewSubmissionCallback{Hash: "close_goals_hash"},
}

const SummaryModalWithGoals = `{
	"external_id": "Close_goals_externalId",
	"hash": "close_goals_hash",
	"view": {
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
				"elements": [
					{
						"text": "*Today*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
			},
			{
				"type": "divider"
			},
			{
				"block_id": "WyKVYV",
				"text": {
					"text": ":small_orange_diamond: 10:00ads",
					"type": "mrkdwn"
				},
				"type": "section"
			},
			{
				"type": "divider"
			},
			{
				"elements": [
					{
						"text": "*Tomorrow*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
			},
			{
				"type": "divider"
			},
			{
				"block_id": "PTjSgI",
				"text": {
					"text": ":small_orange_diamond: 11:15dfs",
					"type": "mrkdwn"
				},
				"type": "section"
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
				"type": "divider"
			},
			{
				"elements": [
					{
						"text": "*TeamNeeds*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
			},
			{
				"type": "divider"
			},
			{
				"block_id": "YbiWhf",
				"text": {
					"text": ":small_blue_diamond: lskfd",
					"type": "mrkdwn"
				},
				"type": "section"
			},
			{
				"type": "divider"
			},
			{
				"elements": [
					{
						"text": "*Learnings*",
						"type": "mrkdwn"
					}
				],
				"type": "context"
			},
			{
				"type": "divider"
			},
			{
				"block_id": "mopNVQ",
				"text": {
					"text": ":small_blue_diamond: sdfg",
					"type": "mrkdwn"
				},
				"type": "section"
			},
			{
				"type": "divider"
			},
			{
				"block_id": "actions_block",
				"elements": [
					{
						"action_id": "edit_events",
						"text": {
							"emoji": true,
							"text": "EditEvents",
							"type": "plain_text"
						},
						"type": "button",
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
				],
				"type": "actions"
			}
		],
		"close": {
			"emoji": true,
			"text": "Cancel",
			"type": "plain_text"
		},
		"private_metadata": "{\"channel_id\":\"D7P4LC5G9\",\"response_url\":\"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP\"}",
		"submit": {
			"emoji": true,
			"text": "Submit",
			"type": "plain_text"
		},
		"title": {
			"emoji": true,
			"text": "bZapp",
			"type": "plain_text"
		},
		"type": "modal"
	},
	"view_id": "V01DBEZ35GQ"
}`
