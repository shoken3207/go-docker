package models

import "gorm.io/gorm"

type VisitedFacility struct {
	gorm.Model
	Name         string     `json:"name" gorm:"size:200;not null"`
	Address      string     `json:"address" gorm:"size:200;not null"`
	Latitude     float64    `json:"latitude" gorm:"not null"`
	Longitude    float64    `json:"longitude" gorm:"not null"`
	Icon         string     `json:"icon" gorm:"not null"`
	Color        string     `json:"color" gorm:"not null"`
	ExpeditionId uint       `json:"expeditionId" gorm:"column:expedition_id;not null"`
	Expedition   Expedition `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
