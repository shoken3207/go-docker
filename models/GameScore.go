package models

import "gorm.io/gorm"

type GameScore struct {
	gorm.Model
	Order int `json:"order" gorm:"not null;uniqueIndex:game_team_order_unique"`
	Score int `json:"score" gorm:"not null"`
	GameId uint `json:"gameId" gorm:"column:game_id;not null;uniqueIndex:game_team_order_unique"`
	TeamId uint `json:"teamId" gorm:"column:team_id;not null;uniqueIndex:game_team_order_unique"`
	Game Game `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Team Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
