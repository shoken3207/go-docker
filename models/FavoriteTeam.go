package models

type FavoriteTeam struct {
	BaseModel
	UserId uint `json:"userId" gorm:"column:user_id;not null;uniqueIndex:user_team_unique"`
	TeamId uint `json:"teamId" gorm:"column:team_id;not null;uniqueIndex:user_team_unique"`
	User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Team Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}