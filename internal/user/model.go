package user

type Competition struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type League struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type Stadium struct {
	Id uint `json:"id"`
}

type Team struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Stadium string `json:"stadium"`
}

type FavoriteTeam struct {
	Id uint `json:"id"`
	Competition
}

// レスポンスモデル
type GetUserResponse struct {
	Id            uint   `json:"id"`
	Email         string `json:"email"`
	NickName      string `json:"nickName"`
	Introduction  string `json:"introduction"`
	FavoriteTeams string `json:"favoriteTeams"`
}