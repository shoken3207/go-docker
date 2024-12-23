package models

import "gorm.io/gorm"

type ExpeditionImage struct {
	gorm.Model
	FileId       string     `json:"fileId" gorm:"not null;unique"`
	Image        string     `json:"image" gorm:"not null"`
	ExpeditionId uint       `json:"expeditionId" gorm:"column:expedition_id;not null"`
	Expedition   Expedition `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
