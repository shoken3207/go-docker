package models

import "gorm.io/gorm"

type League struct {
	gorm.Model
	Name string `json:"name" gorm:"size:200;not null;uniqueIndex:name_sport_unique"`
	SportId uint `json:"sportId" gorm:"not null;column:sport_id;uniqueIndex:name_sport_unique"`
	Teams []Team `gorm:"foreignKey:LeagueId"`
	Sport Sport `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}