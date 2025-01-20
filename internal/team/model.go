package team

type TeamResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	IsFavorite bool   `json:"isFavorite"`
}

type LeagueResponse struct {
	League string         `json:"league"`
	Teams  []TeamResponse `json:"teams"`
}

type SportResponse struct {
	Sport string           `json:"sports"`
	Icon  string           `json:"icon"`
	Team  []LeagueResponse `json:"team"`
}