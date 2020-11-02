package view

import (
	"github.com/slack-go/slack"
)

func NewErrorModal(errorMessage string) *slack.ModalViewRequest {
	return &slack.ModalViewRequest{
		Type:   slack.VTModal,
		Title:  slack.NewTextBlockObject(slack.PlainTextType, "bZapp", true, false),
		Close:  slack.NewTextBlockObject(slack.PlainTextType, "Close", true, false),
		Blocks: slack.Blocks{
			BlockSet: []slack.Block{
				slack.NewSectionBlock(
					slack.NewTextBlockObject(slack.PlainTextType, errorMessage, false, false),
					nil,
					nil,
				),
			},
		},
	}
}
