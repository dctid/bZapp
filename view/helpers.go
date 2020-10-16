package view

import (
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
)

func convertEventsToSectionBlocks(includeRemoveButton bool, events []model.Event) []slack.Block {

	convertedBlocks := make([]slack.Block, len(events))

	for index, event := range events {
		convertedBlocks[index] = slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf(":small_orange_diamond: %s", event.ToString()), false, false),
			nil,
			getRemoveButton(RemoveEventActionId, includeRemoveButton, event.Day, event.Id),
			slack.SectionBlockOptionBlockID(event.Id),
		)
	}
	return convertedBlocks
}

func convertToGoalBlocks(editable bool, category string, goals []model.Goal) []slack.Block {
	convertedBlocks := make([]slack.Block, len(goals))

	for index, goal := range goals {
		convertedBlocks[index] = slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf(":small_blue_diamond: %s", goal.Value), false, false),
			nil,
			getRemoveButton(RemoveGoalActionId, editable, category, goal.Id),
			slack.SectionBlockOptionBlockID(goal.Id),
		)
	}
	return convertedBlocks
}



func getRemoveButton(actionId string, includeRemoveButton bool, typeOfButton string, id string) *slack.Accessory {
	if includeRemoveButton {
		return slack.NewAccessory(
			slack.NewButtonBlockElement(
				actionId,
				fmt.Sprintf("remove_%s_%s", typeOfButton, id),
				slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false),
			),
		)
	}
	return nil
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
	return slack.NewOptionBlockObject(goal, slack.NewTextBlockObject(slack.PlainTextType, goal, true, false))
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
		slack.NewHeaderBlock( slack.NewTextBlockObject(slack.PlainTextType,  title, false, false)),
	}
}


func buildEventsBlock(editable bool, events model.Events) []slack.Block {
	blocks := sectionHeader(EventsHeader)
	if events.IsEmpty() {
		blocks = append(blocks, NoEventYetSection...)
	} else {
		blocks = addEvents(editable, blocks, TodayEventsHeader, events.TodaysEvents)
		blocks = addEvents(editable, blocks, TomorrowEventsHeader, events.TomorrowsEvents)
	}

	return blocks
}

func addEvents(editable bool, blocks []slack.Block, headerText string, events []model.Event) []slack.Block{
	if len(events) > 0 {
		blocks = append(blocks, header(markupBold(headerText))...)
		blocks = append(blocks, convertEventsToSectionBlocks(editable, events)...)
	}
	return blocks
}


func buildGoalsBlock(editable bool, goals model.Goals) []slack.Block {
	blocks := sectionHeader(GoalsHeader)
	blocks = append(blocks, convertGoalsToBlocks(editable, goals)...)
	return blocks
}

func convertGoalsToBlocks(editable bool, goals model.Goals) []slack.Block {
	var goalBlocks []slack.Block
	for _, category := range GoalCategories {
		if len(goals[category]) > 0 {
			goalBlocks = append(goalBlocks, header(fmt.Sprintf("*%s*", category))...)
			goalBlocks = append(goalBlocks, convertToGoalBlocks(editable, category, goals[category])...)
		}
	}

	if len(goalBlocks) == 0 {
		goalBlocks = NoGoalsYetSection
	}
	return goalBlocks
}


func markupBold(value string) string {
	return fmt.Sprintf("*%s*", value)
}
