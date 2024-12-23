package expedition

import (
	"go-docker/internal/db"
	"go-docker/models"
	"go-docker/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpeditionService struct{}

func (s *ExpeditionService) CreateExpedition(newExpedition *models.Expedition) error {
	if err := db.DB.Create(newExpedition).Error; err != nil {
		log.Printf("遠征記録作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録作成に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) UpdateExpedition(updateExpedition *models.Expedition) error {
	if err := db.DB.Save(updateExpedition).Error; err != nil {
		log.Printf("遠征記録更新エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録更新に失敗しました。")
	}
	return nil
}

func (s *ExpeditionService) CreatePayment(newPayment *models.Payment) error {
	if err := db.DB.Create(newPayment).Error; err != nil {
		log.Printf("遠征記録出費作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録出費作成に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) UpdatePayment(updatePayment *models.Payment) error {
	if err := db.DB.Save(updatePayment).Error; err != nil {
		log.Printf("遠征記録出費更新エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録出費更新に失敗しました。")
	}
	return nil
}

func (s *ExpeditionService) CreateExpeditionImage(newExpeditionImage *models.ExpeditionImage) error {
	if err := db.DB.Create(newExpeditionImage).Error; err != nil {
		log.Printf("遠征記録画像データ作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録画像データ作成に失敗しました。")
	}
	return nil
}

func (s *ExpeditionService) CreateVisitedFacility(newVisitedFacilitiy *models.VisitedFacility) error {
	if err := db.DB.Create(newVisitedFacilitiy).Error; err != nil {
		log.Printf("訪れた施設作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "訪れた施設作成に失敗しました。")
	}
	return nil
}

func (s *ExpeditionService) CreateGame(newGame *models.Game) error {
	if err := db.DB.Create(newGame).Error; err != nil {
		log.Printf("試合記録作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "試合記録作成に失敗しました。")
	}
	return nil
}

func (s *ExpeditionService) CreateGameScore(newGameScore *models.GameScore) error {
	if err := db.DB.Create(newGameScore).Error; err != nil {
		log.Printf("試合得点作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "試合得点作成に失敗しました。")
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
		StadiumId: request.StadiumId,
	}
	if err := s.CreateExpedition(&newExpedition); err != nil {
		return err
	}

	games := request.Games
	for _, game := range games {
		newGame := models.Game{
			ExpeditionId: newExpedition.ID,
			Date:         game.Date,
			Team1Id:      game.Team1Id,
			Team2Id:      game.Team2Id,
			Comment:      game.Comment,
		}
		if err := s.CreateGame(&newGame); err != nil {
			return err
		}

		for _, gameScore := range game.Scores {
			newGameScore := models.GameScore{
				GameId: newGame.ID,
				TeamId: gameScore.TeamId,
				Score:  gameScore.Score,
				Order:  gameScore.Order,
			}

			if err := s.CreateGameScore(&newGameScore); err != nil {
				return err
			}
		}
	}

	images := request.Images
	for _, image := range images {
		newExpeditionImage := models.ExpeditionImage{
			ExpeditionId: newExpedition.ID,
			Image:        image.Image,
			FileId:       image.FileId,
		}
		if err := s.CreateExpeditionImage(&newExpeditionImage); err != nil {
			return err
		}
	}

	payments := request.Payments
	for _, payment := range payments {
		newPayment := models.Payment{
			Title:        payment.Title,
			Date:         payment.Date,
			Cost:         payment.Cost,
			ExpeditionId: newExpedition.ID,
		}
		if err := s.CreatePayment(&newPayment); err != nil {
			return err
		}
	}

	visitedFacilities := request.VisitedFacilities
	for _, visitedFacility := range visitedFacilities {
		newVisitedFacility := models.VisitedFacility{
			ExpeditionId: newExpedition.ID,
			Name:         visitedFacility.Name,
			Address:      visitedFacility.Address,
			Icon:         visitedFacility.Icon,
			Color:        visitedFacility.Color,
			Latitude:     visitedFacility.Latitude,
			Longitude:    visitedFacility.Longitude,
		}
		if err := s.CreateVisitedFacility(&newVisitedFacility); err != nil {
			return err
		}
	}

	return nil
}

func (s *ExpeditionService) ValidateUpdateExpeditionRequest(c *gin.Context) (*uint, *UpdateExpeditionRequestBody, error) {
	var requestBody UpdateExpeditionRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Printf("リクエストエラー: %v", err)
		return nil, nil, utils.NewCustomError(http.StatusBadRequest, "リクエストに不備があります。")
	}
	var requestPath UpdateExpeditionRequestPath
	if err := c.ShouldBindUri(&requestPath); err != nil {
		log.Printf("リクエストエラー: %v", err)
		return nil, nil, utils.NewCustomError(http.StatusBadRequest, "リクエストに不備があります。")
	}

	return &requestPath.ExpeditionId, &requestBody, nil
}

func (s *ExpeditionService) UpdateExpeditionService(expeditionId *uint, requestBody *UpdateExpeditionRequestBody) error {
	return nil
}

func NewExpeditionService() *ExpeditionService {
	return &ExpeditionService{}
}
