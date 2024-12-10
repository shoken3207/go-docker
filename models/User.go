package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"size:100;not null;unique"`
	PassHash string `json:"passHash" gorm:"not null;column:pass_hash"`
	Name      string `json:"name" gorm:"size:100;not null"`
	Description string `json:"description" gorm:"type:text"`
	ProfileImage string `json:"profileImage" gorm:"column:profile_image;null"`
	FavoriteTeams []FavoriteTeam `gorm:"foreignKey:UserId"`
	ExpeditionLikes []ExpeditionLike `gorm:"foreignKey:UserId"`
}