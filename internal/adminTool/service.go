package adminTool

import (
	"go-docker/internal/db"
	"go-docker/models"
	"log"
)

type AdminToolService struct{}

// func (s *AdminToolService) test() {
// }

// チーム情報関連
// チーム重複検索(条件：チーム名)
func (s *AdminToolService) teamCheck(teamName string) (*models.Team, error) {
	team := models.Team{}
	if err := db.DB.Select("name", "league_id", "sport_id").Where("name = ?", teamName).First(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

// 新規チーム追加
func (s *AdminToolService) createTeam(newTeam *models.Team) error {
	if err := db.DB.Create(&newTeam).Error; err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

// スタジアム情報関連
// スタジアム検索(id)
func (s *AdminToolService) stadiumSearch(id uint) (*models.Stadium, error) {
	var stadium models.Stadium
	if err := db.DB.First(&stadium, id).Error; err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &stadium, nil
}

// スタジアム重複検索（条件：競技場名、住所）
// ※スタジアム情報追加時
func (s *AdminToolService) stadiumAddCheck(stadiumName, address string) (*models.Stadium, error) {
	stadium := models.Stadium{}
	if err := db.DB.Select("id", "name", "description", "address", "capacity", "description").Where("name = ?", stadiumName).Or("address = ?", address).First(&stadium).Error; err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &stadium, nil
}

// ※スタジアム情報更新時
func (s *AdminToolService) stadiumUppdateCheck(id uint, stadiumName, address string) (*models.Stadium, error) {
	stadium := models.Stadium{}
	if err := db.DB.Select("id", "name", "description", "address", "capacity", "description").Where("id != ?", id).Where("name = ?", stadiumName).Or("address = ?", address).First(&stadium).Error; err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &stadium, nil
}

// スタジアム追加
func (s *AdminToolService) createStadium(newStadium *models.Stadium) error {
	if err := db.DB.Create(&newStadium).Error; err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	return nil
}

// スタジアム更新
func (s *AdminToolService) UpdateStadium(id uint, updatedStadium *models.Stadium) error {
	result := db.DB.Model(&models.Stadium{}).Where("id = ?", id).Updates(updatedStadium)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// スタジアム削除
func (s *AdminToolService) deleteStadium(stadium *models.Stadium) error {
	if err := db.DB.Delete(stadium).Error; err != nil {
		return err
	}
	return nil
}

func NewAdminToolService() *AdminToolService {
	return &AdminToolService{}
}
