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

/*  ==============================
===== スタジアム情報関連 =======
============================== */
// スタジアム検索(id)
func (s *AdminToolService) stadiumSearchId(id uint) (*models.Stadium, error) {
	var stadium models.Stadium
	if err := db.DB.First(&stadium, id).Error; err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &stadium, nil
}

// スタジアム重複検索（条件：競技場名、住所）
// ※スタジアム情報追加時
func (s *AdminToolService) stadiumAddCheck(address string) (*models.Stadium, error) {
	stadium := models.Stadium{}
	if err := db.DB.Select("id", "name", "description", "address", "capacity", "description").Where("address = ?", address).First(&stadium).Error; err != nil {
		return nil, err
	}
	return &stadium, nil
}

// ※スタジアム情報更新時
func (s *AdminToolService) stadiumUpdateCheck(id uint, name, address string) error {
	stadium := models.Stadium{}
	log.Println("モデル生成直後")
	if err := db.DB.Select("id", "name").Where("id != ?", id).Where("name = ? OR address = ?", name, address).First(&stadium).Error; err != nil {
		log.Println("SQL実行直後")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// レコードが見つからなかった場合の処理
			log.Println("該当するレコードがないため処理を継続")
			return nil
		}
		log.Printf("エラー： %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "内部エラーが発生しました。")
	}
	return utils.NewCustomError(http.StatusInternalServerError, "更新データが他の登録済みデータの競技場名、住所のどちらか、もしくは両方が重複しています。")
}

// スタジアム情報の取得
func (s *AdminToolService) getStadiumsService(keyword string) ([]models.Stadium, error) {
	stadiums, err := adminToolService.stadiumSearchKeyword(keyword)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました。")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewCustomError(http.StatusUnauthorized, "検索結果がありませんでした。")
	}

	return stadiums, nil
}

// スタジアム検索
func (s *AdminToolService) stadiumSearchKeyword(keyword string) ([]models.Stadium, error) {
	stadiums := []models.Stadium{}

	if keyword != "" {
		if err := db.DB.Where("name LIKE ? OR address LIKE ?", "%"+keyword+"%", "%"+keyword+"%").First(&stadiums).Error; err != nil {
			return nil, err
		}

		if err := db.DB.Where("name LIKE ? OR address LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&stadiums).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.DB.Find(&stadiums).Error; err != nil {
			return nil, err
		}
	}
	return stadiums, nil
}

// スタジアム追加
func (s *AdminToolService) createStadiumService(request *StadiumAddRequest) error {
	stadium, err := adminToolService.stadiumAddCheck(request.Address)

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
	err := adminToolService.stadiumUpdateCheck(request.StadiumId, request.Name, request.Address)
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
	stadium, err := adminToolService.stadiumSearchId(request.Id)
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

// スポーツ情報
// スポーツ追加
func (s *AdminToolService) createSportService(request *SportsAddRequest) error {
	sports, err := adminToolService.sportsAddCheck(request.Name)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました。")
	}
	if sports != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "登録済みのスタジアムです")
	}

	newSport := models.Sport{Name: request.Name}

	if err := db.DB.Create(&newSport).Error; err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "内部エラーが発生しました。")
	}
	return nil
}

// スポーツ情報更新
func (s *AdminToolService) UpdateSportService(request *SportsUpdateRequest) error {
	err := adminToolService.sportUppdateCheck(request.SportsId, request.Name)
	if err != nil {
		return err
	}
	log.Println(request.SportsId, request.Name)
	updateSport := models.Sport{Name: request.Name}

	if err := db.DB.Model(&models.Sport{}).Where("id = ?", request.SportsId).Updates(updateSport).Error; err != nil {
		log.Println("エラー", err)
		return utils.NewCustomError(http.StatusInternalServerError, "レコードが更新されませんでした")
	}
	log.Println("SQLは成功したよ")
	return nil
}

// スポーツ削除
func (s *AdminToolService) deleteSportService(request *DeleteRequest) error {
	sport, err := adminToolService.sportSearchId(request.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.NewCustomError(http.StatusUnauthorized, "スポーツが見つかりませんでした")
		}
		return utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました")
	}

	if err := db.DB.Delete(sport).Error; err != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "削除に失敗しました")
	}
	return nil
}

// スタジアム情報の取得
func (s *AdminToolService) getSportService(keyword string) ([]models.Sport, error) {
	sport, err := adminToolService.sportSearchKeyword(keyword)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました。")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewCustomError(http.StatusUnauthorized, "検索結果がありませんでした。")
	}

	return sport, nil
}

// スポーツ情報重複検索（条件：競技場名）
// ※スポーツ情報追加時
func (s *AdminToolService) sportsAddCheck(name string) (*models.Sport, error) {
	sports := models.Sport{}
	if err := db.DB.Select("id", "name").Where("name = ?", name).First(&sports).Error; err != nil {
		return nil, err
	}
	return &sports, nil
}

// 　※スポーツ情報更新時
func (s *AdminToolService) sportUppdateCheck(id uint, name string) error {
	sport := models.Sport{}
	log.Println("モデル生成直後")
	if err := db.DB.Select("id", "name").Where("id != ?", id).Where("name = ?", name).First(&sport).Error; err != nil {
		log.Println("SQL実行直後")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// レコードが見つからなかった場合の処理
			log.Println("該当するレコードがないため処理を継続")
			return nil
		}
		log.Printf("エラー： %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "内部エラーが発生しました。")
	}
	return utils.NewCustomError(http.StatusInternalServerError, "更新データが他の登録済みデータと重複しています。")
}

// スポーツ検索(id)
func (s *AdminToolService) sportSearchId(id uint) (*models.Sport, error) {
	var sport models.Sport
	if err := db.DB.First(&sport, id).Error; err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &sport, nil
}

// スポーツ検索
func (s *AdminToolService) sportSearchKeyword(keyword string) ([]models.Sport, error) {
	sport := []models.Sport{}

	if keyword != "" {
		if err := db.DB.Where("name LIKE ?", "%"+keyword+"%").First(&sport).Error; err != nil {
			return nil, err
		}

		if err := db.DB.Where("name LIKE ? ", "%"+keyword+"%").Find(&sport).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.DB.Find(&sport).Error; err != nil {
			return nil, err
		}
	}
	return sport, nil
}

func NewAdminToolService() *AdminToolService {
	return &AdminToolService{}
}
