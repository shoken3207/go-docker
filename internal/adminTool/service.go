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

// スタジアム情報関連
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
func (s *AdminToolService) UpdateStadiumService(request *StadiumUpdateRequest) error {
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

// リーグ情報
// リーグ情報追加
func (s *AdminToolService) createLeagueService(request *LeagueAddRequest) error {
	sports, err := adminToolService.LeagueAddCheck(request.Name, request.SportsId)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました。")
	}
	if sports != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "登録済みのリーグです")
	}

	newLeague := models.League{Name: request.Name, SportId: request.SportsId}

	if err := db.DB.Create(&newLeague).Error; err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "内部エラーが発生しました。")
	}

	return nil
}

// リーグ情報更新
func (s *AdminToolService) UpdateLeagueService(request *LeagueUpdateRequest) error {
	err := adminToolService.LeagueUpdateCheck(request.LeagueId, request.SportsId, request.Name)
	if err != nil {
		return err
	}
	log.Println(request.LeagueId, request.SportsId, request.Name)
	updateLeague := models.League{Name: request.Name, SportId: request.SportsId}

	if err := db.DB.Model(&models.League{}).Where("id = ?", request.LeagueId).Updates(updateLeague).Error; err != nil {
		log.Println("エラー", err)
		return utils.NewCustomError(http.StatusInternalServerError, "レコードが更新されませんでした")
	}
	log.Println("SQLは成功したよ")
	return nil
}

// リーグ情報削除
func (s *AdminToolService) deleteLeagueService(request *DeleteRequest) error {
	league, err := adminToolService.LeagueSearchId(request.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.NewCustomError(http.StatusUnauthorized, "リーグが見つかりませんでした")
		}
		return utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました")
	}

	if err := db.DB.Delete(league).Error; err != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "削除に失敗しました")
	}
	return nil
}

// リーグ情報の取得
func (s *AdminToolService) getLeagueService(keyword string) ([]models.League, error) {
	league, err := adminToolService.leagueSearchKeyword(keyword)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました。")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewCustomError(http.StatusUnauthorized, "検索結果がありませんでした。")
	}

	return league, nil
}

// リーグ情報重複検索（条件：競技場名）
// ※リーグ情報追加時
func (s *AdminToolService) LeagueAddCheck(name string, sport_id uint) (*models.League, error) {
	league := models.League{}
	if err := db.DB.Select("id", "name", "sport_id").Where("name = ? OR sport_id = ?", name, sport_id).First(&league).Error; err != nil {
		log.Println("エラー:", err)
		return nil, err
	}
	return &league, nil
}

// 　※リーグ情報更新時
func (s *AdminToolService) LeagueUpdateCheck(id, sport_id uint, name string) error {
	league := models.League{}
	log.Println("モデル生成直後")
	if err := db.DB.Select("id", "name", "sport_id").Where("id != ?", id).Where("name = ? OR sport_id = ?", name, sport_id).First(&league).Error; err != nil {
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

// リーグ検索(id)
func (s *AdminToolService) LeagueSearchId(id uint) (*models.League, error) {
	var league models.League
	if err := db.DB.First(&league, id).Error; err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &league, nil
}

// リーグ検索
func (s *AdminToolService) leagueSearchKeyword(keyword string) ([]models.League, error) {
	league := []models.League{}
	log.Println(keyword)
	if keyword != "" {
		if err := db.DB.Where("name LIKE ?", "%"+keyword+"%").First(&league).Error; err != nil {
			return nil, err
		}

		if err := db.DB.Where("name LIKE ? ", "%"+keyword+"%").Find(&league).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.DB.Find(&league).Error; err != nil {
			return nil, err
		}
	}
	return league, nil
}

// チーム情報関連
// チーム情報追加
func (s *AdminToolService) createTeamService(request *TeamAddRequest) error {
	sports, err := adminToolService.teamAddCheck(request.Name, request.SportsId)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました。")
	}
	if sports != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "登録済みのリーグです")
	}

	newLeague := models.Team{StadiumId: request.StadiumId, SportId: request.SportsId, LeagueId: request.LeagueId, Name: request.Name}

	if err := db.DB.Create(&newLeague).Error; err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "内部エラーが発生しました。")
	}

	return nil
}

// チーム情報更新
func (s *AdminToolService) UpdateTeamService(request *TeamUpdateRequest) error {
	err := adminToolService.teamUpdateCheck(request.TeamId, request.SportsId, request.Name)
	if err != nil {
		return err
	}
	log.Println(request.TeamId, request.StadiumId, request.LeagueId, request.SportsId, request.Name)
	updateTeam := models.Team{Name: request.Name, StadiumId: request.StadiumId, LeagueId: request.LeagueId, SportId: request.SportsId}

	if err := db.DB.Model(&models.Team{}).Where("id = ?", request.TeamId).Updates(updateTeam).Error; err != nil {
		log.Println("エラー", err)
		return utils.NewCustomError(http.StatusInternalServerError, "レコードが更新されませんでした")
	}
	log.Println("SQLは成功したよ")
	return nil
}

// チーム情報削除
func (s *AdminToolService) deleteTeamService(request *DeleteRequest) error {
	team, err := adminToolService.teamSearchId(request.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.NewCustomError(http.StatusUnauthorized, "リーグが見つかりませんでした")
		}
		return utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました")
	}

	if err := db.DB.Delete(team).Error; err != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "削除に失敗しました")
	}
	return nil
}

// チーム情報の取得
func (s *AdminToolService) getTeamService(keyword string) ([]models.Team, error) {
	team, err := adminToolService.teamSearchKeyword(keyword)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました。")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewCustomError(http.StatusUnauthorized, "検索結果がありませんでした。")
	}

	return team, nil
}

// チーム情報重複検索（条件：競技場名）
// ※チーム情報追加時
func (s *AdminToolService) teamAddCheck(name string, sport_id uint) (*models.Team, error) {
	team := models.Team{}
	if err := db.DB.Select("id", "name", "stadium_id", "league_id", "sport_id").Where("name = ? OR sport_id = ?", name, sport_id).First(&team).Error; err != nil {
		log.Println("エラー:", err)
		return nil, err
	}
	return &team, nil
}

// 　※チーム情報更新時
func (s *AdminToolService) teamUpdateCheck(id, sport_id uint, name string) error {
	league := models.Team{}
	log.Println("モデル生成直後")
	if err := db.DB.Select("id", "stadium_id", "league_id", "sport_id", "name").Where("id != ?", id).Where("name = ? OR sport_id = ?", name, sport_id).First(&league).Error; err != nil {
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

// チーム検索(id)
func (s *AdminToolService) teamSearchId(id uint) (*models.Team, error) {
	var team models.Team
	if err := db.DB.First(&team, id).Error; err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &team, nil
}

// チーム検索
func (s *AdminToolService) teamSearchKeyword(keyword string) ([]models.Team, error) {
	team := []models.Team{}
	log.Println(keyword)
	if keyword != "" {
		if err := db.DB.Where("name LIKE ?", "%"+keyword+"%").First(&team).Error; err != nil {
			return nil, err
		}

		if err := db.DB.Where("name LIKE ? ", "%"+keyword+"%").Find(&team).Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.DB.Find(&team).Error; err != nil {
			return nil, err
		}
	}
	return team, nil
}

func NewAdminToolService() *AdminToolService {
	return &AdminToolService{}
}
