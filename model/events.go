package model

import (
	"errors"
	"fmt"
)

type Events map[string][]Event

const (
	TodaysEvents    = "today"
	TomorrowsEvents = "tomorrow"
)

type Event struct {
	Id    string
	Title string
	Day   string
	Hour  int
	Min   int
	AmPm  string
}

func (event Event) ToString() string {
	return fmt.Sprintf("%d:%s %s", event.Hour, event.normalizeMins(), event.Title)
}

func (events Events) IsEmpty() bool {
	return len(events[TodaysEvents]) == 0 && len(events[TomorrowsEvents]) == 0
}

func (event Event) normalizeMins() string {
	if event.Min < 10 {
		return fmt.Sprintf("0%d", event.Min)
	} else {
		return fmt.Sprintf("%d", event.Min)
	}
}

func (events Events) AddEvent(newEvent *Event) Events {
	events[newEvent.Day] = addEventInOrder(newEvent, events[newEvent.Day])
	return events
}

func (events Events) ConvertToDate() Events {
	if events != nil {
		days := Days()
		return Events{
			days.Today:      events[TodaysEvents],
			days.NextBizDay: events[TomorrowsEvents],
		}
	}
	return nil
}

func (events Events) ConvertFromDates() Events {
	if events != nil {
		result := Events{}
		days := Days()
		today := events[days.Today]
		if len(today) > 0 {
			result[TodaysEvents] = today
		}
		nextBizDayEvents := events[days.NextBizDay]
		if len(nextBizDayEvents) > 0 {
			result[TomorrowsEvents] = nextBizDayEvents
		}
		return result
	}
	return nil
}

func addEventInOrder(event *Event, events []Event) []Event {

	indexToInsertAt := findIndexToInsertAt(event, events)

	result := append(events, *event)
	copy(result[indexToInsertAt+1:], result[indexToInsertAt:])
	result[indexToInsertAt] = *event

	return result
}

func findIndexToInsertAt(event *Event, events []Event) int {
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

func (events Events) RemoveEvent(id string) Events {
	return Events{
		TodaysEvents:    removeEvent(id, events[TodaysEvents]),
		TomorrowsEvents: removeEvent(id, events[TomorrowsEvents]),
	}
}

func removeEvent(id string, events []Event) []Event {
	index, err := findById(id, events)
	if err != nil {
		return events
	}

	return removeAtIndex(events, index)
}

func removeAtIndex(events []Event, index int) []Event {
	copy(events[index:], events[index+1:])
	events[len(events)-1] = Event{}
	events = events[:len(events)-1]
	return events
}
func findById(id string, events []Event) (int, error) {
	for index, event := range events {
		if event.Id == id {
			return index, nil
		}
	}
	return 0, errors.New("not found")
}
