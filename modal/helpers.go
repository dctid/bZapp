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


func convertToEventsWithRemoveButton(events []model.Event) []*slack.SectionBlock {
	convertedBlocks := make([]*slack.SectionBlock, len(events))

	for index, event := range events {
		convertedBlocks[index] = slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, event.ToString(), false, false),
			nil,
			slack.NewAccessory(
				slack.NewButtonBlockElement(
					RemoveEventActionId,
					fmt.Sprintf("remove_%s_%s", event.Day, event.Id),
					slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false),
				),
			),
			slack.SectionBlockOptionBlockID(event.Id),
		)
	}
	return convertedBlocks
}

func convertToEventsWithoutRemoveButton(events []model.Event) []*slack.SectionBlock {
	convertedBlocks := make([]*slack.SectionBlock, len(events))

	for index, event := range events {
		convertedBlocks[index] = slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, event.ToString(), false, false),
			nil,
			nil,
			slack.SectionBlockOptionBlockID(event.Id),
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
	sectionBlock := block.(*slack.SectionBlock)
	text := sectionBlock.Text.Text
	tokens := spacesOrColon.Split(text, -1)
	log.Printf("text %s, %d", tokens, len(tokens))

	hour, _ := strconv.Atoi(tokens[0])
	mins, _ := strconv.Atoi(tokens[1])
	amPm := amOrPm(hour)

	return model.Event{Id: sectionBlock.BlockID, Title: tokens[2], Day: day, Hour: hour, Min: mins, AmPm: amPm}
}