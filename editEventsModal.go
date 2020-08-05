package bZapp

import (
	"github.com/slack-go/slack"
)

func NewEditEventsModal(todayEvents []*slack.SectionBlock, tomorrowEvents []*slack.SectionBlock) slack.ModalViewRequest {
	hours := []int{9, 10, 11, 12, 1, 2, 3, 4}
	hourOptions := Map(hours, HourOption)

	mins := []int{0, 15, 30, 45}
	minOptions := Map(mins, MinOption)

	blocks := []slack.Block{
		slack.NewDividerBlock(),
		slack.NewContextBlock("", slack.NewTextBlockObject("mrkdwn", "*Today's Events*", false, false)),
		slack.NewDividerBlock(),
	}
	for _, event := range todayEvents {
		blocks = append(blocks, event)
	}

	blocks = append(blocks, slack.NewDividerBlock(),
		slack.NewContextBlock("", slack.NewTextBlockObject("mrkdwn", "*Tomorrow's Events*", false, false)),
		slack.NewDividerBlock())

	for _, event := range tomorrowEvents {
		blocks = append(blocks, event)
	}

	blocks = append(blocks,

		slack.NewDividerBlock(),
		slack.NewInputBlock("", slack.NewTextBlockObject("plain_text", "Add Event", false, false),
			slack.NewPlainTextInputBlockElement(slack.NewTextBlockObject("plain_text", "Title", false, false), "add_event"),
		),
		slack.NewInputBlock("", slack.NewTextBlockObject("plain_text", "Day", true, false),
			slack.NewRadioButtonsBlockElement("today_or_tomorrow",
				slack.NewOptionBlockObject("today", slack.NewTextBlockObject("plain_text", "Today", true, false)),
				slack.NewOptionBlockObject("tomorrow", slack.NewTextBlockObject("plain_text", "Tomorrow", true, false))),
		),
		slack.NewInputBlock("", slack.NewTextBlockObject("plain_text", "Hour", true, false),
			slack.NewOptionsSelectBlockElement("static_select", slack.NewTextBlockObject("plain_text", "Select hour", true, false), "hours_select", hourOptions...),
		),
		slack.NewInputBlock("", slack.NewTextBlockObject("plain_text", "Minutes", true, false),
			slack.NewOptionsSelectBlockElement("static_select", slack.NewTextBlockObject("plain_text", "Select Minutes", true, false), "mins_select", minOptions...),
		))

	return slack.ModalViewRequest{
		Type:   slack.VTModal,
		Title:  slack.NewTextBlockObject("plain_text", "bZapp - Edit Events", true, false),
		Close:  slack.NewTextBlockObject("plain_text", "Cancel", true, false),
		Submit: slack.NewTextBlockObject("plain_text", "Add", true, false),
		Blocks: slack.Blocks{
			BlockSet: blocks,
		},
	}
}
