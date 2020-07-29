package bZapp

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/dctid/bZapp/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-lambda-go/events"
)

var openModalJson = `{
     "type": "modal",
     "submit": {
         "type": "plain_text",
         "text": "Submit",
         "emoji": true
     },
     "close": {
         "type": "plain_text",
         "text": "Cancel",
         "emoji": true
     },
     "title": {
         "type": "plain_text",
         "text": "bZapp",
         "emoji": true
     },
     "blocks": [
         {
             "type": "divider"
         },
         {
             "type": "context",
             "elements": [
                 {
                     "type": "mrkdwn",
                     "text": "*Today's Events*"
                 }
             ]
         },
         {
             "type": "divider"
         },
         {
             "type": "section",
             "text": {
                 "type": "mrkdwn",
                 "text": "9:15 Standup"
             },
             "accessory": {
                 "type": "button",
                 "text": {
                     "type": "plain_text",
                     "text": "Remove",
                     "emoji": true
                 },
                 "value": "remove_today_1"
             }, 
			 "block_id": "today_1"
         },
         {
             "type": "section",
             "text": {
                 "type": "mrkdwn",
                 "text": "11:30 IPM"
             },
             "accessory": {
                 "type": "button",
                 "text": {
                     "type": "plain_text",
                     "text": "Remove",
                     "emoji": true
                 },
                 "value": "remove_today_2"
             },
			 "block_id": "today_2"
         },
         {
             "type": "section",
             "text": {
                 "type": "mrkdwn",
                 "text": "3:15 Retro"
             },
             "accessory": {
                 "type": "button",
                 "text": {
                     "type": "plain_text",
                     "text": "Remove",
                     "emoji": true
                 },
                 "value": "remove_today_3"
             },
			 "block_id": "today_3"
         },
         {
             "type": "divider"
         },
         {
             "type": "context",
             "elements": [
                 {
                     "type": "mrkdwn",
                     "text": "*Tomorrow's Events*"
                 }
             ]
         },
         {
             "type": "divider"
         },
         {
             "type": "section",
             "text": {
                 "type": "mrkdwn",
                 "text": "9:15 Standup"
             },
             "accessory": {
                 "type": "button",
                 "text": {
                     "type": "plain_text",
                     "text": "Remove",
                     "emoji": true
                 },
                 "value": "click_me_123"
             }
         },
         {
             "type": "section",
             "text": {
                 "type": "mrkdwn",
                 "text": "1:30 User Interview"
             },
             "accessory": {
                 "type": "button",
                 "text": {
                     "type": "plain_text",
                     "text": "Remove",
                     "emoji": true
                 },
                 "value": "click_me_123"
             }
         },
         {
             "type": "section",
             "text": {
                 "type": "mrkdwn",
                 "text": "3:00 Synthesis"
             },
             "accessory": {
                 "type": "button",
                 "text": {
                     "type": "plain_text",
                     "text": "Remove",
                     "emoji": true
                 },
                 "value": "click_me_123"
             }
         },
         {
             "type": "divider"
         },
         {
             "type": "input",
             "element": {
                 "type": "plain_text_input",
                 "placeholder": {
                     "type": "plain_text",
                     "text": "Title"
                 }
             },
             "label": {
                 "type": "plain_text",
                 "text": "Add Event",
                 "emoji": false
             }
         },
         {
             "type": "actions",
             "elements": [
                 {
                     "type": "static_select",
                     "placeholder": {
                         "type": "plain_text",
                         "text": "Select hour",
                         "emoji": true
                     },
                     "options": [
                         {
                             "text": {
                                 "type": "plain_text",
                                 "text": "9 AM",
                                 "emoji": true
                             },
                             "value": "value-0"
                         },
                         {
                             "text": {
                                 "type": "plain_text",
                                 "text": "10 AM",
                                 "emoji": true
                             },
                             "value": "value-1"
                         },
                         {
                             "text": {
                                 "type": "plain_text",
                                 "text": "11 AM",
                                 "emoji": true
                             },
                             "value": "value-2"
                         }
                     ]
                 },
                 {
                     "type": "static_select",
                     "placeholder": {
                         "type": "plain_text",
                         "text": "Select minute",
                         "emoji": true
                     },
                     "options": [
                         {
                             "text": {
                                 "type": "plain_text",
                                 "text": "00",
                                 "emoji": true
                             },
                             "value": "value-0"
                         },
                         {
                             "text": {
                                 "type": "plain_text",
                                 "text": "01",
                                 "emoji": true
                             },
                             "value": "value-1"
                         },
                         {
                             "text": {
                                 "type": "plain_text",
                                 "text": "02",
                                 "emoji": true
                             },
                             "value": "value-2"
                         }
                     ]
                 },
                 {
                     "type": "datepicker",
                     "initial_date": "1990-04-28",
                     "placeholder": {
                         "type": "plain_text",
                         "text": "Select a date",
                         "emoji": true
                     }
                 },
                 {
                     "type": "button",
                     "text": {
                         "type": "plain_text",
                         "text": "Add",
                         "emoji": true
                     },
                     "value": "click_me_123"
                 }
             ]
         }
     ]
 }`

const (
	empty = ""
	tab   = "\t"
)


func PrettyJson(data string) (string, error) {
	expectedJson :=[]byte(strings.Join(strings.Fields(data), ""))
	var expectedMap map[string]interface{}
	err := json.Unmarshal(expectedJson, &expectedMap)
	if err != nil {
		return empty, err
	}
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err = encoder.Encode(expectedMap)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

func TestSlash(t *testing.T) {

	var urlCalled *url.URL = nil
	var bodyCalled string
	expectUrl, _ := url.Parse("http://localhost:8080/api/scores")
	Client = &mocks.MockClient{}
	expected := `{
		 "type": "modal",
	     "title": {
	         "type": "plain_text",
	         "text": "bZapp",
	         "emoji": true
	     },
		 "submit": {
	         "type": "plain_text",
	         "text": "Submit",
	         "emoji": true
	     },
		 "close": {
	         "type": "plain_text",
	         "text": "Cancel",
	         "emoji": true
	     },
 	    "blocks": [
         	{
            	 "type": "divider"
        	 },
	        {
	             "type": "context",
	             "elements": [
	                 {
	                     "type": "mrkdwn",
	                     "text": "*Today's Events*"
	                 }
	             ]
	         },
	        {
	             "type": "divider"
	        },
	         {
	             "type": "divider"
	         },
	         {
	             "type": "context",
	             "elements": [
	                 {
	                     "type": "mrkdwn",
	                     "text": "*Tomorrow's Events*"
	                 }
	             ]
	         },
	         {
	             "type": "divider"
	         }
		]
	}`

	prettyJsonExpected, err := PrettyJson(expected)
	assert.NoError(t, err)

	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		log.Printf("url %s ", req.URL)
		urlCalled = req.URL
		body, _ := ioutil.ReadAll(req.Body)
		bodyCalled = string(body)
		return &http.Response{
			StatusCode: 200,
		}, nil
	}

	result, err := Slash(context.Background(), events.APIGatewayProxyRequest{
		Body: `{"text": "none"}`,
	})
	assert.NoError(t, err)
	assert.EqualValues(t, expectUrl, urlCalled)

	prettyJsonActual, err := PrettyJson(bodyCalled)
	assert.NoError(t, err)
	assert.EqualValues(t, prettyJsonExpected, prettyJsonActual)
	assert.EqualValues(t, events.APIGatewayProxyResponse{StatusCode: 200}, result)
}
