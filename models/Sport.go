package models

import "gorm.io/gorm"

type Sport struct {
	gorm.Model
	Name string `json:"name" gorm:"size:200;unique;not null"`
	Leagues []League `gorm:"foreignKey:SportId"`
	Teams []Team `gorm:"foreignKey:SportId"`
	Expeditions []Expedition `gorm:"foreignKey:SportId"`
}