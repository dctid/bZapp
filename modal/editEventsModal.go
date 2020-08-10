package modal

import (
	"github.com/slack-go/slack"
)

const AddEventTitleInputBlock = "add_event_title_input_block"
const AddEventDayInputBlock = "add_event_day_input_block"
const AddEventHoursInputBlock = "add_event_hours_input_block"
const AddEventMinsInputBlock = "add_event_mins_input_block"

const AddEventTitleActionId = "add_event_title"
const AddEventDayActionId = "add_event_day"
const AddEventHoursActionId = "add_event_hour"
const AddEventMinsActionId = "add_event_mins"

const TodayOptionValue = "today"

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
		slack.NewInputBlock(AddEventTitleInputBlock, slack.NewTextBlockObject("plain_text", "Add Event", false, false),
			slack.NewPlainTextInputBlockElement(slack.NewTextBlockObject("plain_text", "Title", false, false), AddEventTitleActionId),
		),
		slack.NewInputBlock(AddEventDayInputBlock, slack.NewTextBlockObject("plain_text", "Day", true, false),
			slack.NewRadioButtonsBlockElement(AddEventDayActionId,
				slack.NewOptionBlockObject(TodayOptionValue, slack.NewTextBlockObject("plain_text", "Today", true, false)),
				slack.NewOptionBlockObject("tomorrow", slack.NewTextBlockObject("plain_text", "Tomorrow", true, false))),
		),
		slack.NewInputBlock(AddEventHoursInputBlock, slack.NewTextBlockObject("plain_text", "Hour", true, false),
			slack.NewOptionsSelectBlockElement("static_select", slack.NewTextBlockObject("plain_text", "Select hour", true, false), AddEventHoursActionId, hourOptions...),
		),
		slack.NewInputBlock(AddEventMinsInputBlock, slack.NewTextBlockObject("plain_text", "Minutes", true, false),
			slack.NewOptionsSelectBlockElement("static_select", slack.NewTextBlockObject("plain_text", "Select Minutes", true, false), AddEventMinsActionId, minOptions...),
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
