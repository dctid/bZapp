package mocks

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type MockDynamoDB struct {
	DeleteItemOutput *dynamodb.DeleteItemOutput
	GetItemOutput    *dynamodb.GetItemOutput
	PutItemOutput    *dynamodb.PutItemOutput
}

var MockDynamoDbCalls = map[string][]interface{}{
	"DeleteItemWithContext" : {},
	"GetItemWithContext" : {},
	"PutItemWithContext" : {},

}

func (m *MockDynamoDB) DeleteItemWithContext(ctx aws.Context, input *dynamodb.DeleteItemInput, opts ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	MockDynamoDbCalls["DeleteItemWithContext"] = append(MockDynamoDbCalls["DeleteItemWithContext"], input)
	return m.DeleteItemOutput, nil
}

func (m *MockDynamoDB) GetItemWithContext(ctx aws.Context, input *dynamodb.GetItemInput, opts ...request.Option) (*dynamodb.GetItemOutput, error) {
	MockDynamoDbCalls["GetItemWithContext"] = append(MockDynamoDbCalls["GetItemWithContext"], input)
	return m.GetItemOutput, nil
}

func (m *MockDynamoDB) PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error) {
	MockDynamoDbCalls["PutItemWithContext"] = append(MockDynamoDbCalls["PutItemWithContext"], input)
	return m.PutItemOutput, nil
}


func ResetMockCalls()  {
	MockDynamoDbCalls = map[string][]interface{}{
		"DeleteItemWithContext" : {},
		"GetItemWithContext" : {},
		"PutItemWithContext" : {},
	}
}
