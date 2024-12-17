package models

import (
	"time"

	"gorm.io/gorm"
)

type Expedition struct {
	gorm.Model
	SportId           uint              `json:"sportId" gorm:"not null;column:sport_id"`
	IsPublic          bool              `json:"isPublic" gorm:"column:is_public;not null"`
	Title             string            `json:"title" gorm:"size:200;not null"`
	StartDate         time.Time         `json:"startDate" gorm:"column:start_date;not null"`
	EndDate           time.Time         `json:"endDate" gorm:"column:end_date;not null"`
	Memo              string            `json:"memo" gorm:"type:text;not null"`
	VisitedFacilities []VisitedFacility `gorm:"foreignKey:ExpeditionId"`
	Payments          []Payment         `gorm:"foreignKey:ExpeditionId"`
	ExpeditionImages  []ExpeditionImage `gorm:"foreignKey:ExpeditionId"`
	ExpeditionLikes   []ExpeditionLike  `gorm:"foreignKey:ExpeditionId"`
	Games             []Game            `gorm:"foreignKey:ExpeditionId"`
	Sport             Sport             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
