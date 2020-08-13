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
	todayHeader := slack.NewContextBlock("", slack.NewTextBlockObject(slack.MarkdownType, "*Today's Events*", false, false))
	tomorrowHeader := slack.NewContextBlock("", slack.NewTextBlockObject(slack.MarkdownType, "*Tomorrow's Events*", false, false))

	actions := slack.NewActionBlock(
		"actions_block",
		slack.NewButtonBlockElement("edit_events", "edit_events", slack.NewTextBlockObject(slack.PlainTextType, "Edit Events", true, false)),
	)
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

	//{
	//	"block_id": "my_block_id",
	//	"type": "input",
	//	"optional": true,
	//	"label": {
	//	"type": "plain_text",
	//		"text": "Select a channel to post the result on"
	//},
	//	"element": {
	//	"action_id": "my_action_id",
	//		"type": "conversations_select",
	//		"response_url_enabled": true
	//}
	//},slack.OptTypeConversations

	conversationSelect := slack.InputBlock{
		Type:     "input",
		BlockID:  "convo_input_id",
		Label:   slack.NewTextBlockObject(slack.PlainTextType, "Select a channel to post the result on", false, false),
		Element:  slack.SelectBlockElement{
			Type:               slack.OptTypeConversations,
			ActionID:           "conversation_select_action_id",
			DefaultToCurrentConversation: true,
			ResponseURLEnabled: true,
		},
		Optional: false,
	}

	//conversationSelect := slack.NewInputBlock("conversation_select",
	//	slack.NewTextBlockObject(slack.PlainTextType, "Choose channel to post to", true, true),
	//	selectBlockElement,
	//	)

	slack.NewConversationsSelect("convo_select", "convo_label")

	blocks = append(blocks,
		slack.NewDividerBlock(),
		actions,
		conversationSelect,
	)

	return blocks
}





