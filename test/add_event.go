package test

import "github.com/slack-go/slack"

var AddEventSubmissionPayload = slack.InteractionCallback{
	Type:      "view_submission",
	View: slack.View{
		Title: &slack.TextBlockObject{Text: "bZapp - Edit Events"},
		State: &slack.ViewState{Values: map[string]map[string]slack.BlockAction{
			"add_event_title_input_block-4": {"add_event_title": slack.BlockAction{Value: "sd"}},
			"add_event_day_input_block-4": {"add_event_day": slack.BlockAction{SelectedOption: slack.OptionBlockObject{Value: "today"}}},
			"add_event_hours_input_block-4": {"add_event_hour": slack.BlockAction{SelectedOption: slack.OptionBlockObject{Text: &slack.TextBlockObject{Text: "1 PM"}}}},
			"add_event_mins_input_block-4": {"add_event_mins": slack.BlockAction{SelectedOption: slack.OptionBlockObject{Text: &slack.TextBlockObject{Text: "15"}}}},
		}},
		PrivateMetadata: `{"channel_id":"add_event_channel_id","response_url":"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"}`,
	},
}

