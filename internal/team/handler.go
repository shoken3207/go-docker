package team

import (
	"go-docker/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct{}

var teamService = NewTeamService()

// @Summary クライアントで推しチームを選択する際に必要なチーム情報を取得するAPI
// @Description sport, leagueで入れ子になったteamを返却<br>最初にお気に入りチームを追加する際は、認証していないため全てのisFavoriteをfalseにして返す
// @Tags team
// @Success 200 {object} utils.ApiResponse[[]SportResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/team/public [get]
func (h *TeamHandler) GetTeamsWithoutFavorites(c *gin.Context) {
	response, err := teamService.GetTeamsService(nil)
    if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
    }

	utils.SuccessResponse[[]SportResponse](c,http.StatusOK, *response, utils.CreateSingleMessage("チームの取得に成功しました。"))
}

// @Summary クライアントで推しチームを選択する際に必要なチーム情報を取得するAPI
// @Description sport, leagueで入れ子になったteamを返却<br>認証後、ログイン済みのuserIdからfavoriteTeamsを取得し、isFavoriteにtrueかfalseを設定する
// @Tags team
// @Security BearerAuth
// @Success 200 {object} utils.ApiResponse[[]SportResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/team/me [get]
func (h *TeamHandler) GetTeamsWithFavorites(c *gin.Context) {
	userId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}
	response, err := teamService.GetTeamsService(userId)
    if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
    }

	utils.SuccessResponse[[]SportResponse](c,http.StatusOK, *response, utils.CreateSingleMessage("チームの取得に成功しました。"))
}

// @Summary クライアントでチーム一覧を表示する際に必要なチーム情報を取得するAPI
// @Description スポーツIDからチームをすべて取得する
// @Tags team
// @Security BearerAuth
// @Param sportsId path int true "スポーツID"
// @Success 200 {object} utils.ApiResponse[[]TeamListResponse] "成功"
// @Failure 400 {object} utils.ErrorBasicResponse "リクエストエラー"
// @Failure 500 {object} utils.ErrorBasicResponse "内部エラー"
// @Router /api/team/{sportsId} [get]
func (h *TeamHandler) GetTeamBySportsId(c *gin.Context) {
	var requestPath GetTeamBySportsIdRequestPath
	_, err, customErr := utils.ValidateRequest(c, &requestPath, nil, nil, true)
	if err != nil {
		if customErr, ok := customErr.(*utils.CustomError); ok {
			utils.HandleCustomError(c, customErr, err, requestPath)
			return
		}
	}
	response, err := teamService.GetTeamsBySportsId(&requestPath.SportsId)
	if err != nil {
		if customErr, ok := err.(*utils.CustomError); ok {
			utils.ErrorResponse[any](c, customErr.Code, utils.CreateSingleMessage(customErr.Error()))
			return
		}
	}

	utils.SuccessResponse[[]TeamListResponse](c, http.StatusOK, *response, utils.CreateSingleMessage("チームの取得に成功しました。"))
}

func NewTeamHandler() *TeamHandler {
	return &TeamHandler{}
}
