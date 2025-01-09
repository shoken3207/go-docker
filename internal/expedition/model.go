package expedition

import "time"

// リクエスト
type PaymentRequest struct {
	Title string    `json:"title" binding:"required" example:"チケット代"`
	Date  time.Time `json:"date" binding:"required" example:"2025-01-01T00:00:00Z"`
	Cost  int       `json:"cost" binding:"required" example:"5000"`
}

type UpdatePaymentRequest struct {
	ID uint `json:"id" binding:"required" example:"1"`
	PaymentRequest
}

type BaseGameRequest struct {
	Date    time.Time `json:"date" binding:"required" example:"2025-01-01T00:00:00Z"`
	Comment string    `json:"comment" binding:"required" example:"熱い試合でした！！"`
	Team1Id uint      `json:"team1Id" binding:"required" example:"1"`
	Team2Id uint      `json:"team2Id" binding:"required" example:"2"`
}
type GameRequest struct {
	BaseGameRequest
	Scores []GameScoreRequest `json:"scores" binding:"required"`
}

type UpdateGameRequest struct {
	ID uint `json:"id" binding:"required" example:"1"`
	BaseGameRequest
	Scores UpdateGameScoresRequest `json:"scores" binding:"required"`
}
type GameScoreRequest struct {
	Order  int  `json:"order" binding:"required" example:"1"`
	Score  int  `json:"score" binding:"required" example:"1"`
	TeamId uint `json:"teamId" binding:"required" example:"1"`
}

type UpdateGameScoreRequest struct {
	ID uint `json:"id" binding:"required" example:"1"`
	GameScoreRequest
}

type VisitedFacilityRequest struct {
	Name      string  `json:"name" binding:"required" example:"東京駅"`
	Address   string  `json:"address" binding:"required" example:"東京都千代田区丸の内1-1-1"`
	Icon      string  `json:"icon" binding:"required" example:"train"`
	Color     string  `json:"color" binding:"required" example:"#00FF00"`
	Latitude  float64 `json:"latitude" binding:"required" example:"35.6812"`
	Longitude float64 `json:"longitude" binding:"required" example:"139.7671"`
}

type UpdateVisitedFacilityRequest struct {
	ID uint `json:"id" binding:"required" example:"1"`
	VisitedFacilityRequest
}

type ExpeditionImageRequest struct {
	FileId string `json:"fileId" binding:"required" example:"file_1234567890"`
	Image  string `json:"image" binding:"required" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
}

type CreateExpeditionRequest struct {
	SportId           uint                     `json:"sportId" binding:"required" example:"1"`
	IsPublic          bool                     `json:"isPublic" example:"true"`
	Title             string                   `json:"title" binding:"required" example:"野球観戦の遠征記録"`
	StartDate         time.Time                `json:"startDate" binding:"required" example:"2025-01-01T00:00:00Z"`
	EndDate           time.Time                `json:"endDate" binding:"required" example:"2025-01-01T00:00:00Z"`
	StadiumId         uint                     `json:"stadiumId" binding:"required" example:"1"`
	Memo              string                   `json:"memo" binding:"required" example:"初めてのスタジアム訪問。とても楽しかった！"`
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
	Add    []ExpeditionImageRequest `json:"add"`
	Delete []string                 `json:"delete"`
}

type UpdateExpeditionRequestBody struct {
	SportId           uint                           `json:"sportId" binding:"required" example:"1"`
	IsPublic          bool                           `json:"isPublic" example:"true"`
	Title             string                         `json:"title" binding:"required" example:"野球観戦の遠征記録"`
	StartDate         time.Time                      `json:"startDate" binding:"required" example:"2025-01-01T00:00:00Z"`
	EndDate           time.Time                      `json:"endDate" binding:"required" example:"2025-01-01T00:00:00Z"`
	StadiumId         uint                           `json:"stadiumId" binding:"required" example:"1"`
	Memo              string                         `json:"memo" binding:"required" example:"初めてのスタジアム訪問。とても楽しかった！"`
	Payments          UpdatePaymentsRequest          `json:"payments" binding:"required"`
	Games             UpdateGamesRequest             `json:"games" binding:"required"`
	VisitedFacilities UpdateVisitedFacilitiesRequest `json:"visitedFacilities" binding:"required"`
	Images            UpdateExpeditionImagesRequest  `json:"images" binding:"required"`
}
type UpdateExpeditionRequestPath struct {
	ExpeditionId uint `uri:"expeditionId" binding:"required" example:"1"`
}

// ... existing code ...

type LikeExpeditionRequestPath struct {
	ExpeditionId uint `uri:"expeditionId" binding:"required" example:"1"`
}

type UnlikeExpeditionRequestPath struct {
	ExpeditionId uint `uri:"expeditionId" binding:"required" example:"1"`
}

type DeleteExpeditionRequestPath struct {
	ExpeditionId uint `uri:"expeditionId" binding:"required" example:"1"`
}
