package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"net/url"
	"strings"
)

func UrlEncode(i interface{}) string {
	commandBytes, _ := json.Marshal(i)

	var f interface{}
	json.Unmarshal(commandBytes, &f)

	m := f.(map[string]interface{})
	b := new(bytes.Buffer)

	for key, value := range m {
		if len(fmt.Sprintf("%v", value)) > 0 {
			fmt.Fprintf(b, "%s=%v&", key, value)
		}
	}
	return url.PathEscape(strings.TrimSuffix(b.String(), "&"))
}

func MakePayload(callback slack.InteractionCallback) string {
	marshal, _ := json.Marshal(callback)
	return fmt.Sprintf("payload=%s", url.PathEscape(string(marshal)))
}