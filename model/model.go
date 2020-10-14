package model

import "encoding/json"

type Model struct {
	Index int
	Events Events
	Goals Goals
}

func (modelToConvert *Model) ConvertModelToJson() string {
	jsonBytes, _ := json.Marshal(modelToConvert)
	return string(jsonBytes)
}