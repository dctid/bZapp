package modal

import (
	"errors"
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func convertToEventsWithoutRemoveButton(todaysSectionBlocks []*slack.SectionBlock, tomorrowsSectionBlocks []*slack.SectionBlock) ([]*slack.SectionBlock, []*slack.SectionBlock) {
	return removeAccessory(todaysSectionBlocks), removeAccessory(tomorrowsSectionBlocks)
}

func removeAccessory(sectionBlocks []*slack.SectionBlock) []*slack.SectionBlock {
	result := make([]*slack.SectionBlock, len(sectionBlocks))
	for index, sectionBlock := range sectionBlocks {
		result[index] = slack.NewSectionBlock(
			sectionBlock.Text,
			sectionBlock.Fields,
			nil,
		)
	}
	return result
}

func extractDay(actionValue string) string {
	return strings.Split(actionValue, "_")[1]
}

func extractIndex(actionId string, events []*slack.SectionBlock) (int, error) {
	for index, event := range events {
		if event.Accessory.ButtonElement.Value == actionId {
			return index, nil
		}
	}
	return -1, errors.New("couldn't find matching event")
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
}

func convertToEventsWithRemoveButton(day string, sectionBlocks []*slack.SectionBlock) []*slack.SectionBlock {
	convertedBlocks := make([]*slack.SectionBlock, len(sectionBlocks))

	for index, sectionBlock := range sectionBlocks {
		convertedBlocks[index] = slack.NewSectionBlock(
			sectionBlock.Text,
			sectionBlock.Fields,
			slack.NewAccessory(
				slack.NewButtonBlockElement(
					RemoveEventActionId,
					fmt.Sprintf("remove_%s_%d", day, index),
					slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false),
				),
			),
		)
	}
	return convertedBlocks
}

func convertToEventsWithRemoveButton2(events []model.Event) []*slack.SectionBlock {
	convertedBlocks := make([]*slack.SectionBlock, len(events))

	for index, event := range events {
		convertedBlocks[index] = slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, event.ToString(), false, false),
			nil,
			slack.NewAccessory(
				slack.NewButtonBlockElement(
					RemoveEventActionId,
					fmt.Sprintf("remove_%s_%d", event.Day, index),
					slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false),
				),
			),
		)
	}
	return convertedBlocks
}

func convertToEventsWithoutRemoveButton2(events []model.Event) []*slack.SectionBlock {
	convertedBlocks := make([]*slack.SectionBlock, len(events))

	for index, event := range events {
		convertedBlocks[index] = slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, event.ToString(), false, false),
			nil,
			nil,
		)
	}
	return convertedBlocks
}



func filterBlocks(vs []slack.Block, f func(slack.Block) bool) []slack.Block {
	vsf := make([]slack.Block, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func firstBlockOfTypeIndex(vs []slack.Block, t slack.MessageBlockType) int {
	for i, v := range vs {
		if v.BlockType() == t {
			return i
		}
	}
	return -1
}

func eventSectionWithoutRemoveButton(title string, hour string, mins string) *slack.SectionBlock {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("%s:%s %s", strings.Fields(hour)[0], mins, title), false, false),
		nil,
		nil,
	)
}


func minOption(num int) *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(fmt.Sprintf("min-%d", num), slack.NewTextBlockObject(slack.PlainTextType, fmt.Sprintf(func() string {
		if num < 10 {
			return "0%d"
		} else {
			return "%d"
		}
	}(), num), true, false))
}

func hourOption(num int) *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(fmt.Sprintf("hour-%d", num), slack.NewTextBlockObject(slack.PlainTextType, fmt.Sprintf("%d %s", num, func() string {
		return amOrPm(num)
	}()), true, false))
}

func amOrPm(num int) string {
	if num < 9 || num == 12 {
		return "PM"
	} else {
		return "AM"
	}
}

func mapOptions(vs []int, f func(int) *slack.OptionBlockObject) []*slack.OptionBlockObject {
	vsm := make([]*slack.OptionBlockObject, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func sectionBlockFilter(block slack.Block) bool {
	return block.BlockType() == slack.MBTSection && block.(*slack.SectionBlock).Text.Text != NoEventsText
}


func mapToEvents(day string, blocks []slack.Block) []model.Event {
	var events = make([]model.Event, len(blocks))
	for index, block := range blocks {
		events[index] = convertToEvent(day, block)
	}
	return events
}

func convertToEvent(day string, block slack.Block) model.Event {
	spacesOrColon := regexp.MustCompile(`(?:\:|\s)+`)
	text := block.(*slack.SectionBlock).Text.Text
	tokens := spacesOrColon.Split(text, -1)
	log.Printf("text %s, %d", tokens, len(tokens))

	hour, _ := strconv.Atoi(tokens[0])
	mins, _ := strconv.Atoi(tokens[1])
	amPm := amOrPm(hour)

	return model.Event{Title: tokens[2], Day: day, Hour: hour, Min: mins, AmPm: amPm}
}