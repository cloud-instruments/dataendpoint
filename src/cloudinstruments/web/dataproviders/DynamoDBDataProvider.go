package dataproviders

import (
	"cloudinstruments/web/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
				N: aws.String(project.Created),
			},
			"LastUpdated": &dynamodb.AttributeValue{
				N: aws.String(project.LastUpdated),
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
			"CycleNumber": &dynamodb.AttributeValue{
				N: aws.String(strconv.Itoa(cycle.CycleNumber)),
			},
			"DeviceName": &dynamodb.AttributeValue{
				S: aws.String(cycle.DeviceName),
			},
			"CycleType": &dynamodb.AttributeValue{
				N: aws.String(strconv.Itoa(int(cycle.Cycle))),
			},
			"Duration": &dynamodb.AttributeValue{
				N: aws.String(strconv.Itoa(cycle.Duration)),
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

func (d *DynamoDBDataProvider) GetProjectsByDeviceName(deviceName string) ([]models.Project, error) {
	cnfg := aws.Config{
		Region: aws.String("us-west-2"),
	}

	sn := session.New(&cnfg)
	db := dynamodb.New(sn)
	queryInput := &dynamodb.ScanInput{
		TableName: aws.String("Projects"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name": &dynamodb.AttributeValue{
				S: aws.String(deviceName),
			},
		},
		FilterExpression: aws.String("DeviceName = :name"),
	}

	result, errRequest := db.Scan(queryInput)
	if errRequest != nil {
		return nil, errRequest
	}

	projects := []models.Project{}
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (d *DynamoDBDataProvider) GetProjectCyclesByProjectName(projectName string) ([]models.BatteryCycle, error) {
	cnfg := aws.Config{
		Region: aws.String("us-west-2"),
	}

	sn := session.New(&cnfg)
	db := dynamodb.New(sn)
	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String("ProjectCycles"),
		KeyConditionExpression: aws.String("ProjectName = :name"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":name": &dynamodb.AttributeValue{
				S: aws.String(projectName),
			},
		},
	}

	result, errRequest := db.Query(queryInput)
	if errRequest != nil {
		return nil, errRequest
	}

	cycles := []models.BatteryCycle{}
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &cycles); err != nil {
		return nil, err
	}

	return cycles, nil
}

func (d *DynamoDBDataProvider) DeleteProject(projectName string) (*dynamodb.DeleteItemOutput, error) {
	cnfg := aws.Config{
		Region: aws.String("us-west-2"),
	}

	sn := session.New(&cnfg)
	db := dynamodb.New(sn)
	deleteInput := &dynamodb.DeleteItemInput{
		TableName: aws.String("Projects"),
		Key: map[string]*dynamodb.AttributeValue{
			"ProjectName": &dynamodb.AttributeValue{
				S: aws.String(projectName),
			},
		},
	}

	return db.DeleteItem(deleteInput)
}

func (d *DynamoDBDataProvider) DeleteProjectCycles(projectName string) (*dynamodb.DeleteItemOutput, error) {
	cnfg := aws.Config{
		Region: aws.String("us-west-2"),
	}

	sn := session.New(&cnfg)
	db := dynamodb.New(sn)
	deleteInput := &dynamodb.DeleteItemInput{
		TableName: aws.String("ProjectCycles"),
		Key: map[string]*dynamodb.AttributeValue{
			"ProjectName": &dynamodb.AttributeValue{
				S: aws.String(projectName),
			},
		},
	}

	return db.DeleteItem(deleteInput)
}
