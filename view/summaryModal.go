package view

import (
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
)

func NewSummaryModal(updatedModel *model.Model) slack.ModalViewRequest {
	return slack.ModalViewRequest{
		Type:   slack.VTModal,
		Title:  slack.NewTextBlockObject(slack.PlainTextType, "bZapp", true, false),
		Close:  slack.NewTextBlockObject(slack.PlainTextType, "Cancel", true, false),
		Submit: slack.NewTextBlockObject(slack.PlainTextType, "Submit", true, false),
		Blocks: slack.Blocks{
			BlockSet: buildEventBlocks(false, updatedModel),
		},
		PrivateMetadata: updatedModel.ConvertModelToJson(),
	}
}

func buildEventBlocks(editable bool, updatedModel *model.Model) []slack.Block {

	blocks := buildEventsBlock(false, updatedModel.Events)
	blocks = append(blocks, buildGoalsBlock(editable, updatedModel.Goals)...)

	blocks = append(blocks, actionBlock()...)

	return blocks
}


func actionBlock() []slack.Block {
	return []slack.Block{
		slack.NewDividerBlock(),
		slack.NewActionBlock(
			"actions_block",
			slack.NewButtonBlockElement("edit_events", "edit_events", slack.NewTextBlockObject(slack.PlainTextType, "Edit Events", true, false)),
			slack.NewButtonBlockElement("edit_goals", "edit_goals", slack.NewTextBlockObject(slack.PlainTextType, "Edit Goals", true, false)),
		),
	}
}
