package models

import (
	"time"
)

type CycleType int

const (
	Rest      CycleType = iota
	Discharge           = iota
	ChargeCC            = iota
	ChargeCV            = iota
)

type BatteryCycle struct {
	Id           string
	ProjectName  string
	DeviceName   string
	Cycle        CycleType
	Duration     time.Duration
	StartVoltage float32
	EndVoltage   float32
	VoltageDiff  float32
	StartCurrent float32
	EndCurrent   float32
	CurrentDiff  float32
}
