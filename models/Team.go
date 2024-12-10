package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name string `json:"name" gorm:"size:200;not null;uniqueIndex:name_league_unique"`
	StadiumId uint `json:"stadiumId" gorm:"column:stadium_id"`
	LeagueId uint `json:"leagueId" gorm:"column:league_id;not null;uniqueIndex:name_league_unique"`
	SportId uint `json:"sportId" gorm:"column:sport_id;not null"`
	FavoriteTeams []FavoriteTeam `gorm:"foreignKey:TeamId"`
	GamesAsTeam1 []Game `gorm:"foreignKey:Team1Id;references:ID;constraint:OnDelete:SET NULL;OnUpdate:CASCADE"`
	GamesAsTeam2 []Game `gorm:"foreignKey:Team2Id;references:ID;constraint:OnDelete:SET NULL;OnUpdate:CASCADE"`
	Stadium Stadium `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	League League `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Sport Sport `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}