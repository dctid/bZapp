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

func TestNewSummaryModal(t *testing.T) {
	model.Clock = mocks.NewMockClock("2020-12-02 08:48:21")

	testModel := &model.Model{
		Events: model.Events{
			model.TodaysEvents: []model.Event{
				{
					Title: "Standup",
					Day:   TodayOptionValue,
					Hour:  9,
					Min:   15,
					AmPm:  "AM",
				},
				{
					Title: "IPM",
					Day:   TodayOptionValue,
					Hour:  11,
					Min:   30,
					AmPm:  "AM",
				},
				{
					Title: "Retro",
					Day:   TodayOptionValue,
					Hour:  3,
					Min:   15,
					AmPm:  "PM",
				},
			},
			model.TomorrowsEvents: []model.Event{
				{
					Title: "Standup",
					Day:   TomorrowOptionValue,
					Hour:  9,
					Min:   15,
					AmPm:  "AM",
				},
				{
					Title: "User Interview",
					Day:   TomorrowOptionValue,
					Hour:  1,
					Min:   30,
					AmPm:  "PM",
				},
				{
					Title: "Synthesis",
					Day:   TomorrowOptionValue,
					Hour:  3,
					Min:   0,
					AmPm:  "PM",
				},
			},
		},
		Goals: nil,
	}

	testMetadata := &model.Metadata{ChannelId: "Fakkkee"}

	result := NewSummaryModal(testModel, testMetadata)
	actualJson, _ := json.Marshal(result)
	expectedJsonString := test.ReadFile(t, "view/summary_modal.json")
	actualJsonString := format.PrettyJson(t, string(actualJson))
	assert.EqualValues(t, expectedJsonString, actualJsonString)
}
