package modal

import (
	"fmt"
	"github.com/slack-go/slack"
)

func NewEditEventsModal(index int, todayEvents []*slack.SectionBlock, tomorrowEvents []*slack.SectionBlock) slack.ModalViewRequest {
	return slack.ModalViewRequest{
		Type:   slack.VTModal,
		Title:  slack.NewTextBlockObject(slack.PlainTextType, "bZapp - Edit Events", true, false),
		Close:  slack.NewTextBlockObject(slack.PlainTextType, "Back", true, false),
		Submit: slack.NewTextBlockObject(slack.PlainTextType, "Add", true, false),
		Blocks: slack.Blocks{
			BlockSet: buildSummaryEventBlocks(index, todayEvents, tomorrowEvents),
		},
		NotifyOnClose: true,
	}
}

func buildSummaryEventBlocks(index int, todayEvents []*slack.SectionBlock, tomorrowEvents []*slack.SectionBlock) []slack.Block {
	hours := []int{9, 10, 11, 12, 1, 2, 3, 4}
	hourOptions := mapOptions(hours, hourOption)

	mins := []int{0, 15, 30, 45}
	minOptions := mapOptions(mins, minOption)

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
		slack.NewDividerBlock())

	for _, event := range tomorrowEvents {
		blocks = append(blocks, event)
	}

	blocks = append(blocks,

		slack.NewDividerBlock(),
		slack.NewInputBlock(fmt.Sprintf("%s-%d", AddEventTitleInputBlock, index), slack.NewTextBlockObject(slack.PlainTextType, "Add Event", false, false),
			slack.NewPlainTextInputBlockElement(slack.NewTextBlockObject(slack.PlainTextType, "Title", false, false), AddEventTitleActionId),
		),
		slack.NewInputBlock(fmt.Sprintf("%s-%d", AddEventDayInputBlock, index), slack.NewTextBlockObject(slack.PlainTextType, "Day", true, false),
			slack.NewRadioButtonsBlockElement(AddEventDayActionId,
				slack.NewOptionBlockObject(TodayOptionValue, slack.NewTextBlockObject(slack.PlainTextType, "Today", true, false)),
				slack.NewOptionBlockObject(TomorrowOptionValue, slack.NewTextBlockObject(slack.PlainTextType, "Tomorrow", true, false))),
		),
		slack.NewInputBlock(fmt.Sprintf("%s-%d", AddEventHoursInputBlock, index), slack.NewTextBlockObject(slack.PlainTextType, "Hour", true, false),
			slack.NewOptionsSelectBlockElement("static_select", slack.NewTextBlockObject(slack.PlainTextType, "Select hour", true, false), AddEventHoursActionId, hourOptions...),
		),
		slack.NewInputBlock(fmt.Sprintf("%s-%d", AddEventMinsInputBlock, index), slack.NewTextBlockObject(slack.PlainTextType, "Minutes", true, false),
			slack.NewOptionsSelectBlockElement("static_select", slack.NewTextBlockObject(slack.PlainTextType, "Select Minutes", true, false), AddEventMinsActionId, minOptions...),
		))
	return blocks
}
