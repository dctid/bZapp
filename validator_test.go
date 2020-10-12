package bZapp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	signingSecret string
	signingSet    bool
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func shutdown() {
	signingSecret, signingSet = os.LookupEnv("SLACK_SIGNING_SECRET")
}

func setup() {
	if signingSet {
		os.Setenv("SLACK_SIGNING_SECRET", signingSecret)
	}
}

func TestValidator_ReturnsTrue_IfSecretNotSet(t *testing.T) {

	os.Unsetenv("SLACK_SIGNING_SECRET")

	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`

	var reqHeaders = map[string]string{"X-Slack-Signature": "Anything"}

	assert.True(t, verifySigning(json, reqHeaders))

}

//func TestValidator_ReturnsTrue_SecretSetAndRequestValid(t *testing.T) {
//   t.Skip("Not sure if the api can do this")
//	os.Setenv("SLACK_SIGNING_SECRET", "")
//	json := `{"token":'8KTh0sVRkeZozlTxrBRqk1NO',"team_id":'T7NS02BFB',"team_domain":'ford-community',"channel_id":'D7P4LC5G9',"channel_name":'directmessage',"user_id":'U7QNBA36K',"user_name":'cdorman1',"command":'/bzapp',"text":'',"response_url":'https://hooks.slack.com/commands/T7NS02BFB/1158151340372/7OcwUt6cv6vpkSbhlykaxTHS',"trigger_id":'1151971965202.260884079521.7e40edbf839d200408a81239cbeacf4d'}`
//
//	var reqHeaders = map[string][]string{"X-Slack-Signature": []string{"v0=64550a5d5c969ce2447a9df41d7fbe830fe5e3e7c352681efd7cb0fc31e0e9cd"},
//		"X-Slack-Request-Timestamp": []string{"1590708241"}}
//
//	assert.True(t, verifySigning([]byte(json), reqHeaders))
//
//}

func TestValidator_ReturnsFalse_SecretSetAndRequestInvalid(t *testing.T) {

	os.Setenv("SLACK_SIGNING_SECRET", "INVALID")
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`

	var reqHeaders = map[string]string{"X-Slack-Signature": "Anything"}

	assert.False(t, verifySigning(json, reqHeaders))

}
