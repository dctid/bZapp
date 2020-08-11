package modal

import (
	"fmt"
	"github.com/slack-go/slack"
	"strings"
)

const EditEventsActionId = "edit_events"

func BuildNewEventSectionBlock(values map[string]map[string]slack.BlockAction) (string, *slack.SectionBlock) {
	eventTitle := values[AddEventTitleInputBlock][AddEventTitleActionId].Value
	eventDay := values[AddEventDayInputBlock][AddEventDayActionId].SelectedOption.Value
	eventHours := values[AddEventHoursInputBlock][AddEventHoursActionId].SelectedOption.Text.Text
	eventMins := values[AddEventMinsInputBlock][AddEventMinsActionId].SelectedOption.Text.Text
	fmt.Printf("Add Event title: %s, day: %s, hour: %s, mins: %s\n", eventTitle, eventDay, eventHours, eventMins)

	newEvent := EventSectionWithoutRemoveButton(eventTitle, eventHours, eventMins)
	return eventDay, newEvent
}

func EventSectionWithRemoveButton(day string, index int, title string, hour string, mins string) *slack.SectionBlock {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("%s:%s %s", strings.Fields(hour)[0], mins, title), false, false),
		nil,
		slack.NewAccessory(
			slack.NewButtonBlockElement(
				RemoveEventActionId,
				fmt.Sprintf("remove_%s_%d", day, index),
				slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false),
			),
		),
	)
}


func EventSectionWithoutRemoveButton(title string, hour string, mins string) *slack.SectionBlock {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("%s:%s %s", strings.Fields(hour)[0], mins, title), false, false),
		nil,
		nil,
	)
}
