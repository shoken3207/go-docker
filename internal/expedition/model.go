package expedition

import "time"

// リクエスト
type PaymentRequest struct {
	Title string    `json:"title" binding:"required" example:"チケット代"`
	Date  time.Time `json:"date" binding:"required" example:"2025-01-01T00:00:00Z"`
	Cost  int       `json:"cost" binding:"required" example:"5000"`
}
type PaymentResponse struct {
	ID    uint      `json:"id" example:"1"`
	Title string    `json:"title" example:"チケット代"`
	Date  time.Time `json:"date" example:"2025-01-01T00:00:00Z"`
	Cost  int       `json:"cost" example:"5000"`
}

type UpdatePaymentRequest struct {
	ID uint `json:"id" binding:"required" example:"1"`
	PaymentRequest
}

type BaseGameRequest struct {
	Date    time.Time `json:"date" binding:"required" example:"2025-01-01T00:00:00Z"`
	Team1Id uint      `json:"team1Id" binding:"required" example:"1"`
	Team2Id uint      `json:"team2Id" binding:"required" example:"2"`
}
type GameRequest struct {
	BaseGameRequest
	Scores []GameScoreRequest `json:"scores" binding:"required"`
}
type GameResponse struct {
	ID      uint                `json:"id" example:"1"`
	Date    time.Time           `json:"date" example:"2025-01-01T00:00:00Z"`
	Team1Id uint                `json:"team1Id" example:"1"`
	Team1Name string `json:"team1Name" example:"ヤクルト"`
	Team2Id uint                `json:"team2Id" example:"2"`
	Team2Name string `json:"team2Name" example:"ソフトバンク"`
	Scores  []GameScoreResponse `json:"scores"`
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
type GameScoreResponse struct {
	ID     uint `json:"id" example:"1"`
	Order  int  `json:"order" example:"1"`
	Score  int  `json:"score" example:"1"`
	TeamId uint `json:"teamId" example:"1"`
	TeamName string `json:"teamName" example:"ヤクルト"`
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
type VisitedFacilityResponse struct {
	ID        int     `json:"id" example:"1"`
	Name      string  `json:"name" example:"東京駅"`
	Address   string  `json:"address" example:"東京都千代田区丸の内1-1-1"`
	Icon      string  `json:"icon" example:"train"`
	Color     string  `json:"color" example:"#00FF00"`
	Latitude  float64 `json:"latitude" example:"35.6812"`
	Longitude float64 `json:"longitude" example:"139.7671"`
}

type UpdateVisitedFacilityRequest struct {
	ID uint `json:"id" binding:"required" example:"1"`
	VisitedFacilityRequest
}

type ExpeditionImageRequest struct {
	FileId string `json:"fileId" binding:"required" example:"file_1234567890"`
	Image  string `json:"image" binding:"required" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
}
type ExpeditionImageResponse struct {
	ID     uint   `json:"id" example:"1"`
	FileId string `json:"fileId" example:"file_1234567890"`
	Image  string `json:"image" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
}

type BaseExpeditionRequest struct {
	SportId   uint      `json:"sportId" binding:"required" example:"1"`
	IsPublic  bool      `json:"isPublic" binding:"required" example:"true"`
	Title     string    `json:"title" binding:"required" example:"野球観戦の遠征記録"`
	StartDate time.Time `json:"startDate" binding:"required" example:"2025-01-01T00:00:00Z"`
	EndDate   time.Time `json:"endDate" binding:"required" example:"2025-01-01T00:00:00Z"`
	StadiumId uint      `json:"stadiumId" binding:"required" example:"1"`
	Memo      string    `json:"memo" binding:"required" example:"初めてのスタジアム訪問。とても楽しかった！"`
}
type ExpeditionResponse struct {
	ID        int       `json:"id" example:"1"`
	SportId   uint      `json:"sportId" example:"1"`
	SportName string    `json:"sportName" example:"野球"`
	IsPublic  bool      `json:"isPublic" example:"true"`
	Title     string    `json:"title" example:"野球観戦の遠征記録"`
	StartDate time.Time `json:"startDate" example:"2025-01-01T00:00:00Z"`
	EndDate   time.Time `json:"endDate" example:"2025-01-01T00:00:00Z"`
	StadiumId uint      `json:"stadiumId" example:"1"`
	StadiumName string `json:"stadiumName" example:"東京ドーム"`
	Memo      string    `json:"memo" example:"初めてのスタジアム訪問。とても楽しかった！"`
}

type CreateExpeditionRequest struct {
	BaseExpeditionRequest
	Payments          []PaymentRequest         `json:"payments" binding:"required"`
	Games             []GameRequest            `json:"games" binding:"required"`
	VisitedFacilities []VisitedFacilityRequest `json:"visitedFacilities" binding:"required"`
	ImageUrls         []string                 `json:"imageUrls" binding:"required"`
}

type GetExpeditionDetailRequest struct {
	ExpeditionId uint `uri:"expeditionId" binding:"required" example:"1"`
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
	Add    []string `json:"add"`
	Delete []string `json:"delete"`
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

type ExpeditionListResponse struct {
	ID          uint                   `json:"id" example:"1"`
	Title       string                 `json:"title" example:"野球観戦の遠征記録"`
	StartDate   time.Time             `json:"startDate" example:"2025-01-01T00:00:00Z"`
	EndDate     time.Time             `json:"endDate" example:"2025-01-01T00:00:00Z"`
	SportId     uint                   `json:"sportId" example:"1"`
	SportName   string                 `json:"sportName" example:"野球"`
	UserID      uint                   `json:"userId" example:"1"`
	UserName    string                 `json:"userName" example:"ユーザー名"`
	UserIcon    string                 `json:"userIcon" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
	LikesCount  int64                  `json:"likesCount" example:"10"`
	Images      []string               `json:"images" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
	Team1Name   string                 `json:"team1Name" example:"ヤクルト"`
	Team2Name   string                 `json:"team2Name" example:"ソフトバンク"`
}

type ExpeditionListRequest struct {
	Page     int        `form:"page" binding:"required,min=1" example:"1"`
	SportId  *uint      `form:"sportId" example:"1"`
	TeamId   *uint      `form:"teamId" example:"1"`
}

type GetExpeditionDetailResponse struct {
	ExpeditionResponse
	VisitedFacilities []VisitedFacilityResponse `json:"visitedFacilities"`
	Payments          []PaymentResponse         `json:"payments"`
	ExpeditionImages  []ExpeditionImageResponse `json:"expeditionImages"`
	Games             []GameResponse            `json:"games"`
}
