package bZapp

import (
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

// var (
// 	signingSecret string
// 	isSet         bool
// )

// func init() {
// 	signingSecret, isSet = os.LookupEnv("SLACK_SIGNING_SECRET")
// }

func VerifySigning(body []byte, header http.Header) bool {
	var signingSecret, isSet = os.LookupEnv("SLACK_SIGNING_SECRET")
	if isSet {
		sv, err := slack.NewSecretsVerifier(header, signingSecret)
		if err != nil {
			log.Printf("[ERROR] Fail to verify SigningSecret: %v", err)
			return false
		}
		sv.Write(body)
		if err := sv.Ensure(); err != nil {
			log.Printf("[ERROR] Fail to verify SigningSecret: %v", err)
			return false
		}

	}
	return true
}
