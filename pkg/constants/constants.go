package constants

import (
	"time"
)

const (
	LIMIT_EXPEDITION_LIST = 15
	EmailVerificationTokenExpDate = time.Minute * 30
	LoginTokenExpDate             = time.Hour * 24 * 7
	RegisteBaserUrl               = "visitorgo://verify"
	MaxFileSize                   = 5 * 1024 * 1024
	MailBody                      = `
%s

こんにちは。

この度は%s

以下のリンクから、%s

確認リンク:
%s

※ 上記のリンクは、発行から30分以内にご利用ください。期限が過ぎると、再度新しいリンクをリクエストする必要があります。

もし、ご不明点がございましたら、お気軽にお問い合わせください。

どうぞよろしくお願いいたします。

ビジターGOサポートチーム
`
)

var AllowedExtensions = [...]string{".jpg", ".jpeg", ".png"}

var SportIcon = map[int]string{
	1: "soccer",
	2: "baseball",
	3: "basketball",
	4: "volleyball",
}