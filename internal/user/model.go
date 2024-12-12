package user

// リクエスト
type GetUserByIdRequest struct {
	UserId uint `uri:"userId" binding:"required"`
}

// レスポンスモデル
type GetUserResponse struct {
	Id            uint   `json:"id"`
	Email         string `json:"email"`
	NickName      string `json:"nickName"`
	Introduction  string `json:"introduction"`
	FavoriteTeams string `json:"favoriteTeams"`
}
