package expedition

import "time"

// リクエスト
type PaymentRequest struct {
	Title string    `json:"title" binding:"required"`
	Date  time.Time `json:"date" binding:"required"`
	Cost  int       `json:"cost" binding:"required"`
}

type GameRequest struct {
	Date      time.Time          `json:"date" binding:"required"`
	Comment   string             `json:"comment" binding:"required"`
	StadiumId uint               `json:"stadiumId" binding:"required"`
	Team1Id   uint               `json:"team1Id" binding:"required"`
	Team2Id   uint               `json:"team2Id" binding:"required"`
	Scores    []GameScoreRequest `json:"scores" binding:"required"`
}

type GameScoreRequest struct {
	Order  int  `json:"order" binding:"required"`
	Score  int  `json:"score" binding:"required"`
	TeamId uint `json:"teamId" binding:"required"`
}

type VisitedFacilityRequest struct {
	Name      string  `json:"name" binding:"required"`
	Address   string  `json:"address" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

type CreateExpeditionRequest struct {
	SportId           uint                     `json:"sportId" binding:"required"`
	IsPublic          bool                     `json:"isPublic" binding:"required"`
	Title             string                   `json:"title" binding:"required"`
	StartDate         time.Time                `json:"startDate" binding:"required"`
	EndDate           time.Time                `json:"endDate" binding:"required"`
	Memo              string                   `json:"memo" binding:"required"`
	Payments          []PaymentRequest         `json:"payments" binding:"required"`
	Games             []GameRequest            `json:"games" binding:"required"`
	VisitedFacilities []VisitedFacilityRequest `json:"visitedFacilities" binding:"required"`
}
