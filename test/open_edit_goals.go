package test

import "github.com/slack-go/slack"

var EditGoalsActionButtonPayload = slack.InteractionCallback{
	Type:      "block_actions",
	TriggerID: "Edit Goals Trigger",
	ActionCallback: slack.ActionCallbacks{
		BlockActions: []*slack.BlockAction{{
			ActionID: "edit_goals",
		}},
	},
	View: slack.View{
		PrivateMetadata: `{"channel_id":"D7P4LC5G9","response_url":"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"}`,
	},
}


const EditGoalsModal = `{
		"blocks": [
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
				"block_id": "add_goal_category_input_block-1",
				"element": {
					"action_id": "add_goal_category",
					"options": [
						{
							"text": {
								"emoji": true,
								"text": "CustomerQuestions?",
								"type": "plain_text"
							},
							"value": "CustomerQuestions?"
						},
						{
							"text": {
								"emoji": true,
								"text": "TeamNeeds",
								"type": "plain_text"
							},
							"value": "TeamNeeds"
						},
						{
							"text": {
								"emoji": true,
								"text": "Learnings",
								"type": "plain_text"
							},
							"value": "Learnings"
						},
						{
							"text": {
								"emoji": true,
								"text": "Questions?",
								"type": "plain_text"
							},
							"value": "Questions?"
						},
						{
							"text": {
								"emoji": true,
								"text": "Other",
								"type": "plain_text"
							},
							"value": "Other"
						}
					],
					"placeholder": {
						"emoji": true,
						"text": "ChooseGoal",
						"type": "plain_text"
					},
					"type": "static_select"
				},
				"label": {
					"emoji": true,
					"text": "GoaltoAddto",
					"type": "plain_text"
				},
				"type": "input"
			},
			{
				"block_id": "add_goal_input_block-1",
				"element": {
					"action_id": "add_goal",
					"placeholder": {
						"text": "Goal",
						"type": "plain_text"
					},
					"type": "plain_text_input"
				},
				"label": {
					"emoji": true,
					"text": "GoaltoAdd",
					"type": "plain_text"
				},
				"type": "input"
			}
		],
		"close": {
			"emoji": true,
			"text": "Back",
			"type": "plain_text"
		},
		"notify_on_close": true,
		"private_metadata": "{\"channel_id\":\"D7P4LC5G9\",\"response_url\":\"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP\"}",
		"submit": {
			"emoji": true,
			"text": "Add",
			"type": "plain_text"
		},
		"title": {
			"emoji": true,
			"text": "bZapp-EditGoals",
			"type": "plain_text"
		},
		"type": "modal"
	}`
