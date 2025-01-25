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
	if err := db.DB.Select("id", "name", "description", "address", "capacity", "file_id").Where("address = ?", address).First(&stadium).Error; err != nil {
		return nil, err
	}
	return &stadium, nil
}

// fieldIdの確認及び値の設定
func (s *AdminToolService) stadiumFileIdCheck(fileId string) string {
	if fileId == "NoImage" {
		fileId = "https://blogger.googleusercontent.com/img/b/R29vZ2xl/AVvXsEhhNiQiwJEndtEtiXGbge_nFBhm48O1veQDVkskf53TwtD9Tf-UsueCE7WkNoLrs3cn05HT07yCtpNkFH8UcmEP4-IA-POvT81HlnsWRnOiCrJQ_MF8lRQxmUURmwhRMJdffXm_RRPXzjZO/s1600/no_image_yoko.jpg"
	}
	return fileId
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
func (s *AdminToolService) getStadiumsService(keyword string) ([]Stadium, error) {
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
func (s *AdminToolService) stadiumSearchKeyword(keyword string) ([]Stadium, error) {
	var stadiums []models.Stadium
	log.Println("キーワード:", keyword)
	query := db.DB.Model(&models.Stadium{})

	if keyword != "" {
		if err := query.Select("id", "name", "description", "address", "capacity", "image").Where("name LIKE ? OR address LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&stadiums).Error; err != nil {
			return nil, err
		}
	} else {
		if err := query.Find(&stadiums).Error; err != nil {
			return nil, err
		}
	}

	var stadiumResponse []Stadium
	for _, stadium := range stadiums {
		stadiumResponse = append(stadiumResponse, Stadium{
			StadiumId:   stadium.ID,
			Name:        stadium.Name,
			Description: stadium.Description,
			Address:     stadium.Address,
			Capacity:    uint(stadium.Capacity),
			Image:       stadium.Image,
			FileId:      stadium.FileId,
		})
	}
	return stadiumResponse, nil
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

	log.Println("リクエスト:", request)
	log.Println("リクエスト:", request.Image)

	return db.DB.Transaction(func(tx *gorm.DB) error {
		var fileId string
		if request.Image != "" {
			tempImages, err := utils.ValidateAndPersistImages(tx, []string{request.Image})
			if err != nil {
				return err
			}
			if len(tempImages) > 0 {
				fileId = tempImages[0].FileId
			}
		}

		newStadium := models.Stadium{
			Name:        request.Name,
			Description: request.Description,
			Address:     request.Address,
			Capacity:    int(request.Capacity),
			Image:       request.Image,
		}
		newStadium.SetFileId(fileId)

		fileId = adminToolService.stadiumFileIdCheck(fileId)
		log.Println("ファイルID:", fileId)

		if err := adminToolService.createStadium(tx, &newStadium); err != nil {
			return err
		}

		return nil
	})
}

func (s *AdminToolService) createStadium(tx *gorm.DB, newStadium *models.Stadium) error {
	if err := tx.Create(&newStadium).Error; err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "内部エラーが発生しました。")
	}
	return nil
}

// スタジアム更新
func (s *AdminToolService) UpdateStadiumService(id uint, request *StadiumUpdateRequest) error {
	err := adminToolService.stadiumUpdateCheck(id, request.Name, request.Address)
	if err != nil {
		return err
	}

	log.Println(request.Name, request.Description, request.Address, int(request.Capacity), request.Image)

	return db.DB.Transaction(func(tx *gorm.DB) error {
		var fileId string
		if request.Image != "" {
			tempImages, err := utils.ValidateAndPersistImages(tx, []string{request.Image})
			if err != nil {
				return err
			}
			if len(tempImages) > 0 {
				fileId = tempImages[0].FileId
			}
		}

		updateStadium := models.Stadium{
			Name:        request.Name,
			Description: request.Description,
			Address:     request.Address,
			Capacity:    int(request.Capacity),
			Image:       request.Image,
		}

		fileId = adminToolService.stadiumFileIdCheck(fileId)

		updateStadium.SetFileId(fileId)

		fileId = adminToolService.stadiumFileIdCheck(fileId)
		log.Println("ファイルID:", fileId)

		if err := adminToolService.updateStadium(tx, id, &updateStadium); err != nil {
			return err
		}

		return nil
	})
}

func (s *AdminToolService) updateStadium(tx *gorm.DB, id uint, updateStadium *models.Stadium) error {
	if err := tx.Model(&models.Stadium{}).Where("id = ?", id).Updates(updateStadium).Error; err != nil {
		log.Println("エラー", err)
		return utils.NewCustomError(http.StatusInternalServerError, "レコードが更新されませんでした")
	}
	return nil
}

// スタジアム削除
func (s *AdminToolService) deleteStadiumService(id uint) error {
	stadium, err := adminToolService.stadiumSearchId(id)
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

// 該当idからレコードを取得
func (s *AdminToolService) StadiumGetIdService(id uint) (*Stadium, error) {
	var stadium Stadium
	if err := db.DB.Select("id", "name", "description").Where("id = ?", id).Find(&stadium).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewCustomError(http.StatusUnauthorized, "スタジアムが見つかりませんでした")
		}
		return nil, utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました")
	}
	return &stadium, nil
}

// スポーツ情報
// スポーツ追加
func (s *AdminToolService) createSportService(request *SportsAddRequest) error {
	sports, err := adminToolService.sportsAddCheck(request.Name)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました。")
	}
	if sports != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "登録済みのスポーツです")
	}

	newSport := models.Sport{Name: request.Name}

	if err := db.DB.Create(&newSport).Error; err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "内部エラーが発生しました。")
	}
	return nil
}

// スポーツ情報更新
func (s *AdminToolService) UpdateSportService(id uint, request *SportsUpdateRequest) error {
	err := adminToolService.sportUppdateCheck(id, request.Name)
	if err != nil {
		return err
	}
	log.Println(id, request.Name)
	updateSport := models.Sport{Name: request.Name}

	if err := db.DB.Model(&models.Sport{}).Where("id = ?", id).Updates(updateSport).Error; err != nil {
		log.Println("エラー", err)
		return utils.NewCustomError(http.StatusInternalServerError, "レコードが更新されませんでした")
	}
	log.Println("SQLは成功したよ")
	return nil
}

// スポーツ削除
func (s *AdminToolService) deleteSportService(id uint) error {
	log.Println("受け取った値(重複検索前):", id)
	sport, err := adminToolService.sportSearchId(id)
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

// スポーツ情報の取得
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
	log.Println("受け取った値(id検索):", id)
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

// 該当idからレコードを取得
func (s *AdminToolService) SportGetIdService(id uint) (*Sports, error) {
	var sport Sports
	if err := db.DB.Select("id", "name").Where("id = ?", id).Find(&sport).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewCustomError(http.StatusUnauthorized, "スポーツが見つかりませんでした")
		}
		return nil, utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました")
	}
	return &sport, nil
}

// リーグ情報
// リーグ情報追加
func (s *AdminToolService) createLeagueService(request *LeagueAddRequest) error {
	league, err := adminToolService.LeagueAddCheck(request.Name, request.SportsId)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました。")
	}
	if league != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "登録済みのリーグです")
	}

	newLeague := models.League{Name: request.Name, SportId: request.SportsId}

	if err := db.DB.Create(&newLeague).Error; err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "内部エラーが発生しました。")
	}

	return nil
}

// リーグ情報更新
func (s *AdminToolService) UpdateLeagueService(id uint, request *LeagueUpdateRequest) error {
	err := adminToolService.LeagueUpdateCheck(id, request.SportsId, request.Name)
	if err != nil {
		return err
	}
	log.Println(id, request.SportsId, request.Name)
	updateLeague := models.League{Name: request.Name, SportId: request.SportsId}

	if err := db.DB.Model(&models.League{}).Where("id = ?", id).Updates(updateLeague).Error; err != nil {
		log.Println("エラー", err)
		return utils.NewCustomError(http.StatusInternalServerError, "レコードが更新されませんでした")
	}
	log.Println("SQLは成功したよ")
	return nil
}

// リーグ情報削除
func (s *AdminToolService) deleteLeagueService(id uint) error {
	league, err := adminToolService.LeagueSearchId(id)
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
	if err := db.DB.Select("id", "name", "sport_id").Where("name = ? ", name).First(&league).Error; err != nil {
		log.Println("エラー:", err)
		return nil, err
	}
	return &league, nil
}

// 　※リーグ情報更新時
func (s *AdminToolService) LeagueUpdateCheck(id, sport_id uint, name string) error {
	league := models.League{}
	log.Println("モデル生成直後")
	if err := db.DB.Select("id", "name", "sport_id").Where("id != ?", id).Where("name = ? AND sport_id = ?", name, sport_id).First(&league).Error; err != nil {
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

// 該当idからレコードを取得
func (s *AdminToolService) LeagueGetIdService(id uint) (*League, error) {
	var league League
	if err := db.DB.Select("id", "name", "sport_id").Where("id = ?", id).Find(&league).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewCustomError(http.StatusUnauthorized, "リーグが見つかりませんでした")
		}
		return nil, utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました")
	}

	log.Println("取得したレコード", &league)
	return &league, nil
}

// チーム情報関連
// チーム情報追加
func (s *AdminToolService) createTeamService(request *TeamAddRequest) error {
	sports, err := adminToolService.teamAddCheck(request.Name)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました。")
	}
	if sports != nil {
		return utils.NewCustomError(http.StatusUnauthorized, "登録済みのチームです")
	}

	newLeague := models.Team{StadiumId: request.StadiumId, SportId: request.SportsId, LeagueId: request.LeagueId, Name: request.Name}

	if err := db.DB.Create(&newLeague).Error; err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, "内部エラーが発生しました。")
	}

	return nil
}

// チーム情報更新
func (s *AdminToolService) UpdateTeamService(id uint, request *TeamUpdateRequest) error {
	err := adminToolService.teamUpdateCheck(id, request.SportsId, request.Name)
	if err != nil {
		return err
	}
	log.Println(id, request.StadiumId, request.LeagueId, request.SportsId, request.Name)
	updateTeam := models.Team{Name: request.Name, StadiumId: request.StadiumId, LeagueId: request.LeagueId, SportId: request.SportsId}

	if err := db.DB.Model(&models.Team{}).Where("id = ?", id).Updates(updateTeam).Error; err != nil {
		log.Println("エラー", err)
		return utils.NewCustomError(http.StatusInternalServerError, "レコードが更新されませんでした")
	}
	log.Println("SQLは成功したよ")
	return nil
}

// チーム情報削除
func (s *AdminToolService) deleteTeamService(id uint) error {
	team, err := adminToolService.teamSearchId(id)
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
func (s *AdminToolService) teamAddCheck(name string) (*models.Team, error) {
	team := models.Team{}
	if err := db.DB.Select("id", "name", "stadium_id", "league_id", "sport_id").Where("name = ? ", name).First(&team).Error; err != nil {
		log.Println("エラー:", err)
		return nil, err
	}
	return &team, nil
}

// 　※チーム情報更新時
func (s *AdminToolService) teamUpdateCheck(id, sport_id uint, name string) error {
	league := models.Team{}
	log.Println("モデル生成直後")
	if err := db.DB.Select("id", "stadium_id", "league_id", "sport_id", "name").Where("id != ?", id).Where("name = ? AND sport_id = ?", name, sport_id).First(&league).Error; err != nil {
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

// 該当idからレコードを取得
func (s *AdminToolService) TeamGetIdService(id uint) (*Team, error) {
	var team Team
	if err := db.DB.Select("id", "stadium_id", "league_id", "sport_id", "name").Where("id = ?", id).Find(&team).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewCustomError(http.StatusUnauthorized, "チームが見つかりませんでした")
		}
		return nil, utils.NewCustomError(http.StatusUnauthorized, "内部エラーが発生しました")
	}

	log.Println("チーム名:", team.Name)

	return &team, nil
}

func NewAdminToolService() *AdminToolService {
	return &AdminToolService{}
}
