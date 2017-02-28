package main

import (
	"cloudinstruments/web/models"
)

type BatteryDataProvider interface {
	GetBatteryTest(testId string) *models.BatteryCycle
	PostBatteryTest(*models.BatteryCycle)
	DeleteBatteryTest(testId, testName string)
}
