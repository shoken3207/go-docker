package admintool

// リクエスト
type teamAddRequest struct {
	Sports   string `uri:"sports" binding:"required"`
	League   string `uri:"sports" binding:"required"`
	TeamName string `uri:"teamName" binding:"required"`
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
	Sports   string `json:"sports"`
	League   string `json:"league"`
	teamId   uint   `json:"id"`
	teamName string `json:"teamName"`
}
