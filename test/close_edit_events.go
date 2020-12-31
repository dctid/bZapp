package test

import "github.com/slack-go/slack"

var CloseEditEventsPayload = slack.InteractionCallback{
	Type:      "view_closed",
	View: slack.View{
		ID: "V01CMKMUWUS",
		RootViewID: "V01DBEZ35GQ",
		ExternalID: "externalOd",
		Title: &slack.TextBlockObject{Text: "bZapp - Edit Events"},
		PrivateMetadata: `{"channel_id":"close_edit_events_channel_id","response_url":"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"}`,
	},
	ViewSubmissionCallback: slack.ViewSubmissionCallback{Hash: "close_events_hash"},
}

