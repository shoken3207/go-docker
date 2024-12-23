package expedition

import "time"

// リクエスト
type PaymentRequest struct {
	Title string    `json:"title" binding:"required"`
	Date  time.Time `json:"date" binding:"required"`
	Cost  int       `json:"cost" binding:"required"`
}
type PaymentResponse struct {
	ID    uint      `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
	Cost  int       `json:"cost"`
}

type UpdatePaymentRequest struct {
	ID uint `json:"id" binding:"required"`
	PaymentRequest
}

type BaseGameRequest struct {
	Date    time.Time `json:"date" binding:"required"`
	Comment string    `json:"comment" binding:"required"`
	Team1Id uint      `json:"team1Id" binding:"required"`
	Team2Id uint      `json:"team2Id" binding:"required"`
}
type GameRequest struct {
	BaseGameRequest
	Scores []GameScoreRequest `json:"scores" binding:"required"`
}
type GameResponse struct {
	ID      uint                `json:"id"`
	Date    time.Time           `json:"date"`
	Comment string              `json:"comment"`
	Team1Id uint                `json:"team1Id"`
	Team2Id uint                `json:"team2Id"`
	Scores  []GameScoreResponse `json:"scores"`
}

type UpdateGameRequest struct {
	ID uint `json:"id" binding:"required"`
	GameRequest
	Scores UpdateGameScoresRequest
}
type GameScoreRequest struct {
	Order  int  `json:"order" binding:"required"`
	Score  int  `json:"score" binding:"required"`
	TeamId uint `json:"teamId" binding:"required"`
}
type GameScoreResponse struct {
	ID     uint `json:"id"`
	Order  int  `json:"order"`
	Score  int  `json:"score"`
	TeamId uint `json:"teamId"`
}

type UpdateGameScoreRequest struct {
	ID uint `json:"id" binding:"required"`
	GameScoreRequest
}

type VisitedFacilityRequest struct {
	Name      string  `json:"name" binding:"required"`
	Address   string  `json:"address" binding:"required"`
	Icon      string  `json:"icon" binding:"required"`
	Color     string  `json:"color" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}
type VisitedFacilityResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Icon      string  `json:"icon"`
	Color     string  `json:"color"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type UpdateVisitedFacilityRequest struct {
	ID uint `json:"id" binding:"required"`
	VisitedFacilityRequest
}

type ExpeditionImageRequest struct {
	FileId string `json:"fileId" binding:"required"`
	Image  string `json:"image" binding:"required"`
}
type ExpeditionImageResponse struct {
	ID     uint   `json:"id"`
	FileId string `json:"fileId"`
	Image  string `json:"image"`
}

type BaseExpeditionRequest struct {
	SportId   uint      `json:"sportId" binding:"required"`
	IsPublic  bool      `json:"isPublic" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	StartDate time.Time `json:"startDate" binding:"required"`
	EndDate   time.Time `json:"endDate" binding:"required"`
	StadiumId uint      `json:"stadiumId" binding:"required"`
	Memo      string    `json:"memo" binding:"required"`
}
type ExpeditionResponse struct {
	ID        int       `json:"id"`
	SportId   uint      `json:"sportId"`
	IsPublic  bool      `json:"isPublic"`
	Title     string    `json:"title"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	StadiumId uint      `json:"stadiumId"`
	Memo      string    `json:"memo"`
}

type CreateExpeditionRequest struct {
	BaseExpeditionRequest
	Payments          []PaymentRequest         `json:"payments" binding:"required"`
	Games             []GameRequest            `json:"games" binding:"required"`
	VisitedFacilities []VisitedFacilityRequest `json:"visitedFacilities" binding:"required"`
	Images            []ExpeditionImageRequest `json:"images" buinding:"required"`
}

type GetExpeditionDetailRequest struct {
	ExpeditionId uint `uri:"expeditionId" binding:"required"`
}

type UpdateGamesRequest struct {
	Add    []GameRequest       `json:"add"`
	Update []UpdateGameRequest `json:"update"`
	Delete []uint              `json:"delete"`
}
type UpdateGameScoresRequest struct {
	Add    []GameScoreRequest       `json:"add"`
	Update []UpdateGameScoreRequest `json:"update"`
	Delete []uint                   `json:"delete"`
}
type UpdatePaymentsRequest struct {
	Add    []PaymentRequest       `json:"add"`
	Update []UpdatePaymentRequest `json:"update"`
	Delete []uint                 `json:"delete"`
}
type UpdateVisitedFacilitiesRequest struct {
	Add    []VisitedFacilityRequest       `json:"add"`
	Update []UpdateVisitedFacilityRequest `json:"update"`
	Delete []uint                         `json:"delete"`
}

type UpdateExpeditionImagesRequest struct {
	Add    []ExpeditionImageRequest `json:"add"`
	Delete []string                 `json:"delete"`
}

type UpdateExpeditionRequestBody struct {
	ID                uint                           `json:"id" binding:"required"`
	SportId           uint                           `json:"sportId" binding:"required"`
	IsPublic          bool                           `json:"isPublic" binding:"required"`
	Title             string                         `json:"title" binding:"required"`
	StartDate         time.Time                      `json:"startDate" binding:"required"`
	EndDate           time.Time                      `json:"endDate" binding:"required"`
	StadiumId         uint                           `json:"stadiumId" binding:"required"`
	Memo              string                         `json:"memo" binding:"required"`
	Payments          UpdatePaymentsRequest          `json:"payments" binding:"required"`
	Games             UpdateGamesRequest             `json:"games" binding:"required"`
	VisitedFacilities UpdateVisitedFacilitiesRequest `json:"visitedFacilities" binding:"required"`
	Images            UpdateExpeditionImagesRequest  `json:"images" buinding:"required"`
}
type UpdateExpeditionRequestPath struct {
	ExpeditionId uint `uri:"expeditionId" binding:"required"`
}

type GetExpeditionDetailResponse struct {
	ExpeditionResponse
	VisitedFacilities []VisitedFacilityResponse `json:"visitedFacilities"`
	Payments          []PaymentResponse         `json:"payments"`
	ExpeditionImages  []ExpeditionImageResponse `json:"expeditionImages"`
	Games             []GameResponse            `json:"games"`
}
