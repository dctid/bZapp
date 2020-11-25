package view

import (
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
)

func NewEditEventsModal(updatedModel *model.Model) *slack.ModalViewRequest {
	return &slack.ModalViewRequest{
		Type:   slack.VTModal,
		Title:  slack.NewTextBlockObject(slack.PlainTextType, EditEventsTitle, true, false),
		Close:  slack.NewTextBlockObject(slack.PlainTextType, "Back", true, false),
		Submit: slack.NewTextBlockObject(slack.PlainTextType, "Add", true, false),
		Blocks: slack.Blocks{
			BlockSet: buildSummaryEventBlocks(updatedModel.Index, updatedModel.Events),
		},
		NotifyOnClose: true,
		PrivateMetadata: updatedModel.ConvertMetadataToJson(),
	}
}

func buildSummaryEventBlocks(index int, events model.Events) []slack.Block {

	blocks := buildEventsBlock(true, events)
	blocks = append(blocks, addEventsActions(index)...)

	return blocks
}

func addEventsActions(index int) []slack.Block {
	hours := []int{9, 10, 11, 12, 1, 2, 3, 4}
	hourOptions := mapOptions(hours, hourOption)

	mins := []int{0, 15, 30, 45}
	minOptions := mapOptions(mins, minOption)
	return []slack.Block{
		slack.NewDividerBlock(),
		slack.NewInputBlock(fmt.Sprintf("%s-%d", AddEventTitleInputBlock, index), slack.NewTextBlockObject(slack.PlainTextType, "Add Event", false, false),
			slack.NewPlainTextInputBlockElement(slack.NewTextBlockObject(slack.PlainTextType, "Title", false, false), AddEventTitleActionId),
		),
		slack.NewInputBlock(fmt.Sprintf("%s-%d", AddEventDayInputBlock, index), slack.NewTextBlockObject(slack.PlainTextType, "Day", true, false),
			slack.NewRadioButtonsBlockElement(AddEventDayActionId,
				slack.NewOptionBlockObject(TodayOptionValue, slack.NewTextBlockObject(slack.PlainTextType, "Today", true, false)),
				slack.NewOptionBlockObject(TomorrowOptionValue, slack.NewTextBlockObject(slack.PlainTextType, "Tomorrow", true, false))),
		),
		slack.NewInputBlock(fmt.Sprintf("%s-%d", AddEventHoursInputBlock, index), slack.NewTextBlockObject(slack.PlainTextType, "Hour", true, false),
			slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, slack.NewTextBlockObject(slack.PlainTextType, "Select hour", true, false), AddEventHoursActionId, hourOptions...),
		),
		slack.NewInputBlock(fmt.Sprintf("%s-%d", AddEventMinsInputBlock, index), slack.NewTextBlockObject(slack.PlainTextType, "Minutes", true, false),
			slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, slack.NewTextBlockObject(slack.PlainTextType, "Select Minutes", true, false), AddEventMinsActionId, minOptions...),
		),
	}
}


