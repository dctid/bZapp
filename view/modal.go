package view

import (
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
	"log"
	"strconv"
	"strings"
)

const AddEventTitleInputBlock = "add_event_title_input_block"
const AddEventDayInputBlock = "add_event_day_input_block"
const AddEventHoursInputBlock = "add_event_hours_input_block"
const AddEventMinsInputBlock = "add_event_mins_input_block"

const AddEventTitleActionId = "add_event_title"
const AddEventDayActionId = "add_event_day"
const AddEventHoursActionId = "add_event_hour"
const AddEventMinsActionId = "add_event_mins"

const TodayOptionValue = "today"
const TomorrowOptionValue = "tomorrow"

const EditEventsActionId = "edit_events"
const EditGoalsActionId = "edit_goals"
const RemoveEventActionId = "remove_event"

const TodaysEventsHeader = "*Today's Events*"
const TomorrowsEventsHeader = "*Tomorrow's Events*"
const GoalsHeader = "Goals"


const NoEventsText = "_No events yet_"
var NoEventYetSection = []slack.Block{slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, NoEventsText, false, false), nil, nil)}

const NoGoalsYetText = "_No goals yet_"
var NoGoalsYetSection = []slack.Block{slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, NoGoalsYetText, false, false), nil, nil)}

type ResponseUrl struct {
	BlockId     string `json:"block_id"`
	ActionId    string `json:"action_id"`
	ChannelId   string `json:"channel_id"`
	ResponseUrl string `json:"response_url"`
}

type InteractionPayload struct {
	slack.InteractionCallback
	ResponseUrls []ResponseUrl `json:"response_urls"`
}

func BuildNewEventSectionBlock(index int, values map[string]map[string]slack.BlockAction) model.Event {
	eventTitle := values[fmt.Sprintf("%s-%d", AddEventTitleInputBlock, index)][AddEventTitleActionId].Value
	eventDay := values[fmt.Sprintf("%s-%d", AddEventDayInputBlock, index)][AddEventDayActionId].SelectedOption.Value
	eventHours, _ := strconv.Atoi(strings.Split(values[fmt.Sprintf("%s-%d", AddEventHoursInputBlock, index)][AddEventHoursActionId].SelectedOption.Text.Text, " ")[0])
	eventMins, _ := strconv.Atoi(values[fmt.Sprintf("%s-%d", AddEventMinsInputBlock, index)][AddEventMinsActionId].SelectedOption.Text.Text)

	fmt.Printf("Add Event title: %s, day: %s, hour: %d, mins: %d\n", eventTitle, eventDay, eventHours, eventMins)

	return model.Event{
		Id:    model.Hash(),
		Title: eventTitle,
		Day:   eventDay,
		Hour:  eventHours,
		Min:   eventMins,
		AmPm:  amOrPm(eventHours),
	}
}

func ExtractModel(blocks []slack.Block) ([]model.Event, []model.Event, []model.Goal) {
	log.Println("New Events")

	contentBlockMap := groupSectionBlocks(blocks)
	todaysBlocks := contentBlockMap[TodaysEventsHeader]
	tomorrowsBlocks := contentBlockMap[TomorrowsEventsHeader]
	goalsBlocks := contentBlockMap[GoalsHeader]

	return mapToEvents(TodayOptionValue, todaysBlocks), mapToEvents(TomorrowOptionValue, tomorrowsBlocks), mapToGoals(goalsBlocks)
}

func ConvertToEventsWithRemoveButton(todaysEvents []model.Event, tomorrowsEvents []model.Event) ([]slack.Block, []slack.Block) {
	return convertToSectionBlocks(true, todaysEvents),
		convertToSectionBlocks(true, tomorrowsEvents)
}

func ConvertToEventsWithoutRemoveButton(todaysEvents []model.Event, tomorrowsEvents []model.Event) ([]slack.Block, []slack.Block) {
	return convertToSectionBlocks(false, todaysEvents),
		convertToSectionBlocks(false, tomorrowsEvents)
}

func ExtractInputIndex(blocks []slack.Block) int {
	for _, block := range blocks {
		if block.BlockType() == slack.MBTInput {
			inputBlock := block.(*slack.InputBlock)
			tokens := strings.Split(inputBlock.BlockID, "-")
			length := len(tokens)
			if length > 1 {
				index, _ := strconv.Atoi(tokens[length-1])
				return index
			}
			return 0
		}
	}
	return 0
}
