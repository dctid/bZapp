package test

import "github.com/slack-go/slack"

var SubmitPayload = slack.InteractionCallback{
	Type:      "view_submission",
	View: slack.View{
		Title: &slack.TextBlockObject{Text: "bZapp"},
		PrivateMetadata: `{"channel_id":"D7P4LC5G9","response_url":"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"}`,
	},
}

