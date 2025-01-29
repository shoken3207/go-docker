package models

type Session struct {
	BaseModel
	UserId uint   `json:"userId" gorm:"not null"`
	Token  string `json:"token" gorm:"not null;unique"`
	User   User   `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}