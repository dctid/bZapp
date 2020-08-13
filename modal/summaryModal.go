package modal

import (
	"github.com/slack-go/slack"
)

func NewSummaryModal(todayEvents []*slack.SectionBlock, tomorrowEvents []*slack.SectionBlock) slack.ModalViewRequest {
	titleText := slack.NewTextBlockObject(slack.PlainTextType, "bZapp", true, false)
	submitText := slack.NewTextBlockObject(slack.PlainTextType, "Submit", true, false)
	closeText := slack.NewTextBlockObject(slack.PlainTextType, "Cancel", true, false)
	blocks := slack.Blocks{BlockSet: buildEventBlocks(todayEvents, tomorrowEvents)}

	var modalRequest slack.ModalViewRequest
	modalRequest.Type = slack.VTModal
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks

	return modalRequest
}

func buildEventBlocks(todayEvents []*slack.SectionBlock, tomorrowEvents []*slack.SectionBlock) []slack.Block {
	blocks := BuildEventsBlock(todayEvents, tomorrowEvents)

	actions := slack.NewActionBlock(
		"actions_block",
		slack.NewButtonBlockElement("edit_events", "edit_events", slack.NewTextBlockObject(slack.PlainTextType, "Edit Events", true, false)),
	)

	conversationSelect := slack.InputBlock{
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
	}

	slack.NewConversationsSelect("convo_select", "convo_label")

	blocks = append(blocks,
		slack.NewDividerBlock(),
		actions,
		conversationSelect,
	)

	return blocks
}

func BuildEventsBlock(todayEvents []*slack.SectionBlock, tomorrowEvents []*slack.SectionBlock) []slack.Block {
	todayHeader := slack.NewContextBlock("", slack.NewTextBlockObject(slack.MarkdownType, "*Today's Events*", false, false))
	tomorrowHeader := slack.NewContextBlock("", slack.NewTextBlockObject(slack.MarkdownType, "*Tomorrow's Events*", false, false))

	blocks := []slack.Block{
		slack.NewDividerBlock(),
		todayHeader,
		slack.NewDividerBlock(),
	}

	for _, event := range todayEvents {
		blocks = append(blocks, event)
	}

	blocks = append(blocks, slack.NewDividerBlock(),
		tomorrowHeader,
		slack.NewDividerBlock(),
	)

	for _, event := range tomorrowEvents {
		blocks = append(blocks, event)
	}
	return blocks
}





