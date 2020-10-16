package view

import (
	"encoding/json"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/model"

	"github.com/stretchr/testify/assert"
	"testing"
)

var summaryModal = `{
  "title": {
    "type": "plain_text",
    "text": "bZapp",
    "emoji": true
  },
  "private_metadata": "{\"Index\":0,\"Events\":{\"TodaysEvents\":[{\"Id\":\"\",\"Title\":\"Standup\",\"Day\":\"today\",\"Hour\":9,\"Min\":15,\"AmPm\":\"AM\"},{\"Id\":\"\",\"Title\":\"IPM\",\"Day\":\"today\",\"Hour\":11,\"Min\":30,\"AmPm\":\"AM\"},{\"Id\":\"\",\"Title\":\"Retro\",\"Day\":\"today\",\"Hour\":3,\"Min\":15,\"AmPm\":\"PM\"}],\"TomorrowsEvents\":[{\"Id\":\"\",\"Title\":\"Standup\",\"Day\":\"tomorrow\",\"Hour\":9,\"Min\":15,\"AmPm\":\"AM\"},{\"Id\":\"\",\"Title\":\"UserInterview\",\"Day\":\"tomorrow\",\"Hour\":1,\"Min\":30,\"AmPm\":\"PM\"},{\"Id\":\"\",\"Title\":\"Synthesis\",\"Day\":\"tomorrow\",\"Hour\":3,\"Min\":0,\"AmPm\":\"PM\"}]},\"Goals\":null}",
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
        "text": ":small_orange_diamond: 9:15 Standup"
      }
    },
    {
      "type": "section",
      "text": {
        "type": "mrkdwn",
        "text": ":small_orange_diamond: 11:30 IPM"
      }
    },
    {
      "type": "section",
      "text": {
        "type": "mrkdwn",
        "text": ":small_orange_diamond: 3:15 Retro"
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
        "text": ":small_orange_diamond: 9:15 Standup"
      }
    },
    {
      "type": "section",
      "text": {
        "type": "mrkdwn",
        "text": ":small_orange_diamond: 1:30 User Interview"
      }
    },
    {
      "type": "section",
      "text": {
        "type": "mrkdwn",
        "text": ":small_orange_diamond: 3:00 Synthesis"
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
}`

func TestNewSummaryModal(t *testing.T) {

	testModel := &model.Model{
		Events: model.Events{
			TodaysEvents: []model.Event{
				{
					Title: "Standup",
					Day:   TodayOptionValue,
					Hour:  9,
					Min:   15,
					AmPm:  "AM",
				},
				{
					Title: "IPM",
					Day:   TodayOptionValue,
					Hour:  11,
					Min:   30,
					AmPm:  "AM",
				},
				{
					Title: "Retro",
					Day:   TodayOptionValue,
					Hour:  3,
					Min:   15,
					AmPm:  "PM",
				},
			},
			TomorrowsEvents: []model.Event{
				{
					Title: "Standup",
					Day:   TomorrowOptionValue,
					Hour:  9,
					Min:   15,
					AmPm:  "AM",
				},
				{
					Title: "User Interview",
					Day:   TomorrowOptionValue,
					Hour:  1,
					Min:   30,
					AmPm:  "PM",
				},
				{
					Title: "Synthesis",
					Day:   TomorrowOptionValue,
					Hour:  3,
					Min:   0,
					AmPm:  "PM",
				},
			},
		},
		Goals: nil,
	}

	result := NewSummaryModal(testModel)
	actualJson, _ := json.Marshal(result)
	expectedJsonString, _ := format.PrettyJson(summaryModal)
	actualJsonString, _ := format.PrettyJson(string(actualJson))
	assert.EqualValues(t, expectedJsonString, actualJsonString)
}
