package models

type ExpeditionLike struct {
	BaseModel
	UserId uint `json:"userId" gorm:"column:user_id;not null;uniqueIndex:user_expedition_unique"`
	ExpeditionId uint `json:"expeditionId" gorm:"column:expedition_id;not null;uniqueIndex:user_expedition_unique"`
	User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Expedition Expedition `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}