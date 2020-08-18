package modal

import (
	"github.com/slack-go/slack"
)

func NewSummaryModal(todayEvents []*slack.SectionBlock, tomorrowEvents []*slack.SectionBlock) slack.ModalViewRequest {
	return slack.ModalViewRequest{
		Type:   slack.VTModal,
		Title:  slack.NewTextBlockObject(slack.PlainTextType, "bZapp", true, false),
		Close:  slack.NewTextBlockObject(slack.PlainTextType, "Cancel", true, false),
		Submit: slack.NewTextBlockObject(slack.PlainTextType, "Submit", true, false),
		Blocks: slack.Blocks{
			BlockSet: buildEventBlocks(todayEvents, tomorrowEvents),
		},
	}
}

func buildEventBlocks(todayEvents []*slack.SectionBlock, tomorrowEvents []*slack.SectionBlock) []slack.Block {
	blocks := BuildEventsBlock(todayEvents, tomorrowEvents)

	blocks = append(blocks,
		slack.NewDividerBlock(),
		slack.NewActionBlock(
			"actions_block",
			slack.NewButtonBlockElement("edit_events", "edit_events", slack.NewTextBlockObject(slack.PlainTextType, "Edit Events", true, false)),
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
	)

	return blocks
}

func BuildEventsBlock(todayEvents []*slack.SectionBlock, tomorrowEvents []*slack.SectionBlock) []slack.Block {
	blocks := []slack.Block{
		slack.NewDividerBlock(),
		slack.NewContextBlock("", slack.NewTextBlockObject(slack.MarkdownType, "*Today's Events*", false, false)),
		slack.NewDividerBlock(),
	}

	for _, event := range todayEvents {
		blocks = append(blocks, event)
	}

	blocks = append(blocks, slack.NewDividerBlock(),
		slack.NewContextBlock("", slack.NewTextBlockObject(slack.MarkdownType, "*Tomorrow's Events*", false, false)),
		slack.NewDividerBlock(),
	)

	for _, event := range tomorrowEvents {
		blocks = append(blocks, event)
	}
	return blocks
}
