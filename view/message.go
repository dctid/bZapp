package view

import (
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
)

func DailySummaryMessage(currentModel *model.Model) *slack.Message {
	eventBlocks := buildEventsBlock(false, currentModel.Events)
	eventBlocks = append(eventBlocks, buildGoalsBlock(currentModel.Goals)...)

	return &slack.Message{
		Msg: slack.Msg{
			Blocks: slack.Blocks{
				BlockSet: eventBlocks,
			},
			Text:         "bZapp - Today's Standup Summary",
			ResponseType: slack.ResponseTypeInChannel,
		},
	}
}
