package stadium

import "go-docker/internal/expedition"

type FacilityResponse struct {
	Id uint `json:"id"`
	Name         string     `json:"name" `
	CustomName string `json:"customName"`
	Address      string     `json:"address"`
	VisitCount int `json:"visitCount"`
}

type GetStadiumResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Capacity    int    `json:"capacity"`
	Image       string `json:"image"`
	Expeditions []expedition.ExpeditionListResponse `json:"expeditions"`
	Facilities [] FacilityResponse `json:"facilities"`
}