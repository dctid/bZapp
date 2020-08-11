package bZapp

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

func JsonHeaders() map[string]string {
	return	map[string]string{
		"content-Type": "application/json",
	}
}

// Post sends a post request to the URL with the body
func Post(url string, headers http.Header, body interface{} ) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers
	return Client.Do(request)
}

// Post sends a post request to the URL with the body
func Get(url string, headers http.Header) (*http.Response, error) {

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header = headers
	return Client.Do(request)
}
