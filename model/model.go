package model

import (
	"encoding/json"
)

type Model struct {
	Index int
	Events Events
	Goals Goals
	ChannelId string `json:"channel_id,omitempty"`
}

type Metadata struct {
	ChannelId string `json:"channel_id,omitempty"`
}

func (modelToConvert *Model) ConvertModelToJson() string {
	jsonBytes, _ := json.Marshal(modelToConvert)
	return string(jsonBytes)
}

func (modelToConvert *Model) ConvertMetadataToJson() string {
	jsonBytes, _ := json.Marshal(Metadata{ChannelId: modelToConvert.ChannelId})
	return string(jsonBytes)
}

func (modelToConvert *Model) ConvertToDbModel() *Model {
	return &Model{
		Index:     modelToConvert.Index,
		Events:    modelToConvert.Events.ConvertToDate(),
		Goals:     modelToConvert.Goals,
		ChannelId: modelToConvert.ChannelId,
	}
}
func (modelToConvert *Model) ConvertFromDbModel() *Model {
	modelToConvert.Events = modelToConvert.Events.ConvertFromDates()
	return modelToConvert
}

