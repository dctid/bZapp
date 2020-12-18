package test

import "github.com/slack-go/slack"

var AddGoalSubmissionPayload = slack.InteractionCallback{
	Type: "view_submission",
	View: slack.View{
		Title: &slack.TextBlockObject{Text: "bZapp - Edit Goals"},
		State: &slack.ViewState{Values: map[string]map[string]slack.BlockAction{
			"add_goal_category_input_block-6": {"add_goal_category": slack.BlockAction{SelectedOption: slack.OptionBlockObject{Value: "Team Needs"}}},
			"add_goal_input_block-6":          {"add_goal": slack.BlockAction{Value: "lskfd"}},
		}},
		PrivateMetadata: `{"channel_id":"D7P4LC5G9","response_url":"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"}`,
	},
}

const AddGoalSubmissionResponse = `{
  "response_action": "update",
  "view": {
    "type": "modal",
    "title": {
      "type": "plain_text",
      "text": "bZapp - Edit Goals",
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
            "text": "*Team Needs*"
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
          "text": ":small_blue_diamond: lskfd"
        },
        "block_id": "Fakehash",
        "accessory": {
          "type": "button",
          "text": {
            "type": "plain_text",
            "text": "Remove",
            "emoji": true
          },
          "action_id": "remove_goal",
          "value": "remove_Team Needs_Fakehash"
        }
      },
      {
        "type": "divider"
      },
      {
        "type": "input",
        "block_id": "add_goal_category_input_block-7",
        "label": {
          "type": "plain_text",
          "text": "Goal to Add to",
          "emoji": true
        },
        "element": {
          "type": "static_select",
          "placeholder": {
            "type": "plain_text",
            "text": "Choose Goal",
            "emoji": true
          },
          "action_id": "add_goal_category",
          "options": [
            {
              "text": {
                "type": "plain_text",
                "text": "Customer Questions?",
                "emoji": true
              },
              "value": "Customer Questions?"
            },
            {
              "text": {
                "type": "plain_text",
                "text": "Team Needs",
                "emoji": true
              },
              "value": "Team Needs"
            },
            {
              "text": {
                "type": "plain_text",
                "text": "Learnings",
                "emoji": true
              },
              "value": "Learnings"
            },
            {
              "text": {
                "type": "plain_text",
                "text": "Questions?",
                "emoji": true
              },
              "value": "Questions?"
            },
            {
              "text": {
                "type": "plain_text",
                "text": "Other",
                "emoji": true
              },
              "value": "Other"
            }
          ]
        }
      },
      {
        "type": "input",
        "block_id": "add_goal_input_block-7",
        "label": {
          "type": "plain_text",
          "text": "Goal to Add",
          "emoji": true
        },
        "element": {
          "type": "plain_text_input",
          "action_id": "add_goal",
          "placeholder": {
            "type": "plain_text",
            "text": "Goal"
          }
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
}`