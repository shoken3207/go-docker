package adminTool

import (
	"errors"
	"go-docker/internal/db"
	"go-docker/models"
	"go-docker/pkg/utils"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type AdminToolService struct{}

// func (s *AdminToolService) test() {
// }

// チーム情報関連
// チーム重複検索(条件：チーム名)
// func (s *AdminToolService) teamCheck(teamName string) (*models.Team, error) {
// 	team := models.Team{}
// 	if err := db.DB.Select("name", "league_id", "sport_id").Where("name = ?", teamName).First(&team).Error; err != nil {
// 		log.Printf("ユーザーデータ取得エラー: %v", err)
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, utils.NewCustomError(http.StatusNotFound, "該当チームが見つかりませんでした。")
// 		} else {
// 			return nil, utils.NewCustomError(http.StatusInternalServerError, "ユーザーデータ取得に失敗しました。")
// 		}
// 	}
// 	return nil
// }

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
	if err := db.DB.Select("id", "name", "description", "address", "capacity", "description").Where("address = ?", address).First(&stadium).Error; err != nil {
		return nil, err
	}
	return &stadium, nil
}

// ※スタジアム情報更新時
func (s *AdminToolService) stadiumUppdateCheck(id uint, name, address string) error {
	stadium := models.Stadium{}
	log.Println("モデル生成直後")
	if err := db.DB.Select("id", "name").Where("id != ?", id).Where("name = ? OR address = ?", name, address).First(&stadium).Error; err != nil {
		log.Println("SQL実行直後")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// レコードが見つからなかった場合の処理
			log.Println("gorm.ErrRecordNotFround処理を検知")
			return nil
		}
		log.Printf("ここがエラー!!!: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "内部エラーが発生しました。")
	}
	return utils.NewCustomError(http.StatusInternalServerError, "更新データが他の登録済みデータの競技場名、住所のどちらか、もしくは両方が重複しています。")
}

// スタジアム追加
func (s *AdminToolService) createStadiumService(request *StadiumAddRequest) error {
	stadium, err := adminToolService.stadiumAddCheck(request.Name, request.Address)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました。")
	}
	if stadium != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "登録済みのスタジアムです")
	}

	newStadium := models.Stadium{Name: request.Name, Description: request.Description, Address: request.Address, Capacity: int(request.Capacity), Image: request.Image}

	if err := db.DB.Create(&newStadium).Error; err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "内部エラーが発生しました。")
	}
	return nil
}

// スタジアム更新
func (s *AdminToolService) UpdateStadiumService(request *StadiumUppdateRequest) error {
	err := adminToolService.stadiumUppdateCheck(request.StadiumId, request.Name, request.Address)
	if err != nil {
		return err
	}
	log.Println(request.Name, request.Description, request.Address, int(request.Capacity), request.Image)
	updateStadium := models.Stadium{Name: request.Name, Description: request.Description, Address: request.Address, Capacity: int(request.Capacity), Image: request.Image}

	if err := db.DB.Model(&models.Stadium{}).Where("id = ?", request.StadiumId).Updates(updateStadium).Error; err != nil {
		log.Println("エラー", err)
		return utils.NewCustomError(http.StatusInternalServerError, "レコードが更新されませんでした")
	}
	log.Println("SQLは成功したよ")
	return nil
}

// スタジアム削除
func (s *AdminToolService) deleteStadiumService(request *DeleteRequest) error {
	stadium, err := adminToolService.stadiumSearch(request.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.NewCustomError(http.StatusUnauthorized, "スタジアムが見つかりませんでした")
		}
		return utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました")
	}

	if err := db.DB.Delete(stadium).Error; err != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "削除に失敗しました")
	}
	return nil
}

func NewAdminToolService() *AdminToolService {
	return &AdminToolService{}
}
