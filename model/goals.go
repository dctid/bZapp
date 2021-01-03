package model

import "errors"

type Goals map[string][]Goal


type Goal struct {
	Id string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
}

func (goals Goals) RemoveGoal(id string) Goals {
	category, index, err := findInMapById(id, goals)
	if err != nil {
		return goals
	}

	return removeFromMap(goals, category, index)
}

func (goals Goals) AddGoal(category string, newGoal string) Goals {
	if goals == nil{
		goals = make(Goals)
	}
	goals[category] = append(goals[category],Goal{
		Id:    Hash(),
		Value: newGoal,
	})
	return goals
}

func removeFromMap(goals Goals, category string, index int) Goals{
	listToRemoveFrom := goals[category]
	copy(listToRemoveFrom[index:], listToRemoveFrom[index+1:])
	listToRemoveFrom[len(listToRemoveFrom)-1] = Goal{}
	goals[category] = listToRemoveFrom[:len(listToRemoveFrom)-1]

	return goals
}

func findInMapById(id string, goals Goals) (string, int, error) {
	for category, goalsForCategory := range goals {
		for index, goal := range goalsForCategory {
			if goal.Id == id {
				return category, index, nil
			}
		}
	}
	return "", 0, errors.New("not found")
}