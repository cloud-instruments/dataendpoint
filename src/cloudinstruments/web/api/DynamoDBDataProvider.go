package main

import (
	"cloudinstruments/web/models"
	//"github.com/aws/aws-sdk-go"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBDataProvider struct {
	db            *dynamodb.DynamoDB
	tableName     string
	readCapacity  int64
	writeCapacity int64
}

func (d *DynamoDBDataProvider) GetBatteryTest(testId string) *models.BatteryCycle {
	return nil
}

func NewDynamoDBDataProvider() *DynamoDBDataProvider {
	return &DynamoDBDataProvider{}
}

func (d *DynamoDBDataProvider) PostBatteryTest(cycle *models.BatteryCycle) {

}

func (d *DynamoDBDataProvider) DeleteBatteryTest(testId, testName string) {

}
