package modal

import (
	"errors"
	"fmt"
	"github.com/slack-go/slack"
	"log"
	"strings"
)

func AddNewEventToDay(blocks []slack.Block, eventDay string, newEvent *slack.SectionBlock) ([]*slack.SectionBlock, []*slack.SectionBlock) {
	todaysSectionBlocks, tomorrowsSectionBlocks := ExtractEvents(blocks)
	log.Printf("Extract got: %v, got1: %v\n", len(todaysSectionBlocks), len(tomorrowsSectionBlocks))
	todaysSectionBlocks, tomorrowsSectionBlocks = addNewEventToCorrectDay(eventDay, todaysSectionBlocks, newEvent, tomorrowsSectionBlocks)
	log.Printf("AddNew got: %v, got1: %v\n", len(todaysSectionBlocks), len(tomorrowsSectionBlocks))
	todaysSectionBlocks, tomorrowsSectionBlocks = ReplaceEmptyEventsWithNoEventsYet(todaysSectionBlocks, tomorrowsSectionBlocks)
	log.Printf("Replace got: %v, got1: %v\n", len(todaysSectionBlocks), len(tomorrowsSectionBlocks))
	return todaysSectionBlocks, tomorrowsSectionBlocks
}
func RemoveEvent(blocks []slack.Block, actionId string) ([]*slack.SectionBlock, []*slack.SectionBlock) {
	log.Printf("remove action id %s\n", actionId)
	todaysSectionBlocks, tomorrowsSectionBlocks := ExtractEvents(blocks)
	eventDay := extractDayFromActionId(actionId)
	index, err := extractIndex(actionId, func() []*slack.SectionBlock {
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
func extractDayFromActionId(actionId string) string {
	return strings.Split(actionId, "_")[1]
}

func extractIndex(actionId string, events []*slack.SectionBlock) (int, error) {
	for index, event := range events {
		if event.Accessory.ButtonElement.Value == actionId {
			return index, nil
		}
	}
	return -1, errors.New("couldn't find matching event")
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

func addNewEventToCorrectDay(eventDay string, todaysSectionBlocks []*slack.SectionBlock, newEvent *slack.SectionBlock, tomorrowsSectionBlocks []*slack.SectionBlock) ([]*slack.SectionBlock, []*slack.SectionBlock) {
	if eventDay == TodayOptionValue {
		todaysSectionBlocks = append(todaysSectionBlocks, newEvent)
	} else {
		tomorrowsSectionBlocks = append(tomorrowsSectionBlocks, newEvent)
	}
	return todaysSectionBlocks, tomorrowsSectionBlocks
}

func removeEventFromCorrectDay(eventDay string, todaysSectionBlocks []*slack.SectionBlock, indexToRemove int, tomorrowsSectionBlocks []*slack.SectionBlock) ([]*slack.SectionBlock, []*slack.SectionBlock) {
	if eventDay == TodayOptionValue {
		todaysSectionBlocks = remove(todaysSectionBlocks, indexToRemove)
	} else {
		tomorrowsSectionBlocks = remove(tomorrowsSectionBlocks, indexToRemove)
	}
	return todaysSectionBlocks, tomorrowsSectionBlocks
}

func remove(sectionBlocks []*slack.SectionBlock, indexToRemove int, ) []*slack.SectionBlock {
	copy(sectionBlocks[indexToRemove:], sectionBlocks[indexToRemove+1:])
	sectionBlocks[len(sectionBlocks)-1] = &slack.SectionBlock{}
	sectionBlocks = sectionBlocks[:len(sectionBlocks)-1]
	return sectionBlocks
	//return make([]*slack.SectionBlock, 0)
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

func Filter(vs []slack.Block, f func(slack.Block) bool) []slack.Block {
	vsf := make([]slack.Block, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Index(vs []slack.Block, t slack.MessageBlockType) int {
	for i, v := range vs {
		if v.BlockType() == t {
			return i
		}
	}
	return -1
}
