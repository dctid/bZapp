package view

import (
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
	"strconv"
	"strings"
)

const EditEventsTitle = "bZapp - Edit Events"
const EventsHeader = "Events"
const TodayEventsHeader = "Today"
const TomorrowEventsHeader = "Tomorrow"
const MondayEventsHeader = "Monday"

const TodayOptionValue = "today"
const TomorrowOptionValue = "tomorrow"

const EditEventsActionId = "edit_events"
const RemoveEventActionId = "remove_event"
const AddEventTitleActionId = "add_event_title"
const AddEventDayActionId = "add_event_day"
const AddEventHoursActionId = "add_event_hour"
const AddEventMinsActionId = "add_event_mins"
const AddEventTitleInputBlock = "add_event_title_input_block"
const AddEventDayInputBlock = "add_event_day_input_block"
const AddEventHoursInputBlock = "add_event_hours_input_block"
const AddEventMinsInputBlock = "add_event_mins_input_block"

const EditGoalsTitle = "bZapp - Edit Goals"
const GoalsHeader = "Goals"

const EditGoalsActionId = "edit_goals"
const RemoveGoalActionId = "remove_goal"
const AddGoalActionId = "add_goal"
const AddGoalCategoryInputBlock = "add_goal_category_input_block"
const AddGoalInputBlock = "add_goal_input_block"
const AddGoalCategoryActionId = "add_goal_category"

var GoalCategories = []string{"Customer Questions?", "Team Needs", "Learnings", "Questions?", "Other"}


var NoEventYetSection = []slack.Block{slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, "_No events yet_", false, false), nil, nil)}

var NoGoalsYetSection = []slack.Block{slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, "_No goals yet_", false, false), nil, nil)}

type ResponseUrl struct {
	BlockId     string `json:"block_id"`
	ActionId    string `json:"action_id"`
	ChannelId   string `json:"channel_id"`
	ResponseUrl string `json:"response_url"`
}


func BuildNewEvent(index int, values map[string]map[string]slack.BlockAction) *model.Event {
	eventTitle := values[fmt.Sprintf("%s-%d", AddEventTitleInputBlock, index)][AddEventTitleActionId].Value
	eventDay := values[fmt.Sprintf("%s-%d", AddEventDayInputBlock, index)][AddEventDayActionId].SelectedOption.Value
	eventHours, _ := strconv.Atoi(strings.Split(values[fmt.Sprintf("%s-%d", AddEventHoursInputBlock, index)][AddEventHoursActionId].SelectedOption.Text.Text, " ")[0])
	eventMins, _ := strconv.Atoi(values[fmt.Sprintf("%s-%d", AddEventMinsInputBlock, index)][AddEventMinsActionId].SelectedOption.Text.Text)

	fmt.Printf("Add Event title: %s, day: %s, hour: %d, mins: %d\n", eventTitle, eventDay, eventHours, eventMins)

	return &model.Event{
		Id:    model.Hash(),
		Title: eventTitle,
		Day:   eventDay,
		Hour:  eventHours,
		Min:   eventMins,
		AmPm:  amOrPm(eventHours),
	}
}

func BuildNewGoalSectionBlock(index int, values map[string]map[string]slack.BlockAction) (string, string) {
	category := values[fmt.Sprintf("%s-%d", AddGoalCategoryInputBlock, index)][AddGoalCategoryActionId].SelectedOption.Value
	goal := values[fmt.Sprintf("%s-%d", AddGoalInputBlock, index)][AddGoalActionId].Value

	fmt.Printf("Add Goal category: %s, goal: %s\n", category, goal)

	return category, goal
}


