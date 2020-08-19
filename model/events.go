package model

import (
	"fmt"
)

type Event struct {
	Title string
	Day   string
	Hour  int
	Min   int
	AmPm  string
}

func (event Event) ToString() string {
	return fmt.Sprintf("%d:%s %s", event.Hour, event.normalizeMins(), event.Title)
}

func (event Event) normalizeMins() string {
	if event.Min < 10 {
		return fmt.Sprintf("0%d", event.Min)
	} else {
		return fmt.Sprintf("%d", event.Min)
	}
}

func AddEventInOrder(event Event, events []Event) []Event {

	indexToInsertAt := findIndexToInsertAt(event, events)

	result := append(events, event)
	copy(result[indexToInsertAt+1:], result[indexToInsertAt:])
	result[indexToInsertAt] = event

	return result
}

func findIndexToInsertAt(event Event, events []Event) int {
	for index, e := range events {
		if event.greaterThan(e) {
			return index
		}
	}
	return len(events)
}

func (event Event) greaterThan(other Event) bool {
	return event.calcEventValue() < other.calcEventValue()
}
func (event Event) calcEventValue() int {
	if event.AmPm == "PM" {
		return (event.Hour+12)*100 + event.Min
	}
	return event.Hour*100 + event.Min
}
