package user

// リクエスト
type GetUserByIdRequest struct {
	UserId uint `uri:"userId" binding:"required"`
}

type GetUserByUsernameRequest struct {
	Username string `uri:"username" binding:"required,min=5,max=255"`
}

type UpdateUserRequestBody struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	ProfileImage string `json:"profileImage" binding:"required"`
}
type UpdateUserRequestPath struct {
	UserId uint `uri:"userId" binding:"required"`
}

type IsUniqueUsernameRequest struct {
	Username string `uri:"username" binding:"required,min=5,max=255"`
}

// レスポンス
type UserResponse struct {
	Id           uint   `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	ProfileImage string `json:"profileImage,omitempty"`
}
type IsUniqueUsernameResponse struct {
	IsUnique bool `json:"isUnique"`
}
