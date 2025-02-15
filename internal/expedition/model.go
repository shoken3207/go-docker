package expedition

import "time"

type PaymentRequest struct {
	Title string    `json:"title" binding:"required" example:"チケット代" field:"タイトル"`
	Date  time.Time `json:"date" binding:"required" example:"2025-01-01T00:00:00Z" field:"日付"`
	Cost  int       `json:"cost" binding:"required" example:"5000" field:"金額"`
}
type PaymentResponse struct {
	ID    uint      `json:"id" example:"1"`
	Title string    `json:"title" example:"チケット代"`
	Date  time.Time `json:"date" example:"2025-01-01T00:00:00Z"`
	Cost  int       `json:"cost" example:"5000"`
}

type UpdatePaymentRequest struct {
	ID uint `json:"id" binding:"required" example:"1" field:"支払いID"`
	PaymentRequest
}

type BaseGameRequest struct {
	Date    time.Time `json:"date" binding:"required" example:"2025-01-01T00:00:00Z" field:"日付"`
	Team1Id uint      `json:"team1Id" binding:"required" example:"1" field:"チーム1ID"`
	Team2Id uint      `json:"team2Id" binding:"required" example:"2" field:"チーム2ID"`
}
type GameRequest struct {
	BaseGameRequest
	Scores []GameScoreRequest `json:"scores" binding:"required,dive" field:"試合スコア"`
}
type GameResponse struct {
	ID      uint                `json:"id" example:"1"`
	Date    time.Time           `json:"date" example:"2025-01-01T00:00:00Z"`
	Team1Id uint                `json:"team1Id" example:"1"`
	Team2Id uint                `json:"team2Id" example:"2"`
	Scores  []GameScoreResponse `json:"scores"`
}

type UpdateGameRequest struct {
	ID uint `json:"id" binding:"required" example:"1" field:"試合ID"`
	BaseGameRequest
	Scores UpdateGameScoresRequest `json:"scores" binding:"required" field:"試合スコア"`
}
type GameScoreRequest struct {
	Order      int  `json:"order" binding:"gte=0" example:"1" field:"順番"`
	Team1Score int  `json:"team1Score" binding:"gte=0" example:"1" field:"チーム1スコア"`
	Team2Score int  `json:"team2Score" binding:"gte=0" example:"2" field:"チーム2スコア"`
}
type GameScoreResponse struct {
	ID         uint   `json:"id" example:"1"`
	Order      int    `json:"order" example:"1"`
	Team1Score int    `json:"team1Score" example:"1"`
	Team2Score int    `json:"team2Score" example:"2"`
	Team1Name  string `json:"team1Name" example:"ヤクルト"`
	Team2Name  string `json:"team2Name" example:"ソフトバンク"`
}

type UpdateGameScoreRequest struct {
	ID uint `json:"id" binding:"required" example:"1" field:"試合スコアID"`
	GameScoreRequest
}

type VisitedFacilityRequest struct {
	Name      string  `json:"name" binding:"required" example:"東京駅" field:"施設名"`
	CustomName      string  `json:"customName" binding:"required" example:"東京駅(おみやげ)" field:"カスタム名"`
	Address   string  `json:"address" binding:"required" example:"東京都千代田区丸の内1-1-1" field:"住所"`
	Icon      string  `json:"icon" binding:"required" example:"train" field:"アイコン"`
	Color     string  `json:"color" binding:"required" example:"#00FF00" field:"色"`
	Latitude  float64 `json:"latitude" binding:"required" example:"35.6812" field:"緯度"`
	Longitude float64 `json:"longitude" binding:"required" example:"139.7671" field:"経度"`
}
type VisitedFacilityResponse struct {
	ID        int     `json:"id" example:"1"`
	Name      string  `json:"name" example:"東京駅"`
	CustomName string `json:"customName" example:"東京駅（おみやげ）"`
	Address   string  `json:"address" example:"東京都千代田区丸の内1-1-1"`
	Icon      string  `json:"icon" example:"train"`
	Color     string  `json:"color" example:"#00FF00"`
	Latitude  float64 `json:"latitude" example:"35.6812"`
	Longitude float64 `json:"longitude" example:"139.7671"`
}

type UpdateVisitedFacilityRequest struct {
	ID uint `json:"id" binding:"required" example:"1" field:"施設ID"`
	VisitedFacilityRequest
}

type ExpeditionImageResponse struct {
	ID     uint   `json:"id" example:"1"`
	FileId string `json:"fileId" example:"file_1234567890"`
	Image  string `json:"image" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
}

type BaseExpeditionRequest struct {
	SportId   uint      `json:"sportId" binding:"required" example:"1" field:"スポーツID"`
	IsPublic  *bool      `json:"isPublic" binding:"required" example:"true" field:"公開設定"`
	Title     string    `json:"title" binding:"required" example:"野球観戦の遠征記録" field:"タイトル"`
	StartDate time.Time `json:"startDate" binding:"required" example:"2025-01-01T00:00:00Z" field:"開始日"`
	EndDate   time.Time `json:"endDate" binding:"required" example:"2025-01-01T00:00:00Z" field:"終了日"`
	StadiumId uint      `json:"stadiumId" binding:"required" example:"1" field:"スタジアムID"`
	Memo      string    `json:"memo" binding:"required" example:"初めてのスタジアム訪問。とても楽しかった！" field:"メモ"`
}
type ExpeditionResponse struct {
	ID        int       `json:"id" example:"1"`
	UserId 	  uint      `json:"userId" example:"1"`
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

type CreateExpeditionRequestBody struct {
	BaseExpeditionRequest
	Payments          []PaymentRequest         `json:"payments" binding:"required,dive" field:"支払い"`
	Games             []GameRequest            `json:"games" binding:"required,dive" field:"試合"`
	VisitedFacilities []VisitedFacilityRequest `json:"visitedFacilities" binding:"required,dive" field:"周辺施設"`
	ImageUrls         []string                 `json:"imageUrls" binding:"required" field:"画像URL配列"`
}

type GetExpeditionDetailRequestPath struct {
	ExpeditionId uint `uri:"expeditionId" binding:"required" example:"1" field:"遠征記録ID"`
}

type UpdateGamesRequest struct {
	Add    []GameRequest       `json:"add" binding:"required,dive" field:"追加"`
	Update []UpdateGameRequest `json:"update" binding:"required,dive" field:"更新"`
	Delete []uint              `json:"delete" binding:"required,dive" field:"削除"`
}
type UpdateGameScoresRequest struct {
	Add    []GameScoreRequest       `json:"add" binding:"required,dive" field:"追加"`
	Update []UpdateGameScoreRequest `json:"update" binding:"required,dive" field:"更新"`
	Delete []uint                   `json:"delete" binding:"required,dive" field:"削除"`
}
type UpdatePaymentsRequest struct {
	Add    []PaymentRequest       `json:"add" binding:"required,dive" field:"追加"`
	Update []UpdatePaymentRequest `json:"update" binding:"required,dive" field:"更新"`
	Delete []uint                 `json:"delete" binding:"required,dive" field:"削除"`
}
type UpdateVisitedFacilitiesRequest struct {
	Add    []VisitedFacilityRequest       `json:"add" binding:"required,dive" field:"追加"`
	Update []UpdateVisitedFacilityRequest `json:"update" binding:"required,dive" field:"更新"`
	Delete []uint                         `json:"delete" binding:"required,dive" field:"削除"`
}

type UpdateExpeditionImagesRequest struct {
	Add    []string `json:"add" binding:"required,dive" field:"追加"`
	Delete []string `json:"delete" binding:"required,dive" field:"削除"`
}

type UpdateExpeditionRequestBody struct {
	SportId           uint                           `json:"sportId" binding:"required" example:"1" field:"スポーツID"`
	IsPublic          bool                           `json:"isPublic" example:"true" field:"公開設定"`
	Title             string                         `json:"title" binding:"required" example:"野球観戦の遠征記録" field:"タイトル"`
	StartDate         time.Time                      `json:"startDate" binding:"required" example:"2025-01-01T00:00:00Z" field:"開始日"`
	EndDate           time.Time                      `json:"endDate" binding:"required" example:"2025-01-01T00:00:00Z" field:"終了日"`
	StadiumId         uint                           `json:"stadiumId" binding:"required" example:"1" field:"スタジアムID"`
	Memo              string                         `json:"memo" binding:"required" example:"初めてのスタジアム訪問。とても楽しかった！" field:"メモ"`
	Payments          UpdatePaymentsRequest          `json:"payments" binding:"required" field:"支払い"`
	Games             UpdateGamesRequest             `json:"games" binding:"required" field:"試合"`
	VisitedFacilities UpdateVisitedFacilitiesRequest `json:"visitedFacilities" binding:"required" field:"周辺施設"`
	Images            UpdateExpeditionImagesRequest  `json:"images" binding:"required" field:"画像URL配列"`
}
type UpdateExpeditionRequestPath struct {
	ExpeditionId uint `uri:"expeditionId" binding:"required" example:"1" field:"遠征記録ID"`
}

type LikeExpeditionRequestPath struct {
	ExpeditionId uint `uri:"expeditionId" binding:"required" example:"1" field:"遠征記録ID"`
}

type UnlikeExpeditionRequestPath struct {
	ExpeditionId uint `uri:"expeditionId" binding:"required" example:"1" field:"遠征記録ID"`
}

type DeleteExpeditionRequestPath struct {
	ExpeditionId uint `uri:"expeditionId" binding:"required" example:"1" field:"遠征記録ID"`
}

type ExpeditionListResponse struct {
	ID          uint                   `json:"id" example:"1"`
	IsPublic bool `json:"isPublic" example:"true"`
	Title       string                 `json:"title" example:"野球観戦の遠征記録"`
	StartDate   time.Time             `json:"startDate" example:"2025-01-01T00:00:00Z"`
	EndDate     time.Time             `json:"endDate" example:"2025-01-01T00:00:00Z"`
	SportId     uint                   `json:"sportId" example:"1"`
	SportName   string                 `json:"sportName" example:"野球"`
	UserID      uint                   `json:"userId" example:"1"`
	UserName    string                 `json:"userName" example:"user123"`
	UserIcon    string                 `json:"userIcon" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
	StadiumId uint      `json:"stadiumId" example:"1"`
	StadiumName string `json:"stadiumName" example:"東京ドーム"`
	IsLiked bool `json:"isLiked" example:"true"`
	LikesCount  int64                  `json:"likesCount" example:"10"`
	Images      []string               `json:"images" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
	Team1Name   string                 `json:"team1Name" example:"ヤクルト"`
	Team2Name   string                 `json:"team2Name" example:"ソフトバンク"`
}

type GetExpeditionListRequestQuery struct {
	Page     int        `form:"page" binding:"required,min=1" example:"1" field:"ページ番号"`
	SportId  *uint      `form:"sportId" example:"1" field:"スポーツID"`
	TeamId   *uint      `form:"teamId" example:"1" field:"チームID"`
	StadiumId *uint `form:"stadiumId" example:"1" field:"スタジアムID"`
	UserId *uint `form:"userId" example:"1" field:"ユーザーID"`
}

type GetExpeditionListByUserIdRequestQuery struct {
	Page     int        `form:"page" binding:"required,min=1" example:"1" field:"ページ番号"`
	UserId     uint        `form:"userId" binding:"required" example:"1" field:"ユーザーID"`
}

type GetExpeditionDetailResponse struct {
	ExpeditionResponse
	Username        string `json:"username" example:"user123"`
	UserIcon    string          `json:"userIcon" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
	IsLiked bool `json:"isLiked" example:"true"`
	LikesCount  int64                  `json:"likesCount" example:"10"`
	VisitedFacilities []VisitedFacilityResponse `json:"visitedFacilities"`
	Payments          []PaymentResponse         `json:"payments"`
	ExpeditionImages  []ExpeditionImageResponse `json:"expeditionImages"`
	Games             []GameResponse            `json:"games"`
}

type LikeExpeditionResponse struct {
	LikesCount  int64                  `json:"likesCount" example:"10"`
	IsLiked bool `json:"isLiked" example:"true"`
}
type UnLikeExpeditionResponse struct {
	LikesCount  int64                  `json:"likesCount" example:"10"`
}