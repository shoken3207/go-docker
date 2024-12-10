package auth

// リクエスト
type EmailVerificationRequest struct {
	Email string `uri:"email" binding:"required,email"`
}

type RegisterRequest struct {
	Name         string `json:"name" binding:"required,min=3,max=100"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6,max=50"`
	Description  string `json:"description"`
	ProfileImage string `json:"profileImage"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=50"`
}

type UpdatePassRequest struct {
	BeforePassword string `json:"beforePassword" binding:"required,min=6,max=50"`
	AfterPassword  string `json:"afterPassword" binding:"required,min=6,max=50"`
}

type ResetPassRequest struct {
	Token         string `json:"token" binding:"required"`
	AfterPassword string `json:"afterPassword" binding:"required,min=6,max=50"`
}


// レスポンス
type LoginResponse struct {
	Token string `json:"token"`
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}