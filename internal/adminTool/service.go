package adminTool

import (
	"go-docker/internal/db"
	"go-docker/models"
	"log"
)

type AdminToolService struct{}

// func (s *AdminToolService) test() {
// }

func (s *AdminToolService) teamSearch(teamName string) (*models.Team, error) {
	team := models.Team{}
	if err := db.DB.Select("name", "league_id", "sport_id").Where("name = ?", teamName).First(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (s *AdminToolService) createTeam(newTeam *models.Team) error {
	if err := db.DB.Create(&newTeam).Error; err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

func NewAdminToolService() *AdminToolService {
	return &AdminToolService{}
}
