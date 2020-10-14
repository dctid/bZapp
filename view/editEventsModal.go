package view

import (
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
)

func NewEditEventsModal(updatedModel model.Model) slack.ModalViewRequest {
	todaysSectionBlocks, tomorrowsSectionBlocks := ConvertToEventsWithRemoveButton(updatedModel.Events.TodaysEvents, updatedModel.Events.TomorrowsEvents)
	return slack.ModalViewRequest{
		Type:   slack.VTModal,
		Title:  slack.NewTextBlockObject(slack.PlainTextType, "bZapp - Edit Events", true, false),
		Close:  slack.NewTextBlockObject(slack.PlainTextType, "Back", true, false),
		Submit: slack.NewTextBlockObject(slack.PlainTextType, "Add", true, false),
		Blocks: slack.Blocks{
			BlockSet: buildSummaryEventBlocks(updatedModel.Index, todaysSectionBlocks, tomorrowsSectionBlocks),
		},
		NotifyOnClose: true,
		PrivateMetadata: updatedModel.ConvertModelToJson(),
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
			slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, slack.NewTextBlockObject(slack.PlainTextType, "Select hour", true, false), AddEventHoursActionId, hourOptions...),
		),
		slack.NewInputBlock(fmt.Sprintf("%s-%d", AddEventMinsInputBlock, index), slack.NewTextBlockObject(slack.PlainTextType, "Minutes", true, false),
			slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, slack.NewTextBlockObject(slack.PlainTextType, "Select Minutes", true, false), AddEventMinsActionId, minOptions...),
		),
	}
}

func AddEventToEditModal(payload InteractionPayload) *slack.ViewSubmissionResponse {
	//action := payload.View.State.Values[AddEventDayInputBlock][AddEventDayActionId]
	//marshal, _ := json.Marshal(action)
	//fmt.Printf("Add Event button pressed by user %s with value %v\n", payload.User.Name, string(marshal))

	currentModel := ExtractModel(payload.View.Blocks.BlockSet)
	index := ExtractInputIndex(payload.View.Blocks.BlockSet)
	newEvent := BuildNewEvent(index, payload.View.State.Values)
	switch newEvent.Day {
	case TodayOptionValue:
		currentModel.Events.TodaysEvents = model.AddEventInOrder(newEvent, currentModel.Events.TodaysEvents)
	case TomorrowOptionValue:
		currentModel.Events.TomorrowsEvents = model.AddEventInOrder(newEvent, currentModel.Events.TomorrowsEvents)
	}
	currentModel.Index = index + 1

	modalRequest := NewEditEventsModal(currentModel)
	return slack.NewUpdateViewSubmissionResponse(&modalRequest)
}

func OpenEditEventModalFromSummaryModal(payload InteractionPayload) slack.ModalViewRequest {
	currentModel := ExtractModel(payload.View.Blocks.BlockSet)
	currentModel.Index = ExtractInputIndex(payload.View.Blocks.BlockSet) + 1

	modalRequest := NewEditEventsModal(currentModel)
	return modalRequest
}

func RemoveEventFromEditModal(payload InteractionPayload) slack.ModalViewRequest {
	blockIdToDelete := payload.ActionCallback.BlockActions[0].BlockID

	currentModel := ExtractModel(payload.View.Blocks.BlockSet)
	currentModel.Events.TodaysEvents = model.RemoveEvent(blockIdToDelete, currentModel.Events.TodaysEvents)
	currentModel.Events.TomorrowsEvents = model.RemoveEvent(blockIdToDelete, currentModel.Events.TomorrowsEvents)
	currentModel.Index = ExtractInputIndex(payload.View.Blocks.BlockSet)

	modalRequest := NewEditEventsModal(currentModel)
	return modalRequest
}
