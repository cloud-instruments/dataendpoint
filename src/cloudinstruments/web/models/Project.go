package models

import (
	"time"
)

type Project struct {
	ProjectName    string
	DeviceName     string
	NumberOfCycles int
	Tag            string
	Comment        string
	Created        time.Duration
	LastUpdated    time.Duration
	FileName       string
}
