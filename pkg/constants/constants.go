package constants

import (
	"time"
)

const (
	EmailVerificationExpDate = time.Minute * 30
	JwtTokenExpDate          = time.Hour * 24 * 7
	RegisteBaserUrl          = "visitorgo://verify"
)
