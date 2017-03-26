package models

type CycleType int

const (
	Rest      CycleType = iota
	Discharge           = iota
	ChargeCC            = iota
	ChargeCV            = iota
)

type BatteryCycle struct {
	ProjectName  string
	CycleNumber  int
	DeviceName   string
	Cycle        int
	Duration     int
	StartVoltage float64
	EndVoltage   float64
	VoltageDiff  float64
	StartCurrent float64
	EndCurrent   float64
	CurrentDiff  float64
}
