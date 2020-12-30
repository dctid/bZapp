package test

import "github.com/slack-go/slack"

var RemoveGoalActionPayload = slack.InteractionCallback{
	Type:      "block_actions",
	ActionCallback: slack.ActionCallbacks{
		BlockActions: []*slack.BlockAction{{
			ActionID: "remove_goal",
			Value: "remove_Questions?_GqrOzx",
			BlockID: "GqrOzx",
		}},

	},
	View: slack.View{
		ID: "V01DBFTR588",
		ExternalID: "remove_goal_outsideId",
		PrivateMetadata: `{"channel_id":"D7P4LC5G9","response_url":"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"}`,
	},
	ViewSubmissionCallback: slack.ViewSubmissionCallback{Hash: "remove_goal_hash"},
}

