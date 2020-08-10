package modal

import (
	"fmt"
	"github.com/slack-go/slack"
	"strings"
)

func NewSummaryModal(todayEvents []*slack.SectionBlock, tomorrowEvents []*slack.SectionBlock) slack.ModalViewRequest {
	titleText := slack.NewTextBlockObject("plain_text", "bZapp", true, false)
	submitText := slack.NewTextBlockObject("plain_text", "Submit", true, false)
	closeText := slack.NewTextBlockObject("plain_text", "Cancel", true, false)
	todayHeader := slack.NewContextBlock("", slack.NewTextBlockObject("mrkdwn", "*Today's Events*", false, false))
	tomorrowHeader := slack.NewContextBlock("", slack.NewTextBlockObject("mrkdwn", "*Tomorrow's Events*", false, false))

	actions := slack.NewActionBlock(
		"actions_block",
		slack.NewButtonBlockElement("edit_events", "edit_events", slack.NewTextBlockObject("plain_text", "Edit Events", true, false)),
	)

	blockSet := []slack.Block{
		slack.NewDividerBlock(),
		todayHeader,
		slack.NewDividerBlock(),
	}

	for _, event := range todayEvents {
		blockSet = append(blockSet, event)
	}

	blockSet = append(blockSet, slack.NewDividerBlock(),
		tomorrowHeader,
		slack.NewDividerBlock(),
	)

	for _, event := range tomorrowEvents {
		blockSet = append(blockSet, event)
	}
	blockSet = append(blockSet,
		slack.NewDividerBlock(),
		actions,
	)

	blocks := slack.Blocks{
		BlockSet: blockSet,
	}

	var modalRequest slack.ModalViewRequest
	modalRequest.Type = "modal"
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks

	return modalRequest
}
const NoEventsText = "_No events yet_"
var NoEventYetSection = []*slack.SectionBlock{slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", NoEventsText, false, false), nil, nil)}

func MinOption(num int) *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(fmt.Sprintf("min-%d", num), slack.NewTextBlockObject("plain_text", fmt.Sprintf(func() string {
		if num < 10 {
			return "0%d"
		} else {
			return "%d"
		}
	}(), num), true, false))
}

func HourOption(num int) *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(fmt.Sprintf("hour-%d", num), slack.NewTextBlockObject("plain_text", fmt.Sprintf("%d %s", num, func() string {
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



func EventSection(index int, title string, hour string, mins string) *slack.SectionBlock {
	id := fmt.Sprintf("remove_today_%d", index)
	return slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", fmt.Sprintf("%s:%s %s", strings.Fields(hour)[0], mins, title), false, false),
		nil,
		slack.NewAccessory(slack.NewButtonBlockElement(id, id, slack.NewTextBlockObject("plain_text", "Remove", true, false))))
}
