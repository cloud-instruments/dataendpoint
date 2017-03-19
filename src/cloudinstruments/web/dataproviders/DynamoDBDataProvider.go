package dataproviders

import (
	"cloudinstruments/web/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strconv"
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

func (d *DynamoDBDataProvider) PostProject(project *models.Project) (*dynamodb.PutItemOutput, error) {
	cnfg := aws.Config{
		Region: aws.String("us-west-2"),
	}

	sn := session.New(&cnfg)
	db := dynamodb.New(sn)
	input := dynamodb.PutItemInput{
		TableName: aws.String("Projects"),
		Item: map[string]*dynamodb.AttributeValue{
			"ProjectName": &dynamodb.AttributeValue{
				S: aws.String(project.ProjectName),
			},
			"DeviceName": &dynamodb.AttributeValue{
				S: aws.String(project.DeviceName),
			},
			"NumberOfCycles": &dynamodb.AttributeValue{
				N: aws.String(strconv.Itoa(project.NumberOfCycles)),
			},
			"Tag": &dynamodb.AttributeValue{
				S: aws.String(project.Tag),
			},
			"Comment": &dynamodb.AttributeValue{
				S: aws.String(project.Comment),
			},
			"Created": &dynamodb.AttributeValue{
				N: aws.String(strconv.FormatFloat(project.Created.Seconds(), 'E', -1, 64)),
			},
			"LastUpdated": &dynamodb.AttributeValue{
				N: aws.String(strconv.FormatFloat(project.LastUpdated.Seconds(), 'E', -1, 64)),
			},
			"FileName": &dynamodb.AttributeValue{
				S: aws.String(project.FileName),
			},
		},
	}

	resp, err := db.PutItem(&input)
	return resp, err
}

func (d *DynamoDBDataProvider) PostBatteryCycle(cycle *models.BatteryCycle) (*dynamodb.PutItemOutput, error) {
	cnfg := aws.Config{
		Region: aws.String("us-west-2"),
	}

	sn := session.New(&cnfg)
	db := dynamodb.New(sn)
	input := dynamodb.PutItemInput{
		TableName: aws.String("ProjectCycles"),
		Item: map[string]*dynamodb.AttributeValue{
			"ProjectName": &dynamodb.AttributeValue{
				S: aws.String(cycle.ProjectName),
			},
			"DeviceName": &dynamodb.AttributeValue{
				S: aws.String(cycle.DeviceName),
			},
			"CycleType": &dynamodb.AttributeValue{
				N: aws.String(strconv.Itoa(int(cycle.Cycle))),
			},
			"Duration": &dynamodb.AttributeValue{
				N: aws.String(strconv.FormatFloat(cycle.Duration.Seconds(), 'E', -1, 64)),
			},
			"StartVoltage": &dynamodb.AttributeValue{
				N: aws.String(strconv.FormatFloat(cycle.StartVoltage, 'E', -1, 64)),
			},
			"EndVoltage": &dynamodb.AttributeValue{
				N: aws.String(strconv.FormatFloat(cycle.EndVoltage, 'E', -1, 64)),
			},
			"VoltageDiff": &dynamodb.AttributeValue{
				N: aws.String(strconv.FormatFloat(cycle.VoltageDiff, 'E', -1, 64)),
			},
			"StartCurrent": &dynamodb.AttributeValue{
				N: aws.String(strconv.FormatFloat(cycle.StartCurrent, 'E', -1, 64)),
			},
			"EndCurrent": &dynamodb.AttributeValue{
				N: aws.String(strconv.FormatFloat(cycle.EndCurrent, 'E', -1, 64)),
			},
			"CurrentDiff": &dynamodb.AttributeValue{
				N: aws.String(strconv.FormatFloat(cycle.CurrentDiff, 'E', -1, 64)),
			},
		},
	}

	resp, err := db.PutItem(&input)
	return resp, err
}

func (d *DynamoDBDataProvider) DeleteBatteryTest(projectName string) {

}
