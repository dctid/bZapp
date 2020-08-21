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

var GoalCategories = []string{"Customer Questions?", "Team Needs", "Learnings", "Questions?", "Other"}

func NewEditGoalsModal(index int,goalMap map[string][]model.Goal) slack.ModalViewRequest {

	return slack.ModalViewRequest{
		Type: slack.VTModal,
		Title: slack.NewTextBlockObject(slack.PlainTextType, "bZapp - Edit Goals", true, false),
		Close:  slack.NewTextBlockObject(slack.PlainTextType, "Back", true, false),
		Submit: slack.NewTextBlockObject(slack.PlainTextType, "Add", true, false),
		Blocks: slack.Blocks{
			BlockSet: buildEditGoalsBlock(index),
		},
		NotifyOnClose: true,
	}
}

func buildEditGoalsBlock(index int) []slack.Block {

	blocks := header("*Customer Questions?*")
	blocks = append(blocks, NoGoalsYetSection...)

	blocks = append(blocks, header("*Team Needs*")...)
	blocks = append(blocks, NoGoalsYetSection...)

	blocks = append(blocks, header("*Learnings*")...)
	blocks = append(blocks, NoGoalsYetSection...)

	blocks = append(blocks, header("*Questions?*")...)
	blocks = append(blocks, NoGoalsYetSection...)

	blocks = append(blocks, header("*Other*")...)
	blocks = append(blocks, NoGoalsYetSection...)

	blocks = append(blocks, actionsBlock(index)...)

	return blocks
}

func actionsBlock(index int) []slack.Block  {

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