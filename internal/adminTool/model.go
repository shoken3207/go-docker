package adminTool

// リクエスト
type TeamAddRequest struct {
	StadiumId uint   `json:"stadiumId" binding:"required"`
	SportsId  uint   `json:"sportsId" binding:"required"`
	LeagueId  uint   `json:"LeagueId" binding:"required"`
	TeamName  string `json:"teamName" binding:"required"`
}

type StadiumAddRequest struct {
	StadiumId   uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Capacity    uint   `json:"capacity" binding:"required"`
	Image       string `json:"image" binding:"required"`
}

type UppdateTeamRequest struct {
	BeforTeamName string `json:"beforeTeamName" binding:"required,max=50"`
	AfterTeamName string `json:"afterTeamName" binding:"required,max=50"`
}

type DeleteRequest struct {
	Id uint `uri:"id" binding:"required"`
}

// レスポンス
type Team struct {
	StadiumId uint   `json:"stadiumId"`
	SportsId  uint   `json:"sportsId"`
	LeagueId  uint   `json:"LeagueId"`
	TeamName  string `json:"teamName"`
}

type Stadium struct {
	StadiumId   uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Capacity    uint   `json:"capacity" binding:"required"`
	Image       string `json:"image" binding:"required"`
}
