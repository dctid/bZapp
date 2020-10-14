package view

import (
	"encoding/json"
	"fmt"
	"github.com/dctid/bZapp/model"
	"github.com/slack-go/slack"
)

const AddGoalCategoryInputBlock = "add_goal_category_input_block"
const AddGoalInputBlock = "add_goal_input_block"

const AddGoalCategoryActionId = "add_goal_category"
const AddGoalActionId = "add_goal"
const EditGoalsTitle = "bZapp - Edit Goals"

var GoalCategories = []string{"Customer Questions?", "Team Needs", "Learnings", "Questions?", "Other"}

func NewEditGoalsModal(updatedModel model.Model) slack.ModalViewRequest {

	return slack.ModalViewRequest{
		Type:   slack.VTModal,
		Title:  slack.NewTextBlockObject(slack.PlainTextType, EditGoalsTitle, true, false),
		Close:  slack.NewTextBlockObject(slack.PlainTextType, "Back", true, false),
		Submit: slack.NewTextBlockObject(slack.PlainTextType, "Add", true, false),
		Blocks: slack.Blocks{
			BlockSet: buildEditGoalsBlock(updatedModel.Index, updatedModel.Goals),
		},
		NotifyOnClose:   true,
		PrivateMetadata: updatedModel.ConvertModelToJson(),
	}
}

func buildEditGoalsBlock(index int, goals model.Goals) []slack.Block {

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
	currentModel := ExtractModel(payload.View.Blocks.BlockSet)
	//todaysSectionBlocks, tomorrowsSectionEvents := ConvertToEventsWithRemoveButton(todaysEvents, tomorrowsEvents)
	index := ExtractInputIndex(payload.View.Blocks.BlockSet)
	currentModel.Index = index+1
	modalRequest := NewEditGoalsModal(currentModel)
	return modalRequest
}

func AddGoalToEditModal(payload InteractionPayload) *slack.ViewSubmissionResponse {
	action := payload.View.State.Values[AddEventDayInputBlock][AddEventDayActionId]
	marshal, _ := json.Marshal(action)
	fmt.Printf("bAdd Event button pressed by user %s with value %v\n", payload.User.Name, string(marshal))

	index := ExtractInputIndex(payload.View.Blocks.BlockSet)
	currentModel := ExtractModel(payload.View.Blocks.BlockSet)

	category, goal := BuildNewGoalSectionBlock(index, payload.View.State.Values)

	currentModel.Goals[category] = append(currentModel.Goals[category], model.Goal{
		Id:    model.Hash(),
		Value: goal,
	})
	currentModel.Index = index+1

	modalRequest := NewEditGoalsModal(currentModel)
	return slack.NewUpdateViewSubmissionResponse(&modalRequest)
}

func RemoveGoalFromEditModal(payload InteractionPayload) slack.ModalViewRequest {
	blockIdToDelete := payload.ActionCallback.BlockActions[0].BlockID

	currentModel := ExtractModel(payload.View.Blocks.BlockSet)

	currentModel.Goals = model.RemoveGoal(blockIdToDelete, currentModel.Goals)

	currentModel.Index = ExtractInputIndex(payload.View.Blocks.BlockSet)

	modalRequest := NewEditGoalsModal(currentModel)
	return modalRequest
}