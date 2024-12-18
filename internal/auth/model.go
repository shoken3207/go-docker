package auth

// リクエスト
type EmailVerificationRequest struct {
	Email     string `form:"email" binding:"required,email"`
	TokenType string `form:"tokenType" binding:"required,oneof=register reset"`
}

type RegisterRequest struct {
	Token        string `json:"token" binding:"required"`
	Username     string `json:"username" binding:"required,min=5,max=255"`
	Name         string `json:"name" binding:"required,min=3,max=100"`
	Password     string `json:"password" binding:"required,min=6,max=50"`
	Description  string `json:"description"`
	ProfileImage string `json:"profileImage"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=50"`
}

type UpdatePassRequestBody struct {
	BeforePassword string `json:"beforePassword" binding:"required,min=6,max=50"`
	AfterPassword  string `json:"afterPassword" binding:"required,min=6,max=50"`
}
type UpdateUserRequestPath struct {
	UserId uint `uri:"userId" binding:"required"`
}

type ResetPassRequest struct {
	Token         string `json:"token" binding:"required"`
	AfterPassword string `json:"afterPassword" binding:"required,min=6,max=50"`
}

// レスポンス
type LoginResponse struct {
	Token string `json:"token"`
}

// その他
type TokenRequest struct {
	UserID    *uint
	Email     *string
	TokenType *string
}
