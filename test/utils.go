package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dctid/bZapp/format"
	"github.com/dctid/bZapp/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"testing"
)

func ReadFile(t *testing.T, name string) string {
	path, err := filepath.Abs("./")
	assert.NoError(t, err)
	log.Printf("Path %s", path)
	dat, err := ioutil.ReadFile(fmt.Sprintf("%s/test/data/%s", path, name))
	assert.NoError(t, err)

	prettyJson := format.PrettyJson(t, string(dat))
	assert.NoError(t, err)
	return prettyJson
}

func ParseUrl(urlString string) *url.URL {
	result, _ := url.Parse(urlString)
	return result
}

func SuccessResponse(response string) *http.Response {
	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(response))),
		StatusCode: 200,
	}
}

func GetItemOutput(id string, modelToReturn *model.Model) *dynamodb.GetItemOutput {
	modelBytes, _ := json.Marshal(modelToReturn)
	return &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
			"model": {
				S: aws.String(string(modelBytes)),
			},
		},
	}
}

func PutItemInput(id string, modelToSave *model.Model) *dynamodb.PutItemInput {
	modelBytes, _ := json.Marshal(modelToSave)
	return &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
			"model": {
				S: aws.String(string(modelBytes)),
			},
		},
		TableName: aws.String("bZappTable"),
	}
}

func GetItemInput(id string) *dynamodb.GetItemInput {
	return &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String("bZappTable"),
	}
}

