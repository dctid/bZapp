package view

import (
	"github.com/slack-go/slack"
	"log"
)

func DailySummaryMessage(payload InteractionPayload) (string, slack.Message) {
	url := payload.ResponseUrls[0].ResponseUrl
	log.Printf("Response Urls: %s", url)

	todaysEvents, tomorrowsEvents := ExtractEvents(payload.View.Blocks.BlockSet)
	todaysSectionBlocks, tomorrowsSectionBlocks := ConvertToEventsWithoutRemoveButton(todaysEvents, tomorrowsEvents)
	eventBlocks := buildEventsBlock(todaysSectionBlocks, tomorrowsSectionBlocks)

	message := slack.NewBlockMessage(eventBlocks...)
	message.Text = "bZapp - Today's Standup Summary"
	message.ResponseType = slack.ResponseTypeInChannel
	return url, message
}

