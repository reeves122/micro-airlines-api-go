package model

import "gorm.io/gorm"

type City struct {
	gorm.Model
	Name        string
	Country     string
	Cost        int64
	CityClass   string
	Population  int
	LayoverSize int
	Layovers    []Job
	Jobs        []Job
	JobsExpire  int64
	Latitude    float64
	Longitude   float64
}
