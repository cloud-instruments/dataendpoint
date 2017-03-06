package dataproviders

import (
	"cloudinstruments/web/models"
)

type BatteryDataProvider interface {
	GetBatteryTest(projectName string) *models.BatteryCycle
	PostBatteryTest(*models.Project)
	DeleteBatteryTest(projectName string)
}
