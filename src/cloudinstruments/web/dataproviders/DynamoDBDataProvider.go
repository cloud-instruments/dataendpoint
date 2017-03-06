package dataproviders

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

func (d *DynamoDBDataProvider) GetBatteryTest(projectName string) *models.BatteryCycle {
	return nil
}

func NewDynamoDBDataProvider() *DynamoDBDataProvider {
	return &DynamoDBDataProvider{}
}

func (d *DynamoDBDataProvider) PostBatteryTest(cycle *models.Project) {

}

func (d *DynamoDBDataProvider) DeleteBatteryTest(projectName string) {

}
