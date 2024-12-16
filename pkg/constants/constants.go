package constants

import (
	"time"
)

const (
	EmailVerificationTokenExpDate = time.Minute * 30
	LoginTokenExpDate             = time.Hour * 24 * 7
	RegisteBaserUrl               = "visitorgo://verify"
	MaxFileSize = 5 * 1024 * 1024
)
var AllowedExtensions = [...]string{".jpg", ".jpeg", ".png"}