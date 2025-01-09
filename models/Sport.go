package models

type Sport struct {
	BaseModel
	Name        string       `json:"name" gorm:"size:200;unique;not null"`
	Leagues     []League     `gorm:"foreignKey:SportId;constraint:OnDelete:CASCADE"`
	Teams       []Team       `gorm:"foreignKey:SportId;constraint:OnDelete:CASCADE"`
	Expeditions []Expedition `gorm:"foreignKey:SportId;constraint:OnDelete:CASCADE"`
}