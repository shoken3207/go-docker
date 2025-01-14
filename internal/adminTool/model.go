package adminTool

// リクエスト
// スタジアム関連
type StadiumAddRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Capacity    uint   `json:"capacity" binding:"required"`
	Image       string `json:"image" binding:"required"`
}

type StadiumUpdateRequest struct {
	StadiumId   uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Capacity    uint   `json:"capacity" binding:"required"`
	Image       string `json:"image" binding:"required"`
}

// スポーツ情報
type SportsAddRequest struct {
	Name string `json:"name" binding:"required"`
}

type SportsUpdateRequest struct {
	SportsId uint   `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

// リーグ情報
type LeagueAddRequest struct {
	Name     string `json:"name" bindning:"required"`
	SportsId uint   `json:"sport_id" binding:"required"`
}

type LeagueUpdateRequest struct {
	LeagueId uint   `json:"id" binding:"required"`
	Name     string `json:"name" bindning:"required"`
	SportsId uint   `json:"sport_id" binding:"required"`
}

// チーム関連
type TeamAddRequest struct {
	StadiumId uint   `json:"stadiumId" binding:"required"`
	SportsId  uint   `json:"sportsId" binding:"required"`
	LeagueId  uint   `json:"LeagueId" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

type TeamUpdateRequest struct {
	TeamId    uint   `json:"teamId" binding:"required"`
	StadiumId uint   `json:"stadiumId" binding:"required"`
	SportsId  uint   `json:"sportsId" binding:"required"`
	LeagueId  uint   `json:"LeagueId" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

// 削除は共通してidを利用するため共通のmodelを利用する
type DeleteRequest struct {
	Id uint `uri:"id" binding:"required"`
}

// レスポンス
// スタジアム情報
type Stadium struct {
	StadiumId   uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Capacity    uint   `json:"capacity" binding:"required"`
	Image       string `json:"image" binding:"required"`
}

// スポーツ情報
type Sports struct {
	SportsId uint   `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

// リーグ情報
type League struct {
	LeagueId uint   `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	SportsId uint   `json:"sport_id" binding:"required"`
}

// チーム情報
type Team struct {
	TeamId    uint   `json:"id"`
	StadiumId uint   `json:"stadiumId"`
	SportsId  uint   `json:"sportsId"`
	LeagueId  uint   `json:"LeagueId"`
	TeamName  string `json:"teamName"`
}
