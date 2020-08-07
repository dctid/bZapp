package test

import (
	"encoding/json"
	"strings"
)

const (
	empty = ""
	tab   = "\t"
)

func PrettyJson(data string) (string, error) {
	expectedJson := []byte(strings.Join(strings.Fields(data), ""))
	var expectedMap map[string]interface{}
	err := json.Unmarshal(expectedJson, &expectedMap)
	if err != nil {
		return empty, err
	}

	indent, err := json.MarshalIndent(expectedMap, empty, tab)
	if err != nil {
		return empty, err
	}
	return string(indent), nil
}

