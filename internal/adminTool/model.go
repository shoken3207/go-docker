package adminTool

// リクエスト
type teamAddRequest struct {
	StadiumId uint   `json:"stadiumId" binding:"required"`
	SportsId  uint   `json:"sportsId" binding:"required"`
	LeagueId  uint   `json:"LeagueId" binding:"required"`
	TeamName  string `json:"teamName" binding:"required"`
}

type UppdateTeamRequest struct {
	BeforTeamName string `json:"beforeTeamName" binding:"required,max=50"`
	AfterTeamName string `json:"afterTeamName" binding:"required,max=50"`
}

type DeleteTeamRequest struct {
	TeamId uint `uri:"id" binding:"required"`
}

// レスポンス
type team struct {
	StadiumId uint   `json:"stadiumId"`
	SportsId  uint   `json:"sportsId"`
	LeagueId  uint   `json:"LeagueId"`
	TeamName  string `json:"teamName"`
}
