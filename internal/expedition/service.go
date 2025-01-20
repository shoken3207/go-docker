package expedition

import (
	"errors"
	"go-docker/internal/db"
	"go-docker/models"
	"go-docker/pkg/constants"
	"go-docker/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
	"gorm.io/gorm"
)

type ExpeditionService struct{}



func (s *ExpeditionService) GetLikesCountByExpeditionId(expeditionId uint) (*int64, error) {
	expeditionLikes := []models.ExpeditionLike{}
	if err := db.DB.Where("expedition_id = ?", expeditionId).Find(&expeditionLikes).Error; err != nil {
		log.Printf("遠征記録お気に入り取得エラー: %v", err)
			return nil, utils.NewCustomError(http.StatusInternalServerError, "遠征記録取得に失敗しました。")
	}
	likesCount := int64(len(expeditionLikes))
	return &likesCount, nil
}
func (s *ExpeditionService) FindExpeditionById(expeditionId uint) (*models.Expedition, error) {
	expedition := models.Expedition{}
	if err := db.DB.Where("id = ?", expeditionId).First(&expedition).Error; err != nil {
		log.Printf("遠征記録取得エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, "遠征記録が見つかりませんでした。。")
		} else {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "遠征記録取得に失敗しました。")
		}
	}
	return &expedition, nil
}
func (s *ExpeditionService) CreateExpedition(tx *gorm.DB, newExpedition *models.Expedition) error {
	if err := tx.Create(newExpedition).Error; err != nil {
		log.Printf("遠征記録作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録作成に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) UpdateExpedition(tx *gorm.DB, updateExpedition *models.Expedition) error {
	if err := tx.Save(updateExpedition).Error; err != nil {
		log.Printf("遠征記録更新エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "遠征記録が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録更新に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) DeleteExpeditionService(expeditionId *uint, userId *uint, ik *imagekit.ImageKit) error {
	expedition, err := s.FindExpeditionById(*expeditionId)
	if err != nil {
		return err
	}

	if expedition.UserId != *userId {
		return utils.NewCustomError(http.StatusForbidden, "この遠征記録を削除する権限がありません")
	}

	var expeditionImages []models.ExpeditionImage
	if err := db.DB.Where("expedition_id = ?", expeditionId).Find(&expeditionImages).Error; err != nil {
		log.Printf("遠征記録の画像情報取得エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録の画像情報の取得に失敗しました")
	}

	for _, image := range expeditionImages {
		utils.DeleteUploadImage(ik, &image.FileId)
	}

	if err := db.DB.Delete(&models.Expedition{}, expeditionId).Error; err != nil {
		log.Printf("遠征記録削除エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "遠征記録が見つかりません")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録削除に失敗しました")
	}

	return nil
}

func (s *ExpeditionService) FindPaymentById(paymentId uint) (*models.Payment, error) {
	payment := models.Payment{}
	if err := db.DB.Where("id = ?", paymentId).First(&payment).Error; err != nil {
		log.Printf("遠征記録出費取得エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, "遠征記録出費が見つかりませんでした。。")
		} else {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "遠征記録出費取得に失敗しました。")
		}
	}
	return &payment, nil
}
func (s *ExpeditionService) CreatePayment(tx *gorm.DB, newPayment *models.Payment) error {
	if err := tx.Create(newPayment).Error; err != nil {
		log.Printf("遠征記録出費作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録出費作成に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) UpdatePayment(tx *gorm.DB, updatePayment *models.Payment) error {
	if err := tx.Save(updatePayment).Error; err != nil {
		log.Printf("遠征記録出費更新エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "遠征記録出費が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録出費更新に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) DeletePayments(tx *gorm.DB, paymentIds *[]uint) error {
	if err := tx.Delete(&models.Payment{}, paymentIds).Error; err != nil {
		log.Printf("遠征記録出費削除エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "遠征記録出費が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録出費削除に失敗しました。")
	}
	return nil
}

func (s *ExpeditionService) CreateExpeditionImages(tx *gorm.DB, newExpeditionImages *[]models.ExpeditionImage) error {
	if err := tx.Create(newExpeditionImages).Error; err != nil {
		log.Printf("遠征記録画像データ作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録画像データ作成に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) DeleteExpeditionImages(tx *gorm.DB, fileIds []string) error {
	if err := tx.Where("file_id IN ?", fileIds).Delete(&models.ExpeditionImage{}).Error; err != nil {
		log.Printf("遠征記録画像データ削除エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "遠征記録画像データが見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録画像データ削除に失敗しました。")
	}
	return nil
}

func (s *ExpeditionService) FindVisitedFacilityById(visitedFacilityId uint) (*models.VisitedFacility, error) {
	visitedFacility := models.VisitedFacility{}
	if err := db.DB.Where("id = ?", visitedFacilityId).First(&visitedFacility).Error; err != nil {
		log.Printf("訪れた施設取得エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, "訪れた施設が見つかりませんでした。。")
		} else {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "訪れた施設取得に失敗しました。")
		}
	}
	return &visitedFacility, nil
}
func (s *ExpeditionService) CreateVisitedFacility(tx *gorm.DB, newVisitedFacilitiy *models.VisitedFacility) error {
	if err := tx.Create(newVisitedFacilitiy).Error; err != nil {
		log.Printf("訪れた施設作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "訪れた施設作成に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) UpdateVisitedFacility(tx *gorm.DB, updateVisitedFacilitiy *models.VisitedFacility) error {
	if err := tx.Save(updateVisitedFacilitiy).Error; err != nil {
		log.Printf("訪れた施設更新エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "訪れた施設が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "訪れた施設更新に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) DeleteVisitedFacilities(tx *gorm.DB, visitedFacilityIds *[]uint) error {
	if err := tx.Delete(&models.VisitedFacility{}, visitedFacilityIds).Error; err != nil {
		log.Printf("訪れた施設削除エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "訪れた施設が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "訪れた施設削除に失敗しました。")
	}
	return nil
}

func (s *ExpeditionService) FindGameById(gameId uint) (*models.Game, error) {
	game := models.Game{}
	if err := db.DB.Where("id = ?", gameId).First(&game).Error; err != nil {
		log.Printf("試合記録取得エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, "試合記録が見つかりませんでした。。")
		} else {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "試合記録取得に失敗しました。")
		}
	}
	return &game, nil
}
func (s *ExpeditionService) CreateGame(tx *gorm.DB, newGame *models.Game) error {
	log.Printf("newGame: %v", newGame)
	if err := tx.Create(newGame).Error; err != nil {
		log.Printf("試合記録作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "試合記録作成に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) UpdateGame(tx *gorm.DB, updateGame *models.Game) error {
	if err := tx.Save(updateGame).Error; err != nil {
		log.Printf("試合記録更新エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "試合記録が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "試合記録更新に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) DeleteGames(tx *gorm.DB, gameIds *[]uint) error {
	if err := tx.Delete(&models.Game{}, gameIds).Error; err != nil {
		log.Printf("試合記録削除エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "試合記録が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "試合記録削除に失敗しました。")
	}
	return nil
}

func (s *ExpeditionService) FindGameScoreById(gameScoreId uint) (*models.GameScore, error) {
	gameScore := models.GameScore{}
	if err := db.DB.Where("id = ?", gameScoreId).First(&gameScore).Error; err != nil {
		log.Printf("試合得点取得エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, "試合得点が見つかりませんでした。。")
		} else {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "試合得点取得に失敗しました。")
		}
	}
	return &gameScore, nil
}
func (s *ExpeditionService) CreateGameScore(tx *gorm.DB, newGameScore *models.GameScore) error {
	if err := tx.Create(newGameScore).Error; err != nil {
		log.Printf("試合得点作成エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "試合得点が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "試合得点作成に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) UpdateGameScore(tx *gorm.DB, updateGameScore *models.GameScore) error {
	if err := tx.Save(updateGameScore).Error; err != nil {
		log.Printf("試合得点更新エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "試合得点が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "試合得点更新に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) DeleteGameScores(tx *gorm.DB, gameScoreIds *[]uint) error {
	if err := tx.Delete(&models.GameScore{}, gameScoreIds).Error; err != nil {
		log.Printf("試合得点削除エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "試合得点が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "試合得点削除に失敗しました。")
	}
	return nil
}

func (s *ExpeditionService) GetExpeditionDetailService(request *GetExpeditionDetailRequest) (*GetExpeditionDetailResponse, error) {
	var expedition models.Expedition

	if err := db.DB.Preload("VisitedFacilities").
		Preload("Payments").
		Preload("ExpeditionImages").
		Preload("ExpeditionLikes").
		Preload("Games.Team1").
		Preload("Games.Team2").
		Preload("Games.GameScores.Team").
		Preload("Games.GameScores").
		Preload("Sport").
		Preload("Stadium").
		Preload("User").
		First(&expedition, request.ExpeditionId).Error; err != nil {
		log.Printf("遠征記録詳細取得エラー: %v", err)
		return nil, utils.NewCustomError(http.StatusInternalServerError, "遠征記録詳細の取得に失敗しました。")
	}

	var visitedFacilities []VisitedFacilityResponse
	for _, vf := range expedition.VisitedFacilities {
		visitedFacilities = append(visitedFacilities, VisitedFacilityResponse{
			ID:        int(vf.ID),
			Name:      vf.Name,
			Address:   vf.Address,
			Icon:      vf.Icon,
			Color:     vf.Color,
			Latitude:  vf.Latitude,
			Longitude: vf.Longitude,
		})
	}

	var payments []PaymentResponse
	for _, p := range expedition.Payments {
		payments = append(payments, PaymentResponse{
			ID:    p.ID,
			Title: p.Title,
			Date:  p.Date,
			Cost:  p.Cost,
		})
	}

	var expeditionImages []ExpeditionImageResponse
	for _, img := range expedition.ExpeditionImages {
		expeditionImages = append(expeditionImages, ExpeditionImageResponse{
			ID:     img.ID,
			FileId: img.FileId,
			Image:  img.Image,
		})
	}

	var games []GameResponse
	for _, g := range expedition.Games {
		var scores []GameScoreResponse
		for _, s := range g.GameScores {
			scores = append(scores, GameScoreResponse{
				ID:       s.ID,
				Order:    s.Order,
				Score:    s.Score,
				TeamId:   s.TeamId,
				TeamName: s.Team.Name,
			})
		}
		games = append(games, GameResponse{
			ID:        g.ID,
			Date:      g.Date,
			Team1Id:   g.Team1Id,
			Team1Name: g.Team1.Name,
			Team2Id:   g.Team2Id,
			Team2Name: g.Team2.Name,
			Scores:    scores,
		})
	}

	response := &GetExpeditionDetailResponse{
		ExpeditionResponse: ExpeditionResponse{
			ID:          int(expedition.ID),
			SportId:     expedition.SportId,
			SportName:   expedition.Sport.Name,
			IsPublic:    expedition.IsPublic,
			Title:       expedition.Title,
			StartDate:   expedition.StartDate,
			EndDate:     expedition.EndDate,
			StadiumId:   expedition.StadiumId,
			StadiumName: expedition.Stadium.Name,
			Memo:        expedition.GetMemo(),
			UserId: expedition.UserId,
		},
		Username: expedition.User.Username,
		UserIcon: expedition.User.GetProfileImage(),
		LikesCount: int64(len(expedition.ExpeditionLikes)),
		VisitedFacilities: visitedFacilities,
		Payments:          payments,
		ExpeditionImages:  expeditionImages,
		Games:             games,
	}

	return response, nil
}

func (s *ExpeditionService) CreateExpeditionService(request *CreateExpeditionRequest, userId *uint) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		newExpedition := models.Expedition{
			SportId:   request.SportId,
			IsPublic:  request.IsPublic,
			Title:     request.Title,
			StartDate: request.StartDate,
			EndDate:   request.EndDate,
			StadiumId: request.StadiumId,
			UserId:    *userId,
		}
		newExpedition.SetMemo(request.Memo)
		if err := s.CreateExpedition(tx, &newExpedition); err != nil {
			return err
		}

		games := request.Games
		for _, game := range games {
			newGame := models.Game{
				ExpeditionId: newExpedition.ID,
				Date:         game.Date,
				Team1Id:      game.Team1Id,
				Team2Id:      game.Team2Id,
			}
			if err := s.CreateGame(tx, &newGame); err != nil {
				return err
			}

		for _, gameScore := range game.Scores {
			newGameScore := models.GameScore{
				GameId: newGame.ID,
				TeamId: gameScore.TeamId,
				Score:  gameScore.Score,
				Order:  gameScore.Order,
			}

				if err :=s.CreateGameScore(tx, &newGameScore); err != nil {
					return err
				}
			}
		}

		imageUrls := request.ImageUrls
		if len(imageUrls) > 0 {
			tempImages, err := utils.ValidateAndPersistImages(tx, imageUrls)
			if err != nil {
				return err
			}

			var expeditionImages []models.ExpeditionImage
			for _, tempImage := range tempImages {
				expeditionImages = append(expeditionImages, models.ExpeditionImage{
					ExpeditionId: newExpedition.ID,
					Image:       tempImage.Image,
					FileId:      tempImage.FileId,
				})
			}

			if err := s.CreateExpeditionImages(tx, &expeditionImages); err != nil {
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
			if err := s.CreatePayment(tx, &newPayment); err != nil {
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
			if err := s.CreateVisitedFacility(tx, &newVisitedFacility); err != nil {
				return err
			}
		}

		return nil
	})
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

func (s *ExpeditionService) UpdateExpeditionService(expeditionId *uint, userId *uint, requestBody *UpdateExpeditionRequestBody, ik *imagekit.ImageKit) error {
	expedition, err := s.FindExpeditionById(*expeditionId)
	if err != nil {
		return err
	}

	if expedition.UserId != *userId {
		return utils.NewCustomError(http.StatusForbidden, "この遠征記録を更新する権限がありません")
	}

	expedition.IsPublic = requestBody.IsPublic
	expedition.Title = requestBody.Title
	expedition.SetMemo(requestBody.Memo)
	expedition.StartDate = requestBody.StartDate
	expedition.EndDate = requestBody.EndDate
	expedition.StadiumId = requestBody.StadiumId

	return db.DB.Transaction(func(tx *gorm.DB) error {
		if err := s.UpdateExpedition(tx, expedition); err != nil {
			return err
		}

		games := requestBody.Games
		for _, game := range games.Add {
			newGame := models.Game{
				ExpeditionId: *expeditionId,
				Date:         game.Date,
				Team1Id:      game.Team1Id,
				Team2Id:      game.Team2Id,
			}
			if err := s.CreateGame(tx, &newGame); err != nil {
				return err
			}

			for _, gameScore := range game.Scores {
				newGameScore := models.GameScore{
					GameId: newGame.ID,
					TeamId: gameScore.TeamId,
					Score:  gameScore.Score,
					Order:  gameScore.Order,
				}

				if err := s.CreateGameScore(tx, &newGameScore); err != nil {
					return err
				}
			}
		}
		for _, game := range games.Update {
			updateGame, err := s.FindGameById(game.ID)
			if err != nil {
				return err
			}

			updateGame.Date = game.Date
			updateGame.Team1Id = game.Team1Id
			updateGame.Team2Id = game.Team2Id
			if err := s.UpdateGame(tx, updateGame); err != nil {
				return err
			}

			gameScores := game.Scores
			for _, gameScore := range gameScores.Add {
				newGameScore := models.GameScore{
					GameId: updateGame.ID,
					TeamId: gameScore.TeamId,
					Score:  gameScore.Score,
					Order:  gameScore.Order,
				}

				if err := s.CreateGameScore(tx, &newGameScore); err != nil {
					return err
				}
			}
			for _, gameScore := range gameScores.Update {
				updateGameScore, err := s.FindGameScoreById(gameScore.ID)
				if err != nil {
					return err
				}

				updateGameScore.TeamId = gameScore.TeamId
				updateGameScore.Score = gameScore.Score
				updateGameScore.Order = gameScore.Order

				if err := s.UpdateGameScore(tx, updateGameScore); err != nil {
					return err
				}
			}
			if len(gameScores.Delete) > 0 {
				if err := s.DeleteGameScores(tx, &gameScores.Delete); err != nil {
					return err
				}
			}
		}
		if len(games.Delete) > 0 {
			if err := s.DeleteGames(tx, &games.Delete); err != nil {
				return err
			}
		}

		images := requestBody.Images
		if len(images.Add) > 0 {
			tempImages, err := utils.ValidateAndPersistImages(tx, images.Add)
			if err != nil {
				return err
			}
			var expeditionImages []models.ExpeditionImage
			for _, tempImage := range tempImages {
				expeditionImages = append(expeditionImages, models.ExpeditionImage{
					ExpeditionId: *expeditionId,
					Image:        tempImage.Image,
					FileId:       tempImage.FileId,
				})
			}
			if err := s.CreateExpeditionImages(tx, &expeditionImages); err != nil {
				return err
			}
		}
		if len(images.Delete) > 0 {
			expeditionImages := []models.ExpeditionImage{}
			if err := tx.Where("image IN ?", images.Delete).Find(&expeditionImages).Error; err != nil {
				log.Printf("fileId取得エラー: %v", err)
				return utils.NewCustomError(http.StatusInternalServerError, "一fileIdの取得に失敗しました")
			}
			var fileIds []string
			for _, expeditionImage := range expeditionImages {
				fileIds = append(fileIds, expeditionImage.FileId)
			}
			for _, fileId := range fileIds {
				if err := utils.DeleteUploadImage(ik, &fileId); err != nil {
					return err
				}
			}
			if err := s.DeleteExpeditionImages(tx, fileIds); err != nil {
				return err
			}
		}

		payments := requestBody.Payments
		for _, payment := range payments.Add {
			newPayment := models.Payment{
				Title:        payment.Title,
				Date:         payment.Date,
				Cost:         payment.Cost,
				ExpeditionId: *expeditionId,
			}
			if err := s.CreatePayment(tx, &newPayment); err != nil {
				return err
			}
		}
		for _, payment := range payments.Update {
			updatePayment, err := s.FindPaymentById(payment.ID)
			if err != nil {
				return err
			}
			updatePayment.Title = payment.Title
			updatePayment.Date = payment.Date
			updatePayment.Cost = payment.Cost
			if err := s.UpdatePayment(tx, updatePayment); err != nil {
				return err
			}
		}
		if len(payments.Delete) > 0 {
			if err := s.DeletePayments(tx, &payments.Delete); err != nil {
				return err
			}
		}

		visitedFacilities := requestBody.VisitedFacilities
		for _, visitedFacility := range visitedFacilities.Add {
			newVisitedFacility := models.VisitedFacility{
				ExpeditionId: *expeditionId,
				Name:         visitedFacility.Name,
				Address:      visitedFacility.Address,
				Icon:         visitedFacility.Icon,
				Color:        visitedFacility.Color,
				Latitude:     visitedFacility.Latitude,
				Longitude:    visitedFacility.Longitude,
			}
			if err := s.CreateVisitedFacility(tx, &newVisitedFacility); err != nil {
				return err
			}
		}
		for _, visitedFacility := range visitedFacilities.Update {
			updateVisitedFacility, err := s.FindVisitedFacilityById(visitedFacility.ID)
			if err != nil {
				return err
			}
			updateVisitedFacility.Name = visitedFacility.Name
			updateVisitedFacility.Address = visitedFacility.Address
			updateVisitedFacility.Icon = visitedFacility.Icon
			updateVisitedFacility.Color = visitedFacility.Color
			updateVisitedFacility.Latitude = visitedFacility.Latitude
			updateVisitedFacility.Longitude = visitedFacility.Longitude
			if err := s.UpdateVisitedFacility(tx, updateVisitedFacility); err != nil {
				return err
			}
		}
		if len(visitedFacilities.Delete) > 0 {
			if err := s.DeleteVisitedFacilities(tx, &visitedFacilities.Delete); err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *ExpeditionService) CreateExpeditionLike(userId *uint, expeditionId *uint) error {
	newLike := models.ExpeditionLike{
		UserId:       *userId,
		ExpeditionId: *expeditionId,
	}

	if err := db.DB.Create(&newLike).Error; err != nil {
		log.Printf("いいね作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "いいねの作成に失敗しました")
	}
	return nil
}

func (s *ExpeditionService) CreateExpeditionLikeService(userId *uint, expeditionId *uint) (*int64, error) {
	var existingLike models.ExpeditionLike
	if err := db.DB.Where("user_id = ? AND expedition_id = ?", *userId, *expeditionId).First(&existingLike).Error; err == nil {
		return nil, utils.NewCustomError(http.StatusBadRequest, "既にいいね済みです")
	}

	if _, err := s.FindExpeditionById(*expeditionId); err != nil {
		return nil, err
	}

	if err := s.CreateExpeditionLike(userId, expeditionId); err != nil {
		return nil, err
	}

	likesCount, err := s.GetLikesCountByExpeditionId(*expeditionId);
	if  err != nil {
		return nil, err
	}

	return likesCount, nil
}

func (s *ExpeditionService) DeleteExpeditionLike(userId *uint, expeditionId *uint) error {
	result := db.DB.Where("user_id = ? AND expedition_id = ?", *userId, *expeditionId).Delete(&models.ExpeditionLike{})

	if result.Error != nil {
		log.Printf("いいね削除エラー: %v", result.Error)
		return utils.NewCustomError(http.StatusInternalServerError, "いいねの削除に失敗しました")
	}

	if result.RowsAffected == 0 {
		return utils.NewCustomError(http.StatusNotFound, "いいねが見つかりません")
	}

	return nil
}

func (s *ExpeditionService) DeleteExpeditionLikeService(userId *uint, expeditionId *uint) (*int64, error) {
	if err := s.DeleteExpeditionLike(userId, expeditionId); err != nil {
		return nil, err
	}

	likesCount, err := s.GetLikesCountByExpeditionId(*expeditionId);
	if  err != nil {
		return nil, err
	}

	return likesCount, nil
}

func (s *ExpeditionService) GetExpeditionList(req *ExpeditionListRequest, userId *uint) ([]ExpeditionListResponse, error) {
	offset := (req.Page - 1) * constants.LIMIT_EXPEDITION_LIST
	var expeditions []ExpeditionListResponse

	query := db.DB.Table("expeditions").
		Select(`
			expeditions.id,
			expeditions.title,
			expeditions.start_date,
			expeditions.end_date,
			stadia.name as stadium_name,
			stadia.id as stadium_id,
			expeditions.sport_id,
			sports.name as sport_name,
			expeditions.user_id,
			users.name as user_name,
			users.profile_image as user_icon,
			(SELECT COUNT(*) FROM expedition_likes WHERE expedition_likes.expedition_id = expeditions.id) as likes_count,
			(
				SELECT t1.name
				FROM games g
				JOIN teams t1 ON g.team1_id = t1.id
				WHERE g.expedition_id = expeditions.id
				ORDER BY g.created_at
				LIMIT 1
			) as team1_name,
			(
				SELECT t2.name
				FROM games g
				JOIN teams t2 ON g.team2_id = t2.id
				WHERE g.expedition_id = expeditions.id
				ORDER BY g.created_at
				LIMIT 1
			) as team2_name,
			EXISTS (
				SELECT 1
				FROM expedition_likes
				WHERE expedition_likes.expedition_id = expeditions.id
				AND expedition_likes.user_id = ?
			) as is_liked
		`, *userId).
		Joins("LEFT JOIN sports ON expeditions.sport_id = sports.id").
		Joins("LEFT JOIN users ON expeditions.user_id = users.id").
		Joins("LEFT JOIN stadia ON expeditions.stadium_id = stadia.id").
		Where("expeditions.is_public = ?", true)

	if req.SportId != nil {
		query = query.Where("expeditions.sport_id = ?", *req.SportId)
	}
	if req.TeamId != nil {
		query = query.Where("EXISTS (SELECT 1 FROM games WHERE games.expedition_id = expeditions.id AND (games.team1_id = ? OR games.team2_id = ?))", *req.TeamId, *req.TeamId)
	}

	if err := query.
		Order("expeditions.created_at DESC").
		Limit(constants.LIMIT_EXPEDITION_LIST).
		Offset(offset).
		Find(&expeditions).Error; err != nil {
		log.Printf("遠征記録一覧の取得に失敗しました: %v", err)
		return nil, utils.NewCustomError(http.StatusInternalServerError, "遠征記録一覧の取得に失敗しました")
	}

	if len(expeditions) == 0 {
		if req.Page == 1 {
			return nil, utils.NewCustomError(http.StatusNotFound, "遠征記録が登録されていません")
		} else {
			return nil, utils.NewCustomError(http.StatusNotFound, "最後のページです")
		}
	}

	var expeditionIds []uint
	for _, exp := range expeditions {
		expeditionIds = append(expeditionIds, exp.ID)
	}

	var images []struct {
		ExpeditionID uint
		Image        string
	}
	if err := db.DB.Model(&models.ExpeditionImage{}).
		Select("expedition_id, image").
		Where("expedition_id IN ?", expeditionIds).
		Order("expedition_id, created_at").
		Find(&images).Error; err != nil {
		log.Printf("遠征記録画像の取得に失敗しました: %v", err)
		return nil, utils.NewCustomError(http.StatusInternalServerError, "遠征記録一覧の取得に失敗しました")
	}

	imageMap := make(map[uint][]string)
	for _, img := range images {
		imageMap[img.ExpeditionID] = append(imageMap[img.ExpeditionID], img.Image)
	}

	for i := range expeditions {
		expeditions[i].Images = make([]string, 0)
		if images, ok := imageMap[expeditions[i].ID]; ok {
			expeditions[i].Images = images
		}
	}

	return expeditions, nil
}

func NewExpeditionService() *ExpeditionService {
	return &ExpeditionService{}
}
