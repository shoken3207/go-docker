package models

import (
	"time"
)

type Game struct {
	BaseModel
	Date         time.Time   `json:"date" gorm:"not null"`
	Comment      string      `json:"comment" gorm:"type:text;not null"`
	ExpeditionId uint        `json:"expeditionId" gorm:"column:expedition_id;not null"`
	Team1Id      uint        `json:"team1Id" gorm:"column:team1_id;not null"`
	Team2Id      uint        `json:"team2Id" gorm:"column:team2_id;not null"`
	Expedition   Expedition  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Team1        Team        `gorm:"foreignKey:Team1Id;references:ID;constraint:OnDelete:SET NULL;OnUpdate:CASCADE"`
	Team2        Team        `gorm:"foreignKey:Team2Id;references:ID;constraint:OnDelete:SET NULL;OnUpdate:CASCADE"`
	GameScores   []GameScore `gorm:"foreignKey:GameId;constraint:OnDelete:CASCADE"`
}
