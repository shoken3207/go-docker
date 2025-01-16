package models

import "time"

type TempImage struct {
    BaseModel
    FileId    string    `json:"fileId" gorm:"column:file_id;not null;unique"`
    Image     string    `json:"image" gorm:"not null;unique"`
    ExpiresAt time.Time `json:"expiresAt" gorm:"not null"`
}