package team

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