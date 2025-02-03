package models

type GameScore struct {
	BaseModel
	GameId     uint `json:"gameId" gorm:"column:game_id;not null"`
	Team1Score int  `json:"team1Score" gorm:"column:team1_score;not null"`
	Team2Score int  `json:"team2Score" gorm:"column:team2_score;not null"`
	Order      int  `json:"order" gorm:"column:order;not null"`
	Game       Game `gorm:"foreignKey:GameId;references:ID;constraint:OnDelete:SET NULL;OnUpdate:CASCADE"`
}
