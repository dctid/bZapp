package model

import "errors"

type Goal struct {
	Id string
	Value string
}

func RemoveGoal(id string, goals map[string][]Goal) map[string][]Goal {
	category, index, err := findInMapById(id, goals)
	if err != nil {
		return goals
	}

	return removeFromMap(goals, category, index)
}

func removeFromMap(goals map[string][]Goal, category string, index int) map[string][]Goal{
	listToRemoveFrom := goals[category]
	copy(listToRemoveFrom[index:], listToRemoveFrom[index+1:])
	listToRemoveFrom[len(listToRemoveFrom)-1] = Goal{}
	goals[category] = listToRemoveFrom[:len(listToRemoveFrom)-1]

	return goals
}

func findInMapById(id string, goals map[string][]Goal) (string, int, error) {
	for category, goalsForCategory := range goals {
		for index, goal := range goalsForCategory {
			if goal.Id == id {
				return category, index, nil
			}
		}
	}
	return "", 0, errors.New("not found")
}