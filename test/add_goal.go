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
		PrivateMetadata: `{"channel_id":"add_goal_channel_id","response_url":"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"}`,
	},
}