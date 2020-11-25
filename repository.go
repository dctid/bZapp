package bZapp

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dctid/bZapp/model"
	"log"
	"os"
	"time"
)

func GetModelFromDb(ctx aws.Context, channel string) (*model.Model, error) {
	table := os.Getenv("DYNAMODB_TABLE_NAME")

	endpoint, isSet := os.LookupEnv("DYNAMODB_ENDPOINT")
	log.Printf("env: %s, %v", endpoint, isSet)

	log.Printf("Table: %s", table)
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(channel),
			},
		},
		TableName: aws.String("bZappTable"),
	}

	dbModel, err := DynamoDB.GetItemWithContext(ctx, input)
	if err != nil {
		log.Printf("Couldn't get model %s", err)
		return nil, err
	}
	log.Printf("model %+v", dbModel)
	var currentModel model.Model
	if len(dbModel.Item) == 0 {
		currentModel = model.Model{ChannelId: channel}
		SaveModel(ctx, channel, &currentModel)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal(dbModel.Item["model"].B, &currentModel)
		if err != nil {
			log.Printf("Couldn't parse metadata %s", err)
		} else {
			log.Printf("Metadata: %v", currentModel)
		}

	}

	return &currentModel, nil
}

func SaveModel(ctx aws.Context, channel string, currentModel *model.Model) error {
	modelBytes, err := json.Marshal(currentModel)
	if err != nil {
		log.Printf("Couldn't convert model %s", err)
		return err
	}
	contextWithTimeout, cancelFunc := context.WithTimeout(ctx, time.Second)
	defer cancelFunc()
	withContext, err := DynamoDB.PutItemWithContext(contextWithTimeout, &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(channel),
			},
			"model": {
				B: modelBytes,
			},
		},
		TableName: aws.String("bZappTable"),
	},
	)
	if err != nil {
		log.Printf("Couldn't save model %s", err)
	}
	log.Printf("Put result: %v", withContext)
	return nil
}
