package view

import (
	"encoding/json"
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
	"log"
)

const AddGoalCategoryInputBlock = "add_goal_category_input_block"
const AddGoalInputBlock = "add_goal_input_block"

const AddGoalCategoryActionId = "add_goal_category"
const AddGoalActionId = "add_goal"
const EditGoalsTitle = "bZapp - Edit Goals"

var GoalCategories = []string{"Customer Questions?", "Team Needs", "Learnings", "Questions?", "Other"}

func NewEditGoalsModal(index int, goals map[string][]model.Goal) slack.ModalViewRequest {

	return slack.ModalViewRequest{
		Type:   slack.VTModal,
		Title:  slack.NewTextBlockObject(slack.PlainTextType, EditGoalsTitle, true, false),
		Close:  slack.NewTextBlockObject(slack.PlainTextType, "Back", true, false),
		Submit: slack.NewTextBlockObject(slack.PlainTextType, "Add", true, false),
		Blocks: slack.Blocks{
			BlockSet: buildEditGoalsBlock(index, goals),
		},
		NotifyOnClose: true,
	}
}

func buildEditGoalsBlock(index int, goals map[string][]model.Goal) []slack.Block {

	var blocks []slack.Block

	for _, category := range GoalCategories {
		blocks = append(blocks, header(fmt.Sprintf("*%s*", category))...)
		blocks = append(blocks, ConvertToGoalsWithRemoveButton(category, goals[category])...)
	}

	blocks = append(blocks, actionsBlock(index)...)

	return blocks
}

func actionsBlock(index int) []slack.Block {

	categoryOptions := mapStringOptions(GoalCategories, goalCategoryOption)
	blocks := []slack.Block{
		slack.NewDividerBlock(),
		slack.NewInputBlock(fmt.Sprintf("%s-%d", AddGoalCategoryInputBlock, index), slack.NewTextBlockObject(slack.PlainTextType, "Goal to Add to", true, false),
			slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, slack.NewTextBlockObject(slack.PlainTextType, "Choose Goal", true, false), AddGoalCategoryActionId, categoryOptions...),
		),
		slack.NewInputBlock(fmt.Sprintf("%s-%d", AddGoalInputBlock, index), slack.NewTextBlockObject(slack.PlainTextType, "Goal to Add", true, false),
			slack.NewPlainTextInputBlockElement(slack.NewTextBlockObject(slack.PlainTextType, "Goal", false, false), AddGoalActionId),
		),
	}

	return blocks
}

func OpenEditGoalsModalFromSummaryModal(payload InteractionPayload) slack.ModalViewRequest {
	//todaysEvents, tomorrowsEvents, _ := ExtractModel(payload.View.Blocks.BlockSet)
	//todaysSectionBlocks, tomorrowsSectionEvents := ConvertToEventsWithRemoveButton(todaysEvents, tomorrowsEvents)
	index := ExtractInputIndex(payload.View.Blocks.BlockSet)

	modalRequest := NewEditGoalsModal(index+1, map[string][]model.Goal{})
	return modalRequest
}

func AddGoalToEditModal(payload InteractionPayload) *slack.ViewSubmissionResponse {
	action := payload.View.State.Values[AddEventDayInputBlock][AddEventDayActionId]
	marshal, _ := json.Marshal(action)
	fmt.Printf("bAdd Event button pressed by user %s with value %v\n", payload.User.Name, string(marshal))



	index := ExtractInputIndex(payload.View.Blocks.BlockSet)
	_, _, goals := ExtractModel(payload.View.Blocks.BlockSet)
	category, goal := BuildNewGoalSectionBlock(index, payload.View.State.Values)

	goals[category] = append(goals[category], model.Goal{
		Id:    model.Hash(),
		Value: goal,
	})

	log.Printf("goals %+v", goals)

	modalRequest := NewEditGoalsModal(index+1, goals)
	return slack.NewUpdateViewSubmissionResponse(&modalRequest)
}
