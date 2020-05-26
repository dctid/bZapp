package bZapp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator_ReturnsTrue_IfSecretNotSet(t *testing.T) {

	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`

	var reqHeaders = map[string][]string{"X-Slack-Signature": []string{"Anything"}}

	assert.True(t, VerifySigning([]byte(json), reqHeaders))

}

func TestValidator_ReturnsTrue_SecretSetAndRequestValid(t *testing.T) {

	os.Setenv("SLACK_SIGNING_SECRET", "VALID")
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`

	var reqHeaders = map[string][]string{"X-Slack-Signature": []string{"Anything"}}

	assert.True(t, VerifySigning([]byte(json), reqHeaders))

}

func TestValidator_ReturnsFalse_SecretSetAndRequestInvalid(t *testing.T) {

	os.Setenv("SLACK_SIGNING_SECRET", "INVALID")
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`

	var reqHeaders = map[string][]string{"X-Slack-Signature": []string{"Anything"}}

	assert.False(t, VerifySigning([]byte(json), reqHeaders))

}
