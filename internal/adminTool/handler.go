package admintool

import (
	"errors"
	"go-docker/models"
	"go-docker/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminToolHandler struct{}

var adminToolService = NewAdminToolService()

// @summary チームの追加
// @Description リクエストからチーム名を取得後、チーム一覧から同一のチームが存在しない場合に登録する。
// @Tags teams
// @Param teamname path string true "チーム名"
// @Success 200 {object} utils.BasicResponse "成功"
// @Failure 400 {object} utils.BasicResponse "リクエストエラー"
// @Failure 500 {object} utils.BasicResponse "内部エラー"
// @Router /api/teams/teamAdd [post]
func (h *AdminToolHandler) teamAdd(c *gin.Context) {
	request := teamAddRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse[any](c, http.StatusBadRequest, "リクエストに不備があります。")
		return
	}
	team, err := adminToolService.teamSearch(request.TeamName)
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
