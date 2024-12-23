package expedition

import "time"

// リクエスト
type PaymentRequest struct {
	Title string    `json:"title" binding:"required"`
	Date  time.Time `json:"date" binding:"required"`
	Cost  int       `json:"cost" binding:"required"`
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

type UpdateVisitedFacilityRequest struct {
	ID uint `json:"id" binding:"required"`
	VisitedFacilityRequest
}

type ExpeditionImageRequest struct {
	FileId string `json:"fileId" binding:"required"`
	Image  string `json:"image" binding:"required"`
}
type UpdateExpeditionImageRequest struct {
	ID uint `json:"id" binding:"required"`
	ExpeditionImageRequest
}

type CreateExpeditionRequest struct {
	SportId           uint                     `json:"sportId" binding:"required"`
	IsPublic          bool                     `json:"isPublic" binding:"required"`
	Title             string                   `json:"title" binding:"required"`
	StartDate         time.Time                `json:"startDate" binding:"required"`
	EndDate           time.Time                `json:"endDate" binding:"required"`
	StadiumId         uint                     `json:"stadiumId" binding:"required"`
	Memo              string                   `json:"memo" binding:"required"`
	Payments          []PaymentRequest         `json:"payments" binding:"required"`
	Games             []GameRequest            `json:"games" binding:"required"`
	VisitedFacilities []VisitedFacilityRequest `json:"visitedFacilities" binding:"required"`
	Images            []ExpeditionImageRequest `json:"images" buinding:"required"`
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
	Add    []ExpeditionImageRequest       `json:"add"`
	Delete []UpdateExpeditionImageRequest `json:"delete"`
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
