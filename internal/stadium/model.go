package stadium

import "go-docker/internal/expedition"

type GetStadiumRequestPath struct {
	StadiumId uint `uri:"stadiumId" binding:"required" example:"1" field:"スタジアムID"`
}

type FacilityResponse struct {
	Name         string     `json:"name" example:"東京駅"`
	Address      string     `json:"address" example:"東京都千代田区丸の内1-1-1"`
	VisitCount int `json:"visitCount" example:"1"`
}

type GetStadiumResponse struct {
	Id          uint   `json:"id" example:"1"`
	Name        string `json:"name" example:"京セラドーム"`
	Description string `json:"description" example:"オリックス・バファローズのホーム球場"`
	Address     string `json:"address" example:"大阪府大阪市西区千代崎3-1-1"`
	Capacity    int    `json:"capacity" example:"36000"`
	Image       string `json:"image" example:"https://ik.imagekit.io/your_imagekit_id/image.jpg"`
	Attribution string `json:"attribution" example:"attribution"`
	Expeditions []expedition.ExpeditionListResponse `json:"expeditions"`
	Facilities [] FacilityResponse `json:"facilities"`
}