package view

import (
	"encoding/json"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEditGoalsModal(t *testing.T) {

	type args struct {
		index int
		model model.Model
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				index: 1,
				model: model.Model{},
			},
			want: format.PrettyJsonNoError(expectModalJson),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEditGoalsModal(tt.args.index, tt.args.model); !assert.EqualValues(t, format.PrettyJsonNoError(marshalNoError(got)), tt.want) {
				t.Errorf("NewEditGoalsModal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func marshalNoError(thing interface{}) string {
	marshal, _ := json.Marshal(thing)
	return string(marshal)
}

func parse(jsonStr string) slack.ModalViewRequest {
	var result slack.ModalViewRequest
	json.Unmarshal([]byte(jsonStr), &result)
	return result
}

const expectModalJson = `{
	"title": {
		"type": "plain_text",
		"text": "bZapp - Edit Goals",
		"emoji": true
	},
	"notify_on_close": true,
	"private_metadata": "{\"Events\":{\"TodaysEvents\":null,\"TomorrowsEvents\":null},\"Goals\":null}",
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
					"text": "*Customer Questions?*"
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
				"text": "_No goals yet_"
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
				"text": "_No goals yet_"
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
					"text": "*Learnings*"
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
				"text": "_No goals yet_"
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
					"text": "*Questions?*"
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
				"text": "_No goals yet_"
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
					"text": "*Other*"
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
				"text": "_No goals yet_"
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "input",
			"block_id": "add_goal_category_input_block-1",
			"element": {
				"type": "static_select",
				"action_id": "add_goal_category",
				"placeholder": {
					"type": "plain_text",
					"text": "Choose Goal",
					"emoji": true
				},
				"options": [
					{
						"text": {
							"type": "plain_text",
							"text": "Customer Questions?",
							"emoji": true
						},
						"value": "goal-Customer Questions?"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "Team Needs",
							"emoji": true
						},
						"value": "goal-Team Needs"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "Learnings",
							"emoji": true
						},
						"value": "goal-Learnings"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "Questions?",
							"emoji": true
						},
						"value": "goal-Questions?"
					},
					{
						"text": {
							"type": "plain_text",
							"text": "Other",
							"emoji": true
						},
						"value": "goal-Other"
					}
				]
			},
			"label": {
				"type": "plain_text",
				"text": "Goal to Add to",
				"emoji": true
			}
		},
		{
			"type": "input",
			"block_id": "add_goal_input_block-1",
			"element": {
				"action_id": "add_goal",
				"type": "plain_text_input",
				"placeholder": {
					"text": "Goal",
					"type": "plain_text"
				}
			},
			"label": {
				"type": "plain_text",
				"text": "Goal to Add",
				"emoji": true
			}
		}
	]
}`
