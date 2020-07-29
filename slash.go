package bZapp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/slack-go/slack"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Title struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji"`
}

type Modal struct {
	Type           string                `json:"type"`
	TriggerID      string                `json:"trigger_id,omitempty"`
	CallbackID     string                `json:"callback_id,omitempty"` // Required
	State          string                `json:"state,omitempty"`       // Optional
	Title          *Title                `json:"title"`
	SubmitLabel    string                `json:"submit_label,omitempty"`
	NotifyOnCancel bool                  `json:"notify_on_cancel,omitempty"`
	Elements       []slack.DialogElement `json:"elements,omitempty"`
}

func NewModal(title *Title) *Modal {
	return &Modal{
		Type:  "modal",
		Title: title,
	}
}

func Map(vs []int, f func(int) *slack.OptionBlockObject) []*slack.OptionBlockObject {
	vsm := make([]*slack.OptionBlockObject, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func HourOption(num int) *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(fmt.Sprintf("hour-%d", num), slack.NewTextBlockObject("plain_text", fmt.Sprintf("%d %s", num, func() string {
		if num < 9 || num == 12 {
			return "PM"
		} else {
			return "AM"
		}
	}()), true, false))
}

func MinOption(num int) *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(fmt.Sprintf("min-%d", num), slack.NewTextBlockObject("plain_text", fmt.Sprintf(func() string {
		if num < 10 {
			return "0%d"
		} else {
			return "%d"
		}
	}(), num), true, false))
}

func Slash(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//var headers = map[string]string{
	//	"Content-Type": "application/json",
	//}
	//var block = slack.NewTextBlockObject("plain_text", "HIII", false, false)

	titleText := slack.NewTextBlockObject("plain_text", "bZapp", true, false)
	submitText := slack.NewTextBlockObject("plain_text", "Submit", true, false)
	closeText := slack.NewTextBlockObject("plain_text", "Cancel", true, false)
	todayHeader := slack.NewContextBlock("", slack.NewTextBlockObject("mrkdwn", "*Today's Events*", false, false))
	tomorrowHeader := slack.NewContextBlock("", slack.NewTextBlockObject("mrkdwn", "*Tomorrow's Events*", false, false))
	addEventsElement := slack.NewPlainTextInputBlockElement(slack.NewTextBlockObject("plain_text", "Title", false, false), "add_event")
	addEvents := slack.NewInputBlock("", slack.NewTextBlockObject("plain_text", "Add Event", false, false), addEventsElement)

	hours := []int{9, 10, 11, 12, 1, 2, 3, 4}
	hourOptions := Map(hours, HourOption)

	mins := []int{0, 15, 30, 45}
	minOptions := Map(mins, MinOption)

	datepicker := slack.DatePickerBlockElement{
		Type:        "datepicker",
		ActionID:    "datepicker",
		Placeholder:  slack.NewTextBlockObject("plain_text", "Select a date", true, false),
		InitialDate: "1990-04-28",
	}

	actions := slack.NewActionBlock(
		"",
		slack.NewOptionsSelectBlockElement("static_select", slack.NewTextBlockObject("plain_text", "Select hour", true, false), "", hourOptions...),
		slack.NewOptionsSelectBlockElement("static_select", slack.NewTextBlockObject("plain_text", "Select minutes", true, false), "", minOptions...),
		datepicker,
		slack.NewButtonBlockElement("", "click_me_123", slack.NewTextBlockObject("plain_text", "Add", true, false)),
	)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			slack.NewDividerBlock(),
			todayHeader,
			slack.NewDividerBlock(),
			slack.NewDividerBlock(),
			tomorrowHeader,
			slack.NewDividerBlock(),
			slack.NewDividerBlock(),
			addEvents,
			actions,
		},
	}

	var modalRequest slack.ModalViewRequest
	modalRequest.Type = "modal"
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks

	jsonBytes, err := json.Marshal(modalRequest)
	log.Printf("json %s", jsonBytes)

	postHeaders := http.Header{"Content-Type": {"application/json"}}

	_, err = Post("http://localhost:8080/api/scores", postHeaders, modalRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	//defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//println(string(body))

	return events.APIGatewayProxyResponse{

		StatusCode: 200,
	}, nil
}
