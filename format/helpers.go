package format

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	empty = ""
	tab   = "\t"
)


func PrettyJson(t *testing.T, data string) string {
	expectedJson := []byte(strings.Join(strings.Fields(data), ""))
	var expectedMap map[string]interface{}
	err := json.Unmarshal(expectedJson, &expectedMap)
	assert.NoError(t, err)

	indent, err := json.MarshalIndent(expectedMap, empty, tab)
	assert.NoError(t, err)
	return string(indent)
}

func PrettyJsonNoError(data string) (string) {
	expectedJson := []byte(strings.Join(strings.Fields(data), ""))
	var expectedMap map[string]interface{}
	err := json.Unmarshal(expectedJson, &expectedMap)
	if err != nil {
		return empty
	}

	indent, err := json.MarshalIndent(expectedMap, empty, tab)
	if err != nil {
		return empty
	}
	return string(indent)
}

