package models

type Stadium struct {
	BaseModel
	Name        string       `json:"name" gorm:"size:200;not null;unique"`
	Description string       `json:"description" gorm:"type:text"`
	Address     string       `json:"address" gorm:"size:200;not null"`
	Capacity    int          `json:"capacity"`
	Image       string       `json:"image" gorm:"not null"`
	Teams       []Team       `gorm:"foreignKey:StadiumId;constraint:OnDelete:CASCADE"`
	Expeditions []Expedition `gorm:"foreignKey:StadiumId;constraint:OnDelete:CASCADE"`
}
