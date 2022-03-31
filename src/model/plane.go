package model

import "gorm.io/gorm"

type Plane struct {
	gorm.Model
	Name              string
	Cost              int64
	Speed             int
	Weight            int
	CapacityType      string
	Capacity          int
	FlightRange       int
	SizeClass         string
	LoadedJobs        []Job
	CurrentCityID     string
	DestinationCityID string
}
