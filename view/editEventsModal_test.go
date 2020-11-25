package view

import (
	"encoding/json"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/model"
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
  "private_metadata": "{\"channel_id\":\"D7P4LC5G9\"}",
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
        "text": ":small_orange_diamond: 9:15 Standup"
      },
      "accessory": {
        "action_id": "remove_event",
        "type": "button",
        "text": {
          "type": "plain_text",
          "text": "Remove",
          "emoji": true
        },
        "value": "remove_today_FakeId1"
      },
      "block_id": "FakeId1"
    },
    {
      "type": "section",
      "text": {
        "type": "mrkdwn",
        "text": ":small_orange_diamond: 11:30 IPM"
      },
      "accessory": {
        "action_id": "remove_event",
        "type": "button",
        "text": {
          "type": "plain_text",
          "text": "Remove",
          "emoji": true
        },
        "value": "remove_today_FakeId2"
      },
      "block_id": "FakeId2"
    },
    {
      "type": "section",
      "block_id": "today_3",
      "text": {
        "type": "mrkdwn",
        "text": ":small_orange_diamond: 3:15 Retro"
      },
      "accessory": {
        "action_id": "remove_event",
        "type": "button",
        "text": {
          "type": "plain_text",
          "text": "Remove",
          "emoji": true
        },
        "value": "remove_today_FakeId3"
      },
      "block_id": "FakeId3"
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
        "text": ":small_orange_diamond: 9:15 Standup"
      },
      "accessory": {
        "action_id": "remove_event",
        "type": "button",
        "text": {
          "type": "plain_text",
          "text": "Remove",
          "emoji": true
        },
        "value": "remove_tomorrow_FakeId4"
      },
      "block_id": "FakeId4"
    },
    {
      "type": "section",
      "text": {
        "type": "mrkdwn",
        "text": ":small_orange_diamond: 1:30 User Interview"
      },
      "accessory": {
        "action_id": "remove_event",
        "type": "button",
        "text": {
          "type": "plain_text",
          "text": "Remove",
          "emoji": true
        },
        "value":"remove_tomorrow_FakeId5"
      },
      "block_id": "FakeId5"
    },
    {
      "type": "section",
      "text": {
        "type": "mrkdwn",
        "text": ":small_orange_diamond: 3:00 Synthesis"
      },
      "accessory": {
        "action_id": "remove_event",
        "type": "button",
        "text": {
          "type": "plain_text",
          "text": "Remove",
          "emoji": true
        },
        "value": "remove_tomorrow_FakeId6"
      },
      "block_id": "FakeId6"
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


	todaysEvents := []model.Event{
		{
			Id:    "FakeId1",
			Title: "Standup",
			Day:   TodayOptionValue,
			Hour:  9,
			Min:   15,
			AmPm:  "AM",
		},
		{
			Id:    "FakeId2",
			Title: "IPM",
			Day:   TodayOptionValue,
			Hour:  11,
			Min:   30,
			AmPm:  "AM",
		},
		{
			Id:    "FakeId3",
			Title: "Retro",
			Day:   TodayOptionValue,
			Hour:  3,
			Min:   15,
			AmPm:  "PM",
		},
	}

	tomorrowsEvents := []model.Event{
		{
			Id:    "FakeId4",
			Title: "Standup",
			Day:   TomorrowOptionValue,
			Hour:  9,
			Min:   15,
			AmPm:  "AM",
		},
		{
			Id:    "FakeId5",
			Title: "User Interview",
			Day:   TomorrowOptionValue,
			Hour:  1,
			Min:   30,
			AmPm:  "PM",
		},
		{
			Id:    "FakeId6",
			Title: "Synthesis",
			Day:   TomorrowOptionValue,
			Hour:  3,
			Min:   00,
			AmPm:  "PM",
		},
	}

	result := NewEditEventsModal( &model.Model{
		Index:  666,
		Events: model.Events{
			TodaysEvents: todaysEvents,
			TomorrowsEvents: tomorrowsEvents,
		},
		ChannelId: "D7P4LC5G9",
	})
	actualJson, _ := json.Marshal(result)
	expectedJsonString, _ := format.PrettyJson(editEventsModal)
	actualJsonString, _ := format.PrettyJson(string(actualJson))

	assert.EqualValues(t, expectedJsonString, actualJsonString)
}
