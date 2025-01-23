package user

import "go-docker/internal/expedition"

// リクエスト
type GetUserByIdRequest struct {
	UserId uint `uri:"userId" binding:"required" example:"1"`
}

type GetUserByUsernameRequest struct {
	Username string `uri:"username" binding:"required,min=5,max=255" example:"user123"`
}

type UpdateUserRequestBody struct {
	Username     string `json:"username" binding:"required,min=1,max=255" example:"user123"`
	Name         string `json:"name" binding:"required" example:"tanaka taro"`
	Description  string `json:"description" example:"野球が好きです！"`
	ProfileImage string `json:"profileImage" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
	FavoriteTeams []uint `json:"favoriteTeams" example:"1"`
}

type IsUniqueUsernameRequest struct {
	Username string `uri:"username" binding:"required,min=1,max=255" example:"user123"`
}

// レスポンス
type UserResponse struct {
	Id           uint   `json:"id" example:"1"`
	Username     string `json:"username" example:"user123"`
	Email        string `json:"email" example:"tanaka@example.com"`
	FileId       string `json:"fileId" example:"1234567890"`
	Name         string `json:"name" example:"tanaka taro"`
	Description  string `json:"description" example:"野球が好きです！"`
	ProfileImage string `json:"profileImage" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
}

type UserDetailResponse struct {
	UserResponse
	Expeditions []expedition.ExpeditionListResponse `json:"expeditions"`
	LikedExpeditions []expedition.ExpeditionListResponse `json:"likedExpeditions"`
}

type IsUniqueUsernameResponse struct {
	IsUnique bool `json:"isUnique" example:"true"`
	Message string `json:"message" example:"使用できます"`
}
