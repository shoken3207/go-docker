package auth

// リクエスト
type EmailVerificationRequest struct {
	Email     string `form:"email" binding:"required,email" example:"tanaka@example.com"`
	TokenType string `form:"tokenType" binding:"required,oneof=register reset" example:"register"`
}

type RegisterRequest struct {
	Token        string `json:"token" binding:"required" example:"1234567890"`
	Username     string `json:"username" binding:"required,min=1,max=255" example:"user123"`
	Name         string `json:"name" binding:"required,min=1,max=100" example:"tanaka taro"`
	Password     string `json:"password" binding:"required,min=6,max=50" example:"password123"`
	Description  string `json:"description" example:"野球が好きです！"`
	ProfileImage string `json:"profileImage" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
	FavoriteTeamIds []uint `json:"favoriteTeamIds" example:"1"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"tanaka@example.com"`
	Password string `json:"password" binding:"required,min=6,max=50" example:"password123"`
}

type UpdatePassRequestBody struct {
	BeforePassword string `json:"beforePassword" binding:"required,min=6,max=50" example:"password123"`
	AfterPassword  string `json:"afterPassword" binding:"required,min=6,max=50" example:"password456"`
}

type ResetPassRequest struct {
	Token         string `json:"token" binding:"required" example:"1234567890"`
	AfterPassword string `json:"afterPassword" binding:"required,min=6,max=50" example:"password456"`
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
