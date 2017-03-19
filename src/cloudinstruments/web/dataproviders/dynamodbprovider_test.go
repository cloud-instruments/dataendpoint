// here are mostly integration tests
// or tests to test AWS api
// to be removed
package dataproviders

import (
	"cloudinstruments/web/models"
	"fmt"
	"testing"
)

func TestPostProjectCycle(t *testing.T) {
	provider := NewDynamoDBDataProvider()
	cycle := models.BatteryCycle{
		ProjectName:  "Test",
		DeviceName:   "Phone",
		Cycle:        models.Rest,
		Duration:     123,
		StartVoltage: 0,
		EndVoltage:   1,
		VoltageDiff:  1,
		StartCurrent: 1,
		EndCurrent:   2,
		CurrentDiff:  2,
	}

	strings, err := provider.PostBatteryCycle(&cycle)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(strings)
}

func TestPostProject(t *testing.T) {
	provider := NewDynamoDBDataProvider()
	project := models.Project{
		ProjectName:    "Test",
		DeviceName:     "Phone",
		NumberOfCycles: 4,
		Tag:            "test",
		Comment:        "hello project",
		Created:        123,
		LastUpdated:    1234,
		FileName:       "firstproject.xlsx",
	}

	strings, err := provider.PostProject(&project)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(strings)
}
