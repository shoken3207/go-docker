package expedition

import (
	"errors"
	"go-docker/internal/db"
	"go-docker/models"
	"go-docker/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imagekit-developer/imagekit-go"
	"gorm.io/gorm"
)

type ExpeditionService struct{}

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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "遠征記録が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録更新に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) DeleteExpedition(expeditionId *uint, userId *uint, ik *imagekit.ImageKit) error {
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "遠征記録出費が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録出費更新に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) DeletePayments(paymentIds *[]uint) error {
	if err := db.DB.Delete(&models.Payment{}, paymentIds).Error; err != nil {
		log.Printf("遠征記録出費削除エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "遠征記録出費が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "遠征記録出費削除に失敗しました。")
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
func (s *ExpeditionService) DeleteExpeditionImages(fileIds *[]string) error {
	if err := db.DB.Where("file_id IN ?", fileIds).Delete(&models.ExpeditionImage{}).Error; err != nil {
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
func (s *ExpeditionService) CreateVisitedFacility(newVisitedFacilitiy *models.VisitedFacility) error {
	if err := db.DB.Create(newVisitedFacilitiy).Error; err != nil {
		log.Printf("訪れた施設作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "訪れた施設作成に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) UpdateVisitedFacility(updateVisitedFacilitiy *models.VisitedFacility) error {
	if err := db.DB.Save(updateVisitedFacilitiy).Error; err != nil {
		log.Printf("訪れた施設更新エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "訪れた施設が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "訪れた施設更新に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) DeleteVisitedFacilities(visitedFacilityIds *[]uint) error {
	if err := db.DB.Delete(&models.VisitedFacility{}, visitedFacilityIds).Error; err != nil {
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
func (s *ExpeditionService) CreateGame(newGame *models.Game) error {
	if err := db.DB.Create(newGame).Error; err != nil {
		log.Printf("試合記録作成エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "試合記録作成に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) UpdateGame(updateGame *models.Game) error {
	if err := db.DB.Save(updateGame).Error; err != nil {
		log.Printf("試合記録更新エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "試合記録が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "試合記録更新に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) DeleteGames(gameIds *[]uint) error {
	if err := db.DB.Delete(&models.Game{}, gameIds).Error; err != nil {
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
func (s *ExpeditionService) CreateGameScore(newGameScore *models.GameScore) error {
	if err := db.DB.Create(newGameScore).Error; err != nil {
		log.Printf("試合得点作成エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "試合得点が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "試合得点作成に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) UpdateGameScore(updateGameScore *models.GameScore) error {
	if err := db.DB.Save(updateGameScore).Error; err != nil {
		log.Printf("試合得点更新エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "試合得点が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "試合得点更新に失敗しました。")
	}
	return nil
}
func (s *ExpeditionService) DeleteGameScores(gameScoreIds *[]uint) error {
	if err := db.DB.Delete(&models.GameScore{}, gameScoreIds).Error; err != nil {
		log.Printf("試合得点削除エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewCustomError(http.StatusNotFound, "試合得点が見つかりません。")
		}
		return utils.NewCustomError(http.StatusInternalServerError, "試合得点削除に失敗しました。")
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
	expedition.Memo = requestBody.Memo
	expedition.StartDate = requestBody.StartDate
	expedition.EndDate = requestBody.EndDate
	expedition.StadiumId = requestBody.StadiumId

	if err := s.UpdateExpedition(expedition); err != nil {
		return err
	}

	games := requestBody.Games
	for _, game := range games.Add {
		newGame := models.Game{
			ExpeditionId: *expeditionId,
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
	for _, game := range games.Update {
		updateGame, err := s.FindGameById(*&game.ID)
		if err != nil {
			return err
		}

		updateGame.Date = game.Date
		updateGame.Team1Id = game.Team1Id
		updateGame.Team2Id = game.Team2Id
		updateGame.Comment = game.Comment
		if err := s.UpdateGame(updateGame); err != nil {
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

			if err := s.CreateGameScore(&newGameScore); err != nil {
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

			if err := s.UpdateGameScore(updateGameScore); err != nil {
				return err
			}
		}
		if err := s.DeleteGameScores(&gameScores.Delete); err != nil {
			return err
		}
	}
	if err := s.DeleteGames(&games.Delete); err != nil {
		return err
	}

	images := requestBody.Images
	for _, image := range images.Add {
		newExpeditionImage := models.ExpeditionImage{
			ExpeditionId: *expeditionId,
			Image:        image.Image,
			FileId:       image.FileId,
		}
		if err := s.CreateExpeditionImage(&newExpeditionImage); err != nil {
			return err
		}
	}
	for _, fileId := range images.Delete {
		utils.DeleteUploadImage(ik, &fileId)
	}
	s.DeleteExpeditionImages(&images.Delete)

	payments := requestBody.Payments
	for _, payment := range payments.Add {
		newPayment := models.Payment{
			Title:        payment.Title,
			Date:         payment.Date,
			Cost:         payment.Cost,
			ExpeditionId: *expeditionId,
		}
		if err := s.CreatePayment(&newPayment); err != nil {
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
		if err := s.UpdatePayment(updatePayment); err != nil {
			return err
		}
	}
	s.DeletePayments(&payments.Delete)

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
		if err := s.CreateVisitedFacility(&newVisitedFacility); err != nil {
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
		if err := s.UpdateVisitedFacility(updateVisitedFacility); err != nil {
			return err
		}
	}
	s.DeleteVisitedFacilities(&visitedFacilities.Delete)

	return nil
}

func (s *ExpeditionService) CreateExpeditionLike(userId *uint, expeditionId *uint) error {
	var existingLike models.ExpeditionLike
	if err := db.DB.Where("user_id = ? AND expedition_id = ?", *userId, *expeditionId).First(&existingLike).Error; err == nil {
		return utils.NewCustomError(http.StatusBadRequest, "既にいいね済みです")
	}

	if _, err := s.FindExpeditionById(*expeditionId); err != nil {
		return err
	}

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

func NewExpeditionService() *ExpeditionService {
	return &ExpeditionService{}
}
