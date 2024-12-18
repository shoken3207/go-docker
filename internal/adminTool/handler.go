package adminTool

import (
	"errors"
	"go-docker/models"
	"go-docker/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminToolHandler struct{}

var adminToolService = NewAdminToolService()

// @summary チームの追加
// @Description リクエストからチーム名を取得後、チーム一覧から同一のチームが存在しない場合に登録する。
// @Tags teams
// @Param request body TeamAddRequest true "チーム名"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/teams/teamAdd [post]
func (h *AdminToolHandler) TeamAdd(c *gin.Context) {
	request := TeamAddRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("リクエストエラー: %v", err)
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}
	team, err := adminToolService.teamCheck(request.TeamName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse[any](c, http.StatusInternalServerError, "内部エラーが発生しました。")
		return
	}
	if team != nil {
		utils.ErrorResponse[any](c, http.StatusConflict, "登録済みのチームです。")
		return
	}

	newTeam := models.Team{StadiumId: request.StadiumId, SportId: request.SportsId}

	if err := adminToolService.createTeam(&newTeam); err != nil {
		utils.ErrorResponse[any](c, http.StatusInternalServerError, "内部エラーが発生しました。")
		return
	}
	utils.SuccessResponse[any](c, http.StatusOK, nil, "チーム登録に成功しました。")
}

// @Summary スタジアム追加
// @Description リクエストからスタジアム情報を追加後、重複確認を行い登録する。
// @Tags stadium
// @Param request body StadiumAddRequest true "スタジアム情報"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/stadium/stadiumAdd [post]
func (h *AdminToolHandler) StadiumAdd(c *gin.Context) {
	request := StadiumAddRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります")
		return
	}
	stadium, err := adminToolService.stadiumCheck(request.Name, request.Address)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse[any](c, http.StatusInternalServerError, "内部エラーが発生しました。")
		return
	}
	if stadium != nil {
		utils.ErrorResponse[any](c, http.StatusConflict, "登録済みのスタジアムです")
		return
	}

	newStadium := models.Stadium{Name: request.Name, Description: request.Description, Address: request.Address, Capacity: int(request.Capacity), Image: request.Image}

	if err := adminToolService.createStadium(&newStadium); err != nil {
		utils.ErrorResponse[any](c, http.StatusInternalServerError, "内部エラーが発生しました。")
		return
	}
	utils.SuccessResponse[any](c, http.StatusOK, nil, "スタジアム登録に成功しました。")
}

// @Summary スタジアム削除
// @Description リクエストボディに削除対象のIDを指定してスタジアムを削除します
// @Tags stadium
// @Param request body DeleteRequest true "スタジアムID"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/stadium/delete [delete]
func (h *AdminToolHandler) DeleteStadium(c *gin.Context) {
	var request DeleteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "無効なリクエスト形式です")
		return
	}

	stadium, err := adminToolService.stadiumSearch(request.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse[any](c, http.StatusNotFound, "スタジアムが見つかりませんでした")
			return
		}
		utils.ErrorResponse[any](c, http.StatusInternalServerError, "内部エラーが発生しました。")
		return
	}

	if err := adminToolService.deleteStadium(stadium); err != nil {
		utils.ErrorResponse[any](c, http.StatusInternalServerError, "削除に失敗しました")
		return
	}
	utils.SuccessResponse[any](c, http.StatusOK, nil, "チームが正常に削除されました")
}

func NewAdminToolHandler() *AdminToolHandler {
	return &AdminToolHandler{}
}
