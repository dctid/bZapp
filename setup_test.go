package bZapp

import (
	"os"
	"testing"
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

