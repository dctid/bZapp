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

const TodaysEventsHeader = "Today's Events"
const TomorrowsEventsHeader = "Tomorrow's Events"
const GoalsHeader = "Goals"
const RemoveGoalActionId = "remove_goal"
const GoalCategoryDropdownPrefix = "goal-"

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

func BuildNewEvent(index int, values map[string]map[string]slack.BlockAction) model.Event {
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

func BuildNewGoalSectionBlock(index int, values map[string]map[string]slack.BlockAction) (string, string) {
	log.Printf("index: %d, goals values: %v", index, values)
	category := values[fmt.Sprintf("%s-%d", AddGoalCategoryInputBlock, index)][AddGoalCategoryActionId].SelectedOption.Value
	goal := values[fmt.Sprintf("%s-%d", AddGoalInputBlock, index)][AddGoalActionId].Value

	fmt.Printf("Add Goal category: %s, goal: %s\n", category, goal)

	return strings.TrimPrefix(category, GoalCategoryDropdownPrefix), goal
}

func ConvertToEventsBlocks(editable bool, events model.Events) ([]slack.Block, []slack.Block) {
	return convertToSectionBlocks(editable, events.TodaysEvents),
		convertToSectionBlocks(editable, events.TomorrowsEvents)
}

func ConvertToGoalBlocks(editable bool, category string, goals []model.Goal) []slack.Block {
	numEvents := len(goals)
	if numEvents == 0 {
		return NoGoalsYetSection
	}
	convertedBlocks := make([]slack.Block, numEvents)

	for index, goal := range goals {
		convertedBlocks[index] = slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, goal.Value, false, false),
			nil,
			getGoalRemoveButton(editable, category, goal),
			slack.SectionBlockOptionBlockID(goal.Id),
		)
	}
	return convertedBlocks
}

