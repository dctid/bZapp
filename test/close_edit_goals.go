package test

import (
	"github.com/slack-go/slack"
)

var CloseEditGoalsPayload = slack.InteractionCallback{
	Type:      "view_closed",
	View: slack.View{
		ID: "V01DBFTR588",
		RootViewID: "V01DBEZ35GQ",
		ExternalID: "Close_goals_externalId",
		Title: &slack.TextBlockObject{Text: "bZapp - Edit Goals"},
		PrivateMetadata: `{"channel_id":"close_edit_goals_channel_id","response_url":"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"}`,
	},
	ViewSubmissionCallback: slack.ViewSubmissionCallback{Hash: "close_goals_hash"},
}