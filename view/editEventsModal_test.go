package view

import (
	"encoding/json"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/mocks"
	"github.com/dctid/bZapp/model"
	"github.com/dctid/bZapp/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEditEventsModal(t *testing.T) {

	model.Clock = mocks.NewMockClock("2020-12-02 08:48:21")

	todaysEvents := []model.Event{
		{
			Id:    "FakeId1",
			Title: "Standup",
			Day:   TodayOptionValue,
			Hour:  9,
			Min:   15,
			AmPm:  "AM",
		},
		{
			Id:    "FakeId2",
			Title: "IPM",
			Day:   TodayOptionValue,
			Hour:  11,
			Min:   30,
			AmPm:  "AM",
		},
		{
			Id:    "FakeId3",
			Title: "Retro",
			Day:   TodayOptionValue,
			Hour:  3,
			Min:   15,
			AmPm:  "PM",
		},
	}

	tomorrowsEvents := []model.Event{
		{
			Id:    "FakeId4",
			Title: "Standup",
			Day:   TomorrowOptionValue,
			Hour:  9,
			Min:   15,
			AmPm:  "AM",
		},
		{
			Id:    "FakeId5",
			Title: "User Interview",
			Day:   TomorrowOptionValue,
			Hour:  1,
			Min:   30,
			AmPm:  "PM",
		},
		{
			Id:    "FakeId6",
			Title: "Synthesis",
			Day:   TomorrowOptionValue,
			Hour:  3,
			Min:   00,
			AmPm:  "PM",
		},
	}

	result := NewEditEventsModal( &model.Model{
		Index:  666,
		Events: model.Events{
			model.TodaysEvents: todaysEvents,
			model.TomorrowsEvents: tomorrowsEvents,
		},
	}, &model.Metadata{ChannelId: "edit_events_modal_channel_id"})
	actualJson, _ := json.Marshal(result)
	expectedJsonString := test.ReadFile(t, "view/edit_events_modal.json")
	actualJsonString := format.PrettyJson(t, string(actualJson))

	assert.EqualValues(t, expectedJsonString, actualJsonString)
}

func TestNewEditEventsModalOnFriday(t *testing.T) {

	model.Clock = mocks.NewMockClock("2020-12-04 08:48:21")

	todaysEvents := []model.Event{
		{
			Id:    "FakeId1",
			Title: "Standup",
			Day:   TodayOptionValue,
			Hour:  9,
			Min:   15,
			AmPm:  "AM",
		},
		{
			Id:    "FakeId2",
			Title: "IPM",
			Day:   TodayOptionValue,
			Hour:  11,
			Min:   30,
			AmPm:  "AM",
		},
		{
			Id:    "FakeId3",
			Title: "Retro",
			Day:   TodayOptionValue,
			Hour:  3,
			Min:   15,
			AmPm:  "PM",
		},
	}

	tomorrowsEvents := []model.Event{
		{
			Id:    "FakeId4",
			Title: "Standup",
			Day:   TomorrowOptionValue,
			Hour:  9,
			Min:   15,
			AmPm:  "AM",
		},
		{
			Id:    "FakeId5",
			Title: "User Interview",
			Day:   TomorrowOptionValue,
			Hour:  1,
			Min:   30,
			AmPm:  "PM",
		},
		{
			Id:    "FakeId6",
			Title: "Synthesis",
			Day:   TomorrowOptionValue,
			Hour:  3,
			Min:   00,
			AmPm:  "PM",
		},
	}

	result := NewEditEventsModal( &model.Model{
		Index:  666,
		Events: model.Events{
			model.TodaysEvents: todaysEvents,
			model.TomorrowsEvents: tomorrowsEvents,
		},
	}, &model.Metadata{ChannelId: "edit_events_on_friday_modal_channel_id"})
	actualJson, _ := json.Marshal(result)
	expectedJsonString := test.ReadFile(t, "view/edit_events_modal_on_friday.json")
	actualJsonString := format.PrettyJson(t, string(actualJson))

	assert.EqualValues(t, expectedJsonString, actualJsonString)
}
