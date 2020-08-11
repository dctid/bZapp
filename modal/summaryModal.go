package modal

import (
	"fmt"
	"github.com/slack-go/slack"
	"strings"
)

func NewSummaryModal(todayEvents []*slack.SectionBlock, tomorrowEvents []*slack.SectionBlock) slack.ModalViewRequest {
	titleText := slack.NewTextBlockObject(slack.PlainTextType, "bZapp", true, false)
	submitText := slack.NewTextBlockObject(slack.PlainTextType, "Submit", true, false)
	closeText := slack.NewTextBlockObject(slack.PlainTextType, "Cancel", true, false)
	blocks := slack.Blocks{BlockSet: BuildEventBlocks(todayEvents, tomorrowEvents)}

	var modalRequest slack.ModalViewRequest
	modalRequest.Type = slack.VTModal
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks

	return modalRequest
}

func BuildEventBlocks(todayEvents []*slack.SectionBlock, tomorrowEvents []*slack.SectionBlock) []slack.Block {
	todayHeader := slack.NewContextBlock("", slack.NewTextBlockObject(slack.MarkdownType, "*Today's Events*", false, false))
	tomorrowHeader := slack.NewContextBlock("", slack.NewTextBlockObject(slack.MarkdownType, "*Tomorrow's Events*", false, false))

	actions := slack.NewActionBlock(
		"actions_block",
		slack.NewButtonBlockElement("edit_events", "edit_events", slack.NewTextBlockObject(slack.PlainTextType, "Edit Events", true, false)),
	)
	blocks := []slack.Block{
		slack.NewDividerBlock(),
		todayHeader,
		slack.NewDividerBlock(),
	}

	for _, event := range todayEvents {
		blocks = append(blocks, event)
	}

	blocks = append(blocks, slack.NewDividerBlock(),
		tomorrowHeader,
		slack.NewDividerBlock(),
	)

	for _, event := range tomorrowEvents {
		blocks = append(blocks, event)
	}
	blocks = append(blocks,
		slack.NewDividerBlock(),
		actions,
	)

	return blocks
}

const NoEventsText = "_No events yet_"

var NoEventYetSection = []*slack.SectionBlock{slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, NoEventsText, false, false), nil, nil)}

func MinOption(num int) *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(fmt.Sprintf("min-%d", num), slack.NewTextBlockObject(slack.PlainTextType, fmt.Sprintf(func() string {
		if num < 10 {
			return "0%d"
		} else {
			return "%d"
		}
	}(), num), true, false))
}

func HourOption(num int) *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(fmt.Sprintf("hour-%d", num), slack.NewTextBlockObject(slack.PlainTextType, fmt.Sprintf("%d %s", num, func() string {
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

const RemoveEventActionId = "remove_event"
func EventSection(day string, index int, title string, hour string, mins string) *slack.SectionBlock {
	id := fmt.Sprintf("remove_%s_%d", day, index)
	return slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("%s:%s %s", strings.Fields(hour)[0], mins, title), false, false),
		nil,
		slack.NewAccessory(slack.NewButtonBlockElement(RemoveEventActionId, id, slack.NewTextBlockObject(slack.PlainTextType, "Remove", true, false))))
}
