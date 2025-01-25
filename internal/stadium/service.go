package stadium

import (
	"errors"
	"go-docker/internal/db"
	"go-docker/internal/expedition"
	"go-docker/models"
	"go-docker/pkg/utils"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type StadiumService struct{}
var expeditionService = expedition.NewExpeditionService()

func (s *StadiumService) GetStadium(stadiumId *uint) (*models.Stadium, error) {
	var stadium models.Stadium
	if err := db.DB.First(&stadium, *stadiumId).Error; err != nil {
		log.Printf("スタジアム取得エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusUnauthorized, "スタジアムが見つかりません。")
		}
		return nil, utils.NewCustomError(http.StatusInternalServerError, "スタジアム取得に失敗しました。")
	}

	return &stadium, nil
}

func (s *StadiumService) GetStadiumService(loginUserId *uint, request *GetStadiumRequestPath) (*GetStadiumResponse, error) {
	stadium, err := s.GetStadium(&request.StadiumId)
	if err != nil {
		return nil, err
	}

	var facilities []FacilityResponse
	query := `
		SELECT
			name,
			address,
			COUNT(*) AS visit_count
		FROM
			visited_facilities
		WHERE
			expedition_id IN (
				SELECT id
				FROM expeditions
				WHERE stadium_id = ?
			)
		GROUP BY name, address
		ORDER BY visit_count DESC
		LIMIT 20
	`

	if err := db.DB.Raw(query, request.StadiumId).Scan(&facilities).Error; err != nil {
		log.Printf("施設取得エラー: %v", err)
		return nil, utils.NewCustomError(http.StatusInternalServerError, "施設取得に失敗しました。")
	}
	expeditions, err := expeditionService.GetExpeditionListService(&expedition.GetExpeditionListRequestQuery{Page: 1,StadiumId: &request.StadiumId}, loginUserId)
	if err != nil {
		return nil, err
	}

	stadiumResponse := GetStadiumResponse{
		Id: stadium.ID,
		Name: stadium.Name,
		Description: stadium.Description,
		Address: stadium.Address,
		Capacity: stadium.Capacity,
		Image: stadium.Image,
		Expeditions: expeditions,
		Facilities: facilities,
	}
	return &stadiumResponse, nil
}

func NewStadiumService() *StadiumService {
	return &StadiumService{}
}