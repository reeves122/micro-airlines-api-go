package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	OriginCityID      string
	DestinationCityID string
	Revenue           int64
	JobType           string
}
