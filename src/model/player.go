package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Username string `gorm:"unique"`
	Balance  int64
	Cities   []City
	Planes   []Plane
}
