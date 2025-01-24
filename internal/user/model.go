package user

import "go-docker/internal/expedition"

// リクエスト
type GetUserByIdRequestPath struct {
	UserId uint `uri:"userId" binding:"required" example:"1" field:"ユーザーID"`
}

type GetUserByUsernameRequestPath struct {
	Username string `uri:"username" binding:"required,min=1,max=255" example:"user123" field:"ユーザー名"`
}

type UpdateUserRequestBody struct {
	Username     string `json:"username" binding:"required,min=1,max=255" example:"user123" field:"ユーザー名"`
	Name         string `json:"name" binding:"required" example:"tanaka taro" field:"名前"`
	Description  string `json:"description" example:"野球が好きです！" field:"自己紹介"`
	ProfileImage string `json:"profileImage" binding:"omitempty,url" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg" field:"プロフィール画像"`
	FavoriteTeams []uint `json:"favoriteTeams" example:"1" field:"お気に入りチーム"`
}

type IsUniqueUsernameRequestPath struct {
	Username string `uri:"username" binding:"required,min=1,max=255" example:"user123" field:"ユーザー名"`
}

// レスポンス
type FavoriteTeamResponse struct {
	ID        uint   `json:"id"`
	TeamID    uint   `json:"teamId"`
	TeamName  string `json:"teamName"`
	LeagueName string `json:"leagueName"`
	SportName  string `json:"sportName"`
}

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
	FavoriteTeams []FavoriteTeamResponse `json:"favoriteTeams"`
}

type IsUniqueUsernameResponse struct {
	IsUnique bool `json:"isUnique" example:"true"`
	Message string `json:"message" example:"使用できます"`
}
