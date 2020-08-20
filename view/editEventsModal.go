package view

import (
	"encoding/json"
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
)

func NewEditEventsModal(index int, todayEvents []slack.Block, tomorrowEvents []slack.Block) slack.ModalViewRequest {
	return slack.ModalViewRequest{
		Type:   slack.VTModal,
		Title:  slack.NewTextBlockObject(slack.PlainTextType, "bZapp - Edit Events", true, false),
		Close:  slack.NewTextBlockObject(slack.PlainTextType, "Back", true, false),
		Submit: slack.NewTextBlockObject(slack.PlainTextType, "Add", true, false),
		Blocks: slack.Blocks{
			BlockSet: buildSummaryEventBlocks(index, todayEvents, tomorrowEvents),
		},
		NotifyOnClose: true,
	}
}

func buildSummaryEventBlocks(index int, todayEvents []slack.Block, tomorrowEvents []slack.Block) []slack.Block {

	blocks := buildEventsBlock(todayEvents, tomorrowEvents)
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
			slack.NewOptionsSelectBlockElement("static_select", slack.NewTextBlockObject(slack.PlainTextType, "Select hour", true, false), AddEventHoursActionId, hourOptions...),
		),
		slack.NewInputBlock(fmt.Sprintf("%s-%d", AddEventMinsInputBlock, index), slack.NewTextBlockObject(slack.PlainTextType, "Minutes", true, false),
			slack.NewOptionsSelectBlockElement("static_select", slack.NewTextBlockObject(slack.PlainTextType, "Select Minutes", true, false), AddEventMinsActionId, minOptions...),
		),
	}
}

func AddEventToEditModal(payload InteractionPayload) *slack.ViewSubmissionResponse {
	action := payload.View.State.Values[AddEventDayInputBlock][AddEventDayActionId]
	marshal, _ := json.Marshal(action)
	fmt.Printf("aAdd Event button pressed by user %s with value %v\n", payload.User.Name, string(marshal))

	index := ExtractInputIndex(payload.View.Blocks.BlockSet)
	todaysEvents, tomorrowsEvents, _ := ExtractModel(payload.View.Blocks.BlockSet)

	newEvent := BuildNewEventSectionBlock(index, payload.View.State.Values)
	switch newEvent.Day {
	case TodayOptionValue:
		todaysEvents = model.AddEventInOrder(newEvent, todaysEvents)
	case TomorrowOptionValue:
		tomorrowsEvents = model.AddEventInOrder(newEvent, tomorrowsEvents)
	}
	todaysSectionBlocks, tomorrowsSectionBlocks := ConvertToEventsWithRemoveButton(todaysEvents, tomorrowsEvents)

	modalRequest := NewEditEventsModal(index+1, todaysSectionBlocks, tomorrowsSectionBlocks)
	return slack.NewUpdateViewSubmissionResponse(&modalRequest)
}

func OpenEditEventModalFromSummaryModal(payload InteractionPayload) slack.ModalViewRequest {
	todaysEvents, tomorrowsEvents, _ := ExtractModel(payload.View.Blocks.BlockSet)
	todaysSectionBlocks, tomorrowsSectionEvents := ConvertToEventsWithRemoveButton(todaysEvents, tomorrowsEvents)
	index := ExtractInputIndex(payload.View.Blocks.BlockSet)

	modalRequest := NewEditEventsModal(index+1, todaysSectionBlocks, tomorrowsSectionEvents)
	return modalRequest
}

func RemoveEventFromEditModal(payload InteractionPayload) slack.ModalViewRequest {
	blockIdToDelete := payload.ActionCallback.BlockActions[0].BlockID

	todaysEvents, tomorrowsEvents, _ := ExtractModel(payload.View.Blocks.BlockSet)

	todaysEvents = model.RemoveEvent(blockIdToDelete, todaysEvents)
	tomorrowsEvents = model.RemoveEvent(blockIdToDelete, tomorrowsEvents)
	todaysSectionBlocks, tomorrowsSectionBlocks := ConvertToEventsWithRemoveButton(todaysEvents, tomorrowsEvents)

	index := ExtractInputIndex(payload.View.Blocks.BlockSet)

	modalRequest := NewEditEventsModal(index, todaysSectionBlocks, tomorrowsSectionBlocks)
	return modalRequest
}
