package models

import "gorm.io/gorm"

type FavoriteTeam struct {
	gorm.Model
	UserId uint `json:"userId" gorm:"column:user_id;not null;uniqueIndex:user_team_unique"`
	TeamId uint `json:"teamId" gorm:"column:team_id;not null;uniqueIndex:user_team_unique"`
	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Team Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}