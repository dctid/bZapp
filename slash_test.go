package bZapp

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/dctid/bZapp/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-lambda-go/events"
)

func TestSlash(t *testing.T) {

	var urlCalled *url.URL = nil
	expectUrl, _ := url.Parse("https://www.google.com")
	Client = &mocks.MockClient{}
	json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	// create a new reader with that JSON
	r2 := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(req *http.Request) (*http.Response, error) {
		log.Printf("url %s ", req.URL)
		urlCalled = req.URL
		return &http.Response{
			StatusCode: 200,
			Body:       r2,
		}, nil
	}

	_, err := Slash(context.Background(), events.APIGatewayProxyRequest{
		Body: `{"text": "none"}`,
	})
	assert.NoError(t, err)

	assert.EqualValues(t, urlCalled, expectUrl)
	// assert.EqualValues(t,
	// 	events.APIGatewayProxyResponse{
	// 		Body: "{\"text\": \"hi\"}",
	// 		Headers: map[string]string{
	// 			"Content-Type": "application/json",
	// 		},
	// 		StatusCode: 200,
	// 	},
	// 	r,
	// )
}
