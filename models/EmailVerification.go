package models

import (
	"time"
)

type EmailVerification struct {
	BaseModel
	Email string `json:"email" gorm:"size:100;not null"`
	Token string `json:"token" gorm:"not null"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"column:expires_at;not null"`
}