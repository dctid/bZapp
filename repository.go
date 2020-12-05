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
		currentModel = model.Model{
			Index:     0,
			Events:    model.Events{},
			Goals:     model.Goals{},
			ChannelId: channel,
		}
		SaveModel(ctx, channel, &currentModel)
		if err != nil {
			return nil, err
		}
	} else {
		modelString := dbModel.Item["model"].S
		err = json.Unmarshal([]byte(*modelString), &currentModel)
		if err != nil {
			log.Printf("Couldn't parse metadata %s", err)
		} else {
			log.Printf("Metadata: %v", currentModel)
		}

	}

	return currentModel.ConvertFromDbModel(), nil
}

func SaveModel(ctx aws.Context, channel string, currentModel *model.Model) error {
	log.Printf("CurrentModel: %v", currentModel)
	log.Printf("ConvertedModel: %v", currentModel.ConvertToDbModel())
	modelBytes, err := json.Marshal(currentModel.ConvertToDbModel())
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
				S: aws.String(string(modelBytes)),
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
