package team

import (
	"go-docker/internal/db"
	"go-docker/models"
	"go-docker/pkg/constants"
	"go-docker/pkg/utils"
	"log"
	"net/http"
)

type TeamService struct{}

func (s *TeamService) GetFavoriteTeams(userId *uint) (*[]models.FavoriteTeam, error) {
	var favoriteTeams []models.FavoriteTeam
	if err := db.DB.Where("user_id = ?", *userId).Find(&favoriteTeams).Error; err != nil {
		log.Printf("お気に入りチーム取得エラー: %v", err)
		return nil, utils.NewCustomError(http.StatusInternalServerError, "お気に入りチーム取得に失敗しました。")
	}
	return &favoriteTeams, nil
}

func (s *TeamService) GetTeamsService(userId *uint) (*[]SportResponse, error) {
	favoriteMap := make(map[uint]bool)
	if userId != nil {
		favoriteTeams, err := s.GetFavoriteTeams(userId)
		if err != nil {
			return nil, err
		}
		for _, favoriteTeam := range *favoriteTeams {
			favoriteMap[favoriteTeam.TeamId] = true
		}
	}

	var sports []models.Sport
	err := db.DB.Preload("Leagues.Teams").Find(&sports).Error
	if err != nil {
		log.Printf("データベースエラー: %v", err)
		return nil, utils.NewCustomError(http.StatusInternalServerError, "データ取得に失敗しました。")
	}

	var response []SportResponse
	for _, sport := range sports {
		var leagues []LeagueResponse
		for _, league := range sport.Leagues {
			var teams []TeamResponse
			for _, team := range league.Teams {
				isFavorite := false
				if userId != nil {
					isFavorite = favoriteMap[team.ID]
				}
				teams = append(teams, TeamResponse{
					ID: team.ID,
					Name: team.Name,
					IsFavorite: isFavorite,
				})
			}
			leagues = append(leagues, LeagueResponse{
				League: league.Name,
				Teams: teams,
			})
		}
		response = append(response, SportResponse{
			Sport: sport.Name,
			Icon: constants.SportIcon[int(sport.ID)],
			Team: leagues,
		})
	}

	return &response, nil
}

func (s *TeamService) GetTeamsBySportsId(sportsId *uint) (*[]TeamListResponse, error) {
	var teams []models.Team
	if err := db.DB.Select("id", "name").Where("sport_id = ?", *sportsId).Find(&teams).Error; err != nil {
		log.Printf("チーム取得エラー: %v", err)
		return nil, utils.NewCustomError(http.StatusInternalServerError, "チーム取得に失敗しました。")
	}
	var teamResponse []TeamListResponse

	for _, team := range teams {
		teamResponse = append(teamResponse, TeamListResponse{
			ID:   team.ID,
			Name: team.Name,
		})
	}
	return &teamResponse, nil
}

func NewTeamService() *TeamService {
	return &TeamService{}
}