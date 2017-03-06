package models

import ()

type Project struct {
	Id             string
	Project        string
	DeviceName     string
	NumberOfCycles int
	Tag            string
	Comment        string
	Created        string
	LastUpdated    string
	FileName       string
	Cycles         []BatteryCycle
}
