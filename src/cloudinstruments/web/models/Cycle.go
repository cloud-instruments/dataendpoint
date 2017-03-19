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
	ProjectName  string
	DeviceName   string
	Cycle        CycleType
	Duration     time.Duration
	StartVoltage float64
	EndVoltage   float64
	VoltageDiff  float64
	StartCurrent float64
	EndCurrent   float64
	CurrentDiff  float64
}
