package constants

import (
	"time"
)

const (
	EmailVerificationTokenExpDate = time.Minute * 30
	LoginTokenExpDate             = time.Hour * 24 * 7
	RegisteBaserUrl               = "visitorgo://verify"
)
