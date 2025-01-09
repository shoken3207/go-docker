package models

import (
	"time"
)

type Payment struct {
	BaseModel
	Title string `json:"title" gorm:"size:100;not null"`
	Date time.Time `json:"date" gorm:"not null"`
	Cost int `json:"cost" gorm:"not null"`
	ExpeditionId uint `json:"expeditionId" gorm:"column:expedition_id;not null"`
	Expedition Expedition `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}