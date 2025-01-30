package models

import "gorm.io/gorm"

type Stadium struct {
	gorm.Model
	Name        string       `json:"name" gorm:"size:200;not null;unique"`
	Description string       `json:"description" gorm:"type:text"`
	Address     string       `json:"address" gorm:"size:200;not null"`
	Capacity    int          `json:"capacity"`
	Image       string       `json:"image" gorm:"not null"`
	FileId      *string      `json:"fileId" gorm:"column:file_id"`
	Attribution *string      `json:"attribution"`
	Teams       []Team       `gorm:"foreignKey:StadiumId"`
	Expeditions []Expedition `gorm:"foreignKey:StadiumId"`
}

func (s *Stadium) SetFileId(fileId string) {
	if fileId == "" {
		s.FileId = nil
	} else {
		s.FileId = &fileId
	}
}
