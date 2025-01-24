package auth

// リクエスト
type EmailVerificationRequestQuery struct {
	Email     string `form:"email" binding:"required,email" example:"tanaka@example.com" field:"メールアドレス"`
	TokenType string `form:"tokenType" binding:"required,oneof=register reset" example:"register" field:"トークンタイプ"`
}

type RegisterRequestBody struct {
	Token        string `json:"token" binding:"required" example:"1234567890" field:"トークン"`
	Username     string `json:"username" binding:"required,min=1,max=255" example:"user123" field:"ユーザー名"`
	Name         string `json:"name" binding:"required,min=1,max=100" example:"tanaka taro" field:"名前"`
	Password     string `json:"password" binding:"required,min=6,max=50" example:"password123" field:"パスワード"`
	Description  string `json:"description" example:"野球が好きです！" field:"紹介文"`
	ProfileImage string `json:"profileImage" binding:"omitempty,url" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg" field:"プロフィール画像"`
	FavoriteTeamIds []uint `json:"favoriteTeamIds" example:"1" field:"お気に入りチームのid配列"`
}

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email" example:"tanaka@example.com" field:"メールアドレス"`
	Password string `json:"password" binding:"required,min=6,max=50" example:"password123" field:"パスワード"`
}

type UpdatePassRequestBody struct {
	BeforePassword string `json:"beforePassword" binding:"required,min=6,max=50" example:"password123" field:"現在のパスワード"`
	AfterPassword  string `json:"afterPassword" binding:"required,min=6,max=50" example:"password456" field:"新しいパスワード"`
}

type ResetPassRequestBody struct {
	Token         string `json:"token" binding:"required" example:"1234567890" field:"トークン"`
	AfterPassword string `json:"afterPassword" binding:"required,min=6,max=50" example:"password456" field:"新しいパスワード"`
}

// レスポンス
type LoginResponse struct {
	Token string `json:"token" example:"1234567890"`
}

// その他
type TokenRequest struct {
	UserID    *uint
	Email     *string
	TokenType *string
}
