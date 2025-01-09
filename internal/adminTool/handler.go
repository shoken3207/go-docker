package adminTool

import (
	"go-docker/models"
	"go-docker/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminToolHandler struct{}

var adminToolService = NewAdminToolService()

// @summary チームの追加
// @Description リクエストからチーム名を取得後、チーム一覧から同一のチームが存在しない場合に登録する。
// @Tags teams
// @Security BearerAuth
// @Param request body TeamAddRequest true "チーム名"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/teams/teamAdd [post]
// func (h *AdminToolHandler) TeamAdd(c *gin.Context) {
// 	request := TeamAddRequest{}
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		log.Printf("リクエストエラー: %v", err)
// 		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
// 		return
// 	}

// 	if err := adminToolService.teamCheck(request.TeamName); err != nil {
// 		if customErr, ok := err.(*utils.CustomError); ok {
// 			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
// 			return
// 		}
// 	}

// 	newTeam := models.Team{StadiumId: request.StadiumId, SportId: request.SportsId}

// 	if err := adminToolService.createTeam(&newTeam); err != nil {
// 		utils.ErrorResponse[any](c, http.StatusInternalServerError, "内部エラーが発生しました。")
// 		return
// 	}
// 	utils.SuccessResponse[any](c, http.StatusOK, nil, "チーム登録に成功しました。")

// 	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 		utils.ErrorResponse[any](c, http.StatusInternalServerError, "内部エラーが発生しました。")
// 		return
// 	}
// 	if team != nil {
// 		utils.ErrorResponse[any](c, http.StatusConflict, "登録済みのチームです。")
// 		return
// 	}
// }

// @Summary スタジアム全件検索
// @Description スタジアム情報のレコードを全件取得して、一覧として表示する。
// @Tags stadium
// @Secrity BearerAuth
// @Param keyword query string false "キーワード"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/stadium/stadiums [get]
func (h *AdminToolHandler) GetStadiums(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	log.Println("キーワード:", keyword)
	stadiums := []models.Stadium{}

	stadiums, err := adminToolService.getStadiumsService(keyword)

	if err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	log.Println(stadiums)

	utils.SuccessResponse[any](c, http.StatusOK, stadiums, "スタジアムの検索に成功しました。")
}

// @Summary スタジアム追加
// @Description リクエストからスタジアム情報を追加後、重複確認を行い登録する。
// @Tags stadium
// @Security BearerAuth
// @Param request body StadiumAddRequest true "スタジアム情報"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/stadium/stadiumAdd [post]
func (h *AdminToolHandler) StadiumAdd(c *gin.Context) {
	request := StadiumAddRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.createStadiumService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "スタジアム登録に成功しました。")
}

// @Summary スタジアム更新
// @Description リクエストボディに更新対象のIDを指定してスタジアムを更新します
// @Tags stadium
// @Security BearerAuth
// @Param request body StadiumUppdateRequest true "スタジアム情報"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/stadium/update [put]
func (h *AdminToolHandler) StadiumUpdate(c *gin.Context) {
	request := StadiumUppdateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}
	log.Println("重複検索が正常にリターンはされているよ:handler")
	if err := adminToolService.UpdateStadiumService(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "スタジアムが正常に更新更新されました。")
}

// @Summary スタジアム削除
// @Description リクエストボディに削除対象のIDを指定してスタジアムを削除します
// @Tags stadium
// @Security BearerAuth
// @Param request body DeleteRequest true "スタジアムID"
// @Success 200 {object} utils.SuccessBasicResponse "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/stadium/delete [delete]
func (h *AdminToolHandler) DeleteStadium(c *gin.Context) {
	request := DeleteRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}

	if err := adminToolService.deleteStadiumService(&request); err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, customErr.Error())
			return
		}
	}

	utils.SuccessResponse[any](c, http.StatusOK, nil, "スタジアムが正常に削除されました。")
}

func NewAdminToolHandler() *AdminToolHandler {
	return &AdminToolHandler{}
}
