package view

import (
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
)

func convertToSectionBlocks(includeRemoveButton bool, events []model.Event) []slack.Block {

	numEvents := len(events)
	if numEvents == 0 {
		return NoEventYetSection
	}
	convertedBlocks := make([]slack.Block, numEvents)

	for index, event := range events {
		convertedBlocks[index] = slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf(":small_orange_diamond: %s", event.ToString()), false, false),
			nil,
			getRemoveButton(includeRemoveButton, event),
			slack.SectionBlockOptionBlockID(event.Id),
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
	todayEvents, tomorrowEvents := ConvertToEventsBlocks(editable, events)
	blocks := sectionHeader(EventsHeader)
	if len(events.TodaysEvents) > 0 {
		blocks = append(blocks, header(markupBold(TodaysEventsHeader))...)
		blocks = append(blocks, todayEvents...)

	}
	if len(events.TomorrowsEvents) >  0 {
		blocks = append(blocks, header(markupBold(TomorrowsEventsHeader))...)
		blocks = append(blocks, tomorrowEvents...)
	}
	if len(events.TodaysEvents) == 0 && len(events.TomorrowsEvents) == 0 {
		blocks = append(blocks, NoEventYetSection...)
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
			goalBlocks = append(goalBlocks, ConvertToGoalBlocks(editable, category, goals[category])...)
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

func markupItalics(value string) string {
	return fmt.Sprintf("_%s_", value)
}