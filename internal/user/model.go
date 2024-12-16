package user

// リクエスト
type GetUserByIdRequest struct {
	UserId uint `uri:"userId" binding:"required"`
}

type UpdateUserRequestBody struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	ProfileImage string `json:"profileImage" binding:"required"`
}
type UpdateUserRequestPath struct {
	UserId uint `uri:"userId" binding:"required"`
}

// レスポンス
type UserResponse struct {
	Id           uint   `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	ProfileImage string `json:"profileImage,omitempty"`
}
