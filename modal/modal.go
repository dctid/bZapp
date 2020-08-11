package modal

import (
	"fmt"
	"github.com/slack-go/slack"
	"log"
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
const TomorrowOptionValue = "tomorrow"

const EditEventsActionId = "edit_events"
const RemoveEventActionId = "remove_event"

const NoEventsText = "_No events yet_"

var NoEventYetSection = []*slack.SectionBlock{slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, NoEventsText, false, false), nil, nil)}

func BuildNewEventSectionBlock(values map[string]map[string]slack.BlockAction) (string, *slack.SectionBlock) {
	eventTitle := values[AddEventTitleInputBlock][AddEventTitleActionId].Value
	eventDay := values[AddEventDayInputBlock][AddEventDayActionId].SelectedOption.Value
	eventHours := values[AddEventHoursInputBlock][AddEventHoursActionId].SelectedOption.Text.Text
	eventMins := values[AddEventMinsInputBlock][AddEventMinsActionId].SelectedOption.Text.Text
	fmt.Printf("Add Event title: %s, day: %s, hour: %s, mins: %s\n", eventTitle, eventDay, eventHours, eventMins)

	newEvent := eventSectionWithoutRemoveButton(eventTitle, eventHours, eventMins)
	return eventDay, newEvent
}

func ExtractEvents(blocks []slack.Block) ([]*slack.SectionBlock, []*slack.SectionBlock) {
	firstContextBlock := Index(blocks, slack.MBTContext)
	secondContextBlock := Index(blocks[firstContextBlock+1:], slack.MBTContext)

	sectionBlockFilter := func(block slack.Block) bool {
		return block.BlockType() == slack.MBTSection && block.(*slack.SectionBlock).Text.Text != NoEventsText
	}
	todaysBlocks := Filter(blocks[firstContextBlock:secondContextBlock+firstContextBlock], sectionBlockFilter)
	fmt.Printf("Today filtered got: %v\n", len(todaysBlocks))
	todaysSectionBlocks := make([]*slack.SectionBlock, len(todaysBlocks))
	for i, todayBlock := range todaysBlocks {
		todaysSectionBlocks[i] = todayBlock.(*slack.SectionBlock)
	}

	tomorrowsBlocks := Filter(blocks[secondContextBlock+firstContextBlock:], sectionBlockFilter)
	fmt.Printf("tomorrow filtered got: %v\n", len(todaysBlocks))
	tomorrowsSectionBlocks := make([]*slack.SectionBlock, len(tomorrowsBlocks))
	for i, tomorrowBlock := range tomorrowsBlocks {
		sectionBlock := tomorrowBlock.(*slack.SectionBlock)
		tomorrowsSectionBlocks[i] = sectionBlock
	}
	return todaysSectionBlocks, tomorrowsSectionBlocks
}

func ConvertToEventsWithRemoveButton(todaysSectionBlocks []*slack.SectionBlock, tomorrowsSectionBlocks []*slack.SectionBlock) ([]*slack.SectionBlock, []*slack.SectionBlock) {
	return convertToEventsWithRemoveButton(TodayOptionValue, todaysSectionBlocks),
		convertToEventsWithRemoveButton(TomorrowOptionValue, tomorrowsSectionBlocks)
}

func AddNewEventToDay(blocks []slack.Block, eventDay string, newEvent *slack.SectionBlock) ([]*slack.SectionBlock, []*slack.SectionBlock) {
	todaysSectionBlocks, tomorrowsSectionBlocks := ExtractEvents(blocks)
	log.Printf("Extract got: %v, got1: %v\n", len(todaysSectionBlocks), len(tomorrowsSectionBlocks))
	todaysSectionBlocks, tomorrowsSectionBlocks = addNewEventToCorrectDay(eventDay, todaysSectionBlocks, newEvent, tomorrowsSectionBlocks)
	todaysSectionBlocks, tomorrowsSectionBlocks = convertToEventsWithoutRemoveButton(todaysSectionBlocks, tomorrowsSectionBlocks)
	log.Printf("AddNew got: %v, got1: %v\n", len(todaysSectionBlocks), len(tomorrowsSectionBlocks))
	todaysSectionBlocks, tomorrowsSectionBlocks = ReplaceEmptyEventsWithNoEventsYet(todaysSectionBlocks, tomorrowsSectionBlocks)
	log.Printf("Replace got: %v, got1: %v\n", len(todaysSectionBlocks), len(tomorrowsSectionBlocks))
	return todaysSectionBlocks, tomorrowsSectionBlocks
}


func RemoveEvent(blocks []slack.Block, actionValue string) ([]*slack.SectionBlock, []*slack.SectionBlock) {
	log.Printf("remove action id %s\n", actionValue)
	todaysSectionBlocks, tomorrowsSectionBlocks := ExtractEvents(blocks)
	eventDay := extractDay(actionValue)
	index, err := extractIndex(actionValue, func() []*slack.SectionBlock {
		if eventDay == TodayOptionValue {
			return todaysSectionBlocks
		} else {
			return tomorrowsSectionBlocks
		}
	}())
	if err == nil {
		log.Printf("day: %s, index: %d\n", eventDay, index)
		log.Printf("Extracted got: %v, got1: %v\n", len(todaysSectionBlocks), len(tomorrowsSectionBlocks))
		todaysSectionBlocks, tomorrowsSectionBlocks = removeEventFromCorrectDay(eventDay, todaysSectionBlocks, index, tomorrowsSectionBlocks)
		log.Printf("AddNew got: %v, got1: %v\n", len(todaysSectionBlocks), len(tomorrowsSectionBlocks))
	} else {
		log.Printf("err: %s\n", err)
	}
	todaysSectionBlocks, tomorrowsSectionBlocks = ReplaceEmptyEventsWithNoEventsYet(todaysSectionBlocks, tomorrowsSectionBlocks)
	log.Printf("Replace got: %v, got1: %v\n", len(todaysSectionBlocks), len(tomorrowsSectionBlocks))

	return todaysSectionBlocks, tomorrowsSectionBlocks
}
func ReplaceEmptyEventsWithNoEventsYet(todaysSectionBlocks []*slack.SectionBlock, tomorrowsSectionBlocks []*slack.SectionBlock) ([]*slack.SectionBlock, []*slack.SectionBlock) {
	if len(todaysSectionBlocks) == 0 {
		todaysSectionBlocks = NoEventYetSection
	}
	if len(tomorrowsSectionBlocks) == 0 {
		tomorrowsSectionBlocks = NoEventYetSection
	}
	return todaysSectionBlocks, tomorrowsSectionBlocks
}
