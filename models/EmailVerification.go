package models

import (
	"time"

	"gorm.io/gorm"
)

type EmailVerification struct {
	gorm.Model
	Email string `json:"email" gorm:"size:100;not null"`
	Token string `json:"token" gorm:"not null"`
	ExpiresAt time.Time `json:"expiresAt" gorm:"column:expires_at;not null"`
}