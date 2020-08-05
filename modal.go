package bZapp

import (
	"fmt"
	"github.com/slack-go/slack"
)

func NewModal(todayEvents *slack.SectionBlock, tomorrowEvents *slack.SectionBlock) slack.ModalViewRequest {
	titleText := slack.NewTextBlockObject("plain_text", "bZapp", true, false)
	submitText := slack.NewTextBlockObject("plain_text", "Submit", true, false)
	closeText := slack.NewTextBlockObject("plain_text", "Cancel", true, false)
	todayHeader := slack.NewContextBlock("", slack.NewTextBlockObject("mrkdwn", "*Today's Events*", false, false))
	tomorrowHeader := slack.NewContextBlock("", slack.NewTextBlockObject("mrkdwn", "*Tomorrow's Events*", false, false))
	addEventsElement := slack.NewPlainTextInputBlockElement(slack.NewTextBlockObject("plain_text", "Title", false, false), "add_event")
	addEvents := slack.NewInputBlock("", slack.NewTextBlockObject("plain_text", "Add Event", false, false), addEventsElement)

	hours := []int{9, 10, 11, 12, 1, 2, 3, 4}
	hourOptions := Map(hours, HourOption)

	mins := []int{0, 15, 30, 45}
	minOptions := Map(mins, MinOption)

	datepicker := slack.DatePickerBlockElement{
		Type:        "datepicker",
		ActionID:    "datepicker",
		Placeholder: slack.NewTextBlockObject("plain_text", "Select a date", true, false),
	}

	actions := slack.NewActionBlock(
		"",
		slack.NewOptionsSelectBlockElement("static_select", slack.NewTextBlockObject("plain_text", "Select hour", true, false), "hours_select", hourOptions...),
		slack.NewOptionsSelectBlockElement("static_select", slack.NewTextBlockObject("plain_text", "Select minutes", true, false), "mins_select", minOptions...),
		datepicker,
		slack.NewButtonBlockElement("", "add_event", slack.NewTextBlockObject("plain_text", "Add", true, false)),
	)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewDividerBlock(),
			todayHeader,
			slack.NewDividerBlock(),
			todayEvents,
			slack.NewDividerBlock(),
			tomorrowHeader,
			slack.NewDividerBlock(),
			tomorrowEvents,
			slack.NewDividerBlock(),
			addEvents,
			actions,
		},
	}

	var modalRequest slack.ModalViewRequest
	modalRequest.Type = "modal"
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks

	return modalRequest
}

var NoEventYetSection = slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "No events yet", false, false), nil, nil)

func MinOption(num int) *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(fmt.Sprintf("min-%d", num), slack.NewTextBlockObject("plain_text", fmt.Sprintf(func() string {
		if num < 10 {
			return "0%d"
		} else {
			return "%d"
		}
	}(), num), true, false))
}

func HourOption(num int) *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(fmt.Sprintf("hour-%d", num), slack.NewTextBlockObject("plain_text", fmt.Sprintf("%d %s", num, func() string {
		if num < 9 || num == 12 {
			return "PM"
		} else {
			return "AM"
		}
	}()), true, false))
}

func Map(vs []int, f func(int) *slack.OptionBlockObject) []*slack.OptionBlockObject {
	vsm := make([]*slack.OptionBlockObject, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func TodaySection() *slack.SectionBlock {
	return slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "9:15 Standup", false, false),
		nil,
		slack.NewAccessory(slack.NewButtonBlockElement("", "remove_today_1", slack.NewTextBlockObject("plain_text", "Remove", true, false))))
}
