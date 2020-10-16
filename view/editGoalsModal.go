package view

import (
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

func NewEditGoalsModal(updatedModel *model.Model) *slack.ModalViewRequest {

	return &slack.ModalViewRequest{
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
		blocks = append(blocks, ConvertToGoalBlocks(true, category, goals[category])...)
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

func OpenEditGoalsModalFromSummaryModal(currentModel *model.Model) *slack.ModalViewRequest {
	currentModel.Index++
	return NewEditGoalsModal(currentModel)
}

func AddGoalToEditModal(values map[string]map[string]slack.BlockAction, currentModel *model.Model) *slack.ViewSubmissionResponse {

	category, goal := BuildNewGoalSectionBlock(currentModel.Index, values)
	currentModel.Goals = currentModel.Goals.AddGoal(category, goal)
	currentModel.Index++

	modalRequest := NewEditGoalsModal(currentModel)
	return slack.NewUpdateViewSubmissionResponse(modalRequest)
}

func RemoveGoalFromEditModal(blockIdToDelete string, currentModel *model.Model) *slack.ModalViewRequest {
	currentModel.Goals = currentModel.Goals.RemoveGoal(blockIdToDelete)
	return NewEditGoalsModal(currentModel)
}