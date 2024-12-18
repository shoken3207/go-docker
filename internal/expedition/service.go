package expedition

import (
	"go-docker/internal/db"
	"go-docker/models"
	"go-docker/pkg/utils"
	"log"
	"net/http"
)

type ExpeditionService struct{}

func (s *ExpeditionService) CreateExpedition(newExpedition *models.Expedition) error {
	if err := db.DB.Create(newExpedition); err != nil {
		log.Printf("遠征記録作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録作成に失敗しました。")
	}
	return nil
}

func (s *ExpeditionService) CreateExpeditionService(request *CreateExpeditionRequest) error {
	newExpedition := models.Expedition{
		SportId:   request.SportId,
		IsPublic:  request.IsPublic,
		Title:     request.Title,
		Memo:      request.Memo,
		StartDate: request.StartDate,
		EndDate:   request.EndDate,
	}
	if err := s.CreateExpedition(&newExpedition); err != nil {
		return err
	}

	return nil
}

func NewExpeditionService() *ExpeditionService {
	return &ExpeditionService{}
}
