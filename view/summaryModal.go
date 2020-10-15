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
			BlockSet: buildEventBlocks(updatedModel),
		},
		PrivateMetadata: updatedModel.ConvertModelToJson(),
	}
}

func buildEventBlocks(updatedModel *model.Model) []slack.Block {

	blocks := buildEventsBlock(false, updatedModel.Events)
	blocks = append(blocks, buildGoalsBlock(updatedModel.Goals)...)

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
		slack.InputBlock{
			Type:    "input",
			BlockID: "convo_input_id",
			Label:   slack.NewTextBlockObject(slack.PlainTextType, "Select a channel to post the result on", false, false),
			Element: slack.SelectBlockElement{
				Type:                         slack.OptTypeConversations,
				ActionID:                     "conversation_select_action_id",
				DefaultToCurrentConversation: true,
				ResponseURLEnabled:           true,
			},
			Optional: false,
		},
	}
}
