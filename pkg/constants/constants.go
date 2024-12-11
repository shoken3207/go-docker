package constants

import (
	"time"
)

const (
	EmailVerificationExpDate = time.Hour
	JwtTokenExpDate          = time.Hour * 24 * 7
	RegisteBaserUrl          = "http://localhost:3000/register"
)
