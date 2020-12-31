package test

import "github.com/slack-go/slack"

var EditEventsActionButtonPayload = slack.InteractionCallback{
	Type:      "block_actions",
	TriggerID: "Trigger",
	ActionCallback: slack.ActionCallbacks{
		BlockActions: []*slack.BlockAction{{
			ActionID: "edit_events",
		}},
	},
	View: slack.View{
		PrivateMetadata: `{"channel_id":"open_edit_events_channel_id","response_url":"https://hooks.slack.com/commands/T7NS02BFB/1307783467168/Gvz9lFVBwn9xo8TweP2vJHsP"}`,
	},
}

