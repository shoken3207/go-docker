package models

import (
	"time"
)

type Expedition struct {
	BaseModel
	UserId            uint              `json:"userId" gorm:"column:user_id;not null"`
	SportId           uint              `json:"sportId" gorm:"not null;column:sport_id"`
	IsPublic          bool              `json:"isPublic" gorm:"column:is_public;not null"`
	Title             string            `json:"title" gorm:"size:200;not null"`
	StartDate         time.Time         `json:"startDate" gorm:"column:start_date;not null"`
	EndDate           time.Time         `json:"endDate" gorm:"column:end_date;not null"`
	StadiumId         uint              `json:"stadiumId" gorm:"column:stadium_id;not null"`
	Memo              string            `json:"memo" gorm:"type:text;not null"`
	VisitedFacilities []VisitedFacility `gorm:"foreignKey:ExpeditionId;constraint:OnDelete:CASCADE"`
	Payments          []Payment         `gorm:"foreignKey:ExpeditionId;constraint:OnDelete:CASCADE"`
	ExpeditionImages  []ExpeditionImage `gorm:"foreignKey:ExpeditionId;constraint:OnDelete:CASCADE"`
	ExpeditionLikes   []ExpeditionLike  `gorm:"foreignKey:ExpeditionId;constraint:OnDelete:CASCADE"`
	Games             []Game            `gorm:"foreignKey:ExpeditionId;constraint:OnDelete:CASCADE"`
	Sport             Sport             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Stadium           Stadium           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User              User              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
