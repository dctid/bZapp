package test

import "github.com/slack-go/slack"

var SubmitPayload = slack.InteractionCallback{
	Type:      "view_submission",
	View: slack.View{
		Title: &slack.TextBlockObject{Text: "bZapp"},
		PrivateMetadata: `{"channel_id":"D7P4LC5G9","response_url":"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"}`,
	},
}

const SubmissionJson = `{
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
			"block_id": "coEbHc",
			"text": {
				"text": ":small_orange_diamond: 9:00asdf",
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
			"block_id": "DZosTr",
			"text": {
				"text": ":small_orange_diamond: 1:15qewr",
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
			"block_id": "bSAnHN",
			"text": {
				"text": ":small_blue_diamond: sfd",
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
					"text": "*Questions?*",
					"type": "mrkdwn"
				}
			],
			"type": "context"
		},
		{
			"type": "divider"
		},
		{
			"block_id": "RrMdIA",
			"text": {
				"text": ":small_blue_diamond: afsasdf",
				"type": "mrkdwn"
			},
			"type": "section"
		}
	],
	"delete_original": false,
	"replace_original": false,
	"response_type": "in_channel",
	"text": "bZapp-Today'sStandupSummary"
}`
