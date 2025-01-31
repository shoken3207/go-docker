package team

type GetTeamBySportsIdRequestPath struct {
	SportsId uint `uri:"sportsId" binding:"required" example:"1" field:"スポーツId"`
}

type TeamResponse struct {
	ID         uint   `json:"id" example:"1"`
	Name       string `json:"name" example:"ＦＣ東京"`
	IsFavorite bool   `json:"isFavorite" example:"true"`
}

type LeagueResponse struct {
	League string         `json:"league" example:"J1リーグ"`
	Teams  []TeamResponse `json:"teams"`
}

type SportResponse struct {
	Sport string           `json:"sports" example:"サッカー"`
	Icon  string           `json:"icon" example:"soccer"`
	Team  []LeagueResponse `json:"team"`
}

type TeamListResponse struct {
	ID   uint   `json:"id" example:"1"`
	Name string `json:"name" example:"オリックス・バファローズ"`
}