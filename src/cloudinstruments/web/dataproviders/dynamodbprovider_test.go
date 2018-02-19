Copyright (c) <2018> <University of Washington>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.


// here are mostly integration tests
// or tests to test AWS api
// to be removed
package dataproviders

import (
	"cloudinstruments/web/models"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestPostProjectCycle(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	provider := NewDynamoDBDataProvider()
	cycle := models.BatteryCycle{
		ProjectName:  "Test",
		DeviceName:   "Phone",
		CycleNumber:  r.Intn(20),
		Cycle:        2,
		Duration:     123,
		StartVoltage: 0,
		EndVoltage:   1,
		VoltageDiff:  1,
		StartCurrent: 1,
		EndCurrent:   2,
		CurrentDiff:  2,
	}

	_, err := provider.PostBatteryCycle(&cycle)
	if err != nil {
		t.Fail()
	}
}

func TestPostProject(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	provider := NewDynamoDBDataProvider()
	project := models.Project{
		ProjectName:    "Test" + strconv.Itoa(r.Intn(20)),
		DeviceName:     "Phone",
		NumberOfCycles: 4,
		Tag:            "test",
		Comment:        "hello project",
		Created:        "123",
		LastUpdated:    "1234",
		FileName:       "firstproject.xlsx",
	}

	_, err := provider.PostProject(&project)
	if err != nil {
		t.Fail()
	}
}

func TestGetProjectsByDevice(t *testing.T) {
	provider := NewDynamoDBDataProvider()
	deviceName := "Phone"

	projects, err := provider.GetProjectsByDeviceName(deviceName)
	if err != nil || len(projects) == 0 {
		t.Fail()
	}
}

func TestGetProjectCyclesByProjectName(t *testing.T) {
	provider := NewDynamoDBDataProvider()
	projectName := "Test"

	cycles, err := provider.GetProjectCyclesByProjectName(projectName)
	if err != nil || len(cycles) == 0 {
		t.Fail()
	}
}

func TestDeleteProjectByProjectName(t *testing.T) {
	provider := NewDynamoDBDataProvider()
	projectName := "Test47"

	_, err := provider.DeleteProject(projectName)
	if err != nil {
		t.Fail()
	}
}

func TestDeleteProjectCyclesByProjectName(t *testing.T) {
	provider := NewDynamoDBDataProvider()
	projectName := "Test"

	_, err := provider.DeleteProjectCycles(projectName)
	if err != nil {
		t.Fail()
	}
}
