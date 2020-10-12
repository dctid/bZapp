package view

import (
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
	"log"
	"regexp"
	"strconv"
)

func convertToSectionBlocks(includeRemoveButton bool, events []model.Event) []slack.Block {

	numEvents := len(events)
	if numEvents == 0 {
		return NoEventYetSection
	}
	convertedBlocks := make([]slack.Block, numEvents)

	for index, event := range events {
		convertedBlocks[index] = slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, event.ToString(), false, false),
			nil,
			getRemoveButton(includeRemoveButton, event),
			slack.SectionBlockOptionBlockID(event.Id),
		)
	}
	return convertedBlocks
}

func convertGoalToSectionBlocks(includeRemoveButton bool, category string, goals []model.Goal) []slack.Block {

	numEvents := len(goals)
	if numEvents == 0 {
		return NoGoalsYetSection
	}
	convertedBlocks := make([]slack.Block, numEvents)

	for index, goal := range goals {
		convertedBlocks[index] = slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, goal.Value, false, false),
			nil,
			getGoalRemoveButton(includeRemoveButton, category, goal),
			slack.SectionBlockOptionBlockID(goal.Id),
		)
	}
	return convertedBlocks
}

func getRemoveButton(includeRemoveButton bool, event model.Event) *slack.Accessory {
	if includeRemoveButton {
		return slack.NewAccessory(
			slack.NewButtonBlockElement(
				RemoveEventActionId,
				fmt.Sprintf("remove_%s_%s", event.Day, event.Id),
				slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false),
			),
		)
	}
	return nil
}
func getGoalRemoveButton(includeRemoveButton bool, category string, goal model.Goal) *slack.Accessory {
	if includeRemoveButton {
		return slack.NewAccessory(
			slack.NewButtonBlockElement(
				RemoveGoalActionId,
				fmt.Sprintf("remove_%s_%s", category, goal.Id),
				slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false),
			),
		)
	}
	return nil
}

func groupSectionBlocks(blocks []slack.Block) map[string][]slack.Block {
	result := map[string][]slack.Block{}
	var key string
	for _, block := range blocks {
		if block.BlockType() == slack.MBTContext  {
			key = block.(*slack.ContextBlock).ContextElements.Elements[0].(*slack.TextBlockObject).Text
			result[key] = []slack.Block{}
		} else if  block.BlockType() == slack.MBTHeader {
			key = block.(*slack.HeaderBlock).Text.Text
			result[key] = []slack.Block{}
		} else if block.BlockType() == slack.MBTSection {
			if block.(*slack.SectionBlock).Text.Text != NoEventsText {
				blockMap := append(result[key], block)
				result[key] = blockMap
			}
		}
	}
	return result
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

func goalCategoryOption(goal string) *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(fmt.Sprintf("%s%s", GoalCategoryDropdownPrefix, goal), slack.NewTextBlockObject(slack.PlainTextType, fmt.Sprintf("%s", goal ), true, false))
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

func mapStringOptions(vs []string, f func(string) *slack.OptionBlockObject) []*slack.OptionBlockObject {
	vsm := make([]*slack.OptionBlockObject, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}


func mapToEvents(day string, blocks []slack.Block) []model.Event {
	var events = make([]model.Event, len(blocks))
	for index, block := range blocks {
		events[index] = convertToEvent(day, block)
	}
	return events
}

func mapToGoals(blocks []slack.Block) map[string][]model.Goal {
	return map[string][]model.Goal{}
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

func header(title string) []slack.Block {
	return []slack.Block{
		slack.NewDividerBlock(),
		slack.NewContextBlock("", slack.NewTextBlockObject(slack.MarkdownType,  title, false, false)),
		slack.NewDividerBlock(),
	}
}

func sectionHeader(title string) []slack.Block {
	return []slack.Block{
		slack.NewDividerBlock(),
		slack.NewContextBlock("", slack.NewTextBlockObject(slack.MarkdownType,  title, false, false)),
		slack.NewDividerBlock(),
	}
}

func buildEventsBlock(todayEvents []slack.Block, tomorrowEvents []slack.Block) []slack.Block {
	blocks := header(TodaysEventsHeader)
	blocks = append(blocks, todayEvents...)

	blocks = append(blocks, header(TomorrowsEventsHeader)...)
	blocks = append(blocks, tomorrowEvents...)

	return blocks
}

func buildGoalsBlock(goals []slack.Block) []slack.Block {
	blocks := sectionHeader(GoalsHeader)
	blocks = append(blocks, goals...)
	return blocks
}
