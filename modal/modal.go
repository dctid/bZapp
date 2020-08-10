package modal

import (
	"fmt"
	"github.com/slack-go/slack"
)

func BuildNewEventSectionBlock(values map[string]map[string]slack.BlockAction) (string, *slack.SectionBlock) {
	eventTitle := values[AddEventTitleInputBlock][AddEventTitleActionId].Value
	eventDay := values[AddEventDayInputBlock][AddEventDayActionId].SelectedOption.Value
	eventHours := values[AddEventHoursInputBlock][AddEventHoursActionId].SelectedOption.Text.Text
	eventMins := values[AddEventMinsInputBlock][AddEventMinsActionId].SelectedOption.Text.Text
	fmt.Printf("Add Event title: %s, day: %s, hour: %s, mins: %s\n", eventTitle, eventDay, eventHours, eventMins)

	newEvent := EventSection(0, eventTitle, eventHours, eventMins)
	return eventDay, newEvent
}

