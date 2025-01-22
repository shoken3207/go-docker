package user

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

type UserService struct{}

func (s *UserService) createUserResponse(user *models.User) *UserResponse {
	userResponse := UserResponse{
		Id:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Name:         user.Name,
		Description:  user.GetDescription(),
		ProfileImage: user.GetProfileImage(),
		FileId:       user.GetFileId(),
	}
	return &userResponse
}

func (s *UserService) DeleteFavoriteTeams(tx *gorm.DB, userId uint) error {
	if err := tx.Where("user_id = ?", userId).Delete(&models.FavoriteTeam{}).Error; err != nil {
		log.Printf("お気に入りチーム削除エラー: %v", err)
		return utils.NewCustomError(http.StatusInternalServerError, "お気に入りチーム削除に失敗しました。")
	}
	return nil
}

func (s *UserService) findUserById(userId uint) (*models.User, error) {
	user := models.User{}
	if err := db.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		log.Printf("ユーザーデータ取得エラー: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewCustomError(http.StatusNotFound, "ユーザーデータが見つかりませんでした。。")
		} else {
			return nil, utils.NewCustomError(http.StatusInternalServerError, "ユーザーデータ取得に失敗しました。")
		}
	}
	return &user, nil
}

func (s *UserService) updateUser(ik *imagekit.ImageKit, userId *uint, request *UpdateUserRequestBody) (*models.User, error) {
	user, err := s.findUserById(*userId)
	if err != nil {
		return nil, err
	}
	if request.Username != user.Username {
		isUnique, err := utils.CheckUsernameUnique(request.Username)
		if err != nil {
			if customErr, ok := err.(*utils.CustomError); ok && customErr.Code != http.StatusNotFound {
				return nil, err
			}
		}
		if !isUnique {
			return nil, utils.NewCustomError(http.StatusConflict, "ユーザーネームが被っています。")
		}
	}
	return user, db.DB.Transaction(func(tx *gorm.DB) (error) {
		if user.GetProfileImage() != request.ProfileImage {
			if request.ProfileImage == "" && user.GetProfileImage() != "" {
				if err := utils.DeleteUploadImage(ik, user.FileId); err != nil {
					return err
				}
				user.SetFileId("")
			}

			if request.ProfileImage != "" && user.GetProfileImage() != "" {
				if err := utils.DeleteUploadImage(ik, user.FileId); err != nil {
					return err
				}
			}

			if request.ProfileImage != "" {
				validatedImages, err := utils.ValidateAndPersistImages(tx, []string{request.ProfileImage})
				if err != nil {
					return err
				}
				if len(validatedImages) > 0 {
					user.SetFileId(validatedImages[0].FileId)
				}
			}
		}
		user.Username = request.Username
		user.Name = request.Name
		user.SetDescription(request.Description)
		user.SetProfileImage(request.ProfileImage)
		if err := db.DB.Model(user).Select("username", "profile_image", "file_id", "name", "description").Updates(user).Error; err != nil {
			log.Printf("ユーザー更新エラー: %v", err)
			return utils.NewCustomError(http.StatusInternalServerError, "ユーザーデータがの更新に失敗しました。")
		}
		return nil
	})

}

func (s *UserService) getUserByIdService(request *GetUserByIdRequest) (*UserResponse, error) {
	user, err := userService.findUserById(request.UserId)
	if err != nil {
		return nil, err
	}
	userResponse := s.createUserResponse(user)

	return userResponse, nil
}

func (s *UserService) getUserByUsernameService(request *GetUserByUsernameRequest) (*UserResponse, error) {
	user, err := utils.FindUserByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	userResponse := s.createUserResponse(user)

	return userResponse, nil
}

func (s *UserService) validateUpdateUserRequest(c *gin.Context) (*uint, *UpdateUserRequestBody, error) {
	var requestBody UpdateUserRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Printf("リクエストエラー: %v", err)
		return nil, nil, utils.NewCustomError(http.StatusBadRequest, "リクエストに不備があります。")
	}

	userId, err := utils.StringToUint(c.GetString("userId"))
	if err != nil {
		return nil, nil, err
	}

	return userId, &requestBody, nil
}

func (s *UserService) updateUserService(ik *imagekit.ImageKit, userId *uint, requestBody *UpdateUserRequestBody) (*UserResponse, error) {
	var userResponse UserResponse
	return &userResponse, db.DB.Transaction(func (tx *gorm.DB) error {
		user, err := s.updateUser(ik, userId, requestBody)
		if err != nil {
			return err
		}

		if(len(requestBody.FavoriteTeams) > 0) {
			if err := s.DeleteFavoriteTeams(tx, user.ID); err != nil {
				return err
			}
			var favoriteTeams []models.FavoriteTeam
			for _, teamId := range requestBody.FavoriteTeams {
				favoriteTeams = append(favoriteTeams, models.FavoriteTeam{
					UserId: user.ID,
					TeamId: uint(teamId),
				})
			}
			if err := utils.CreateFavoriteTeams(tx, &favoriteTeams); err != nil {
				return err
			}
		}

		userResponse = UserResponse{
			Id:           user.ID,
			Username:     user.Username,
			Email:        user.Email,
			Name:         user.Name,
			Description:  user.GetDescription(),
			ProfileImage: user.GetProfileImage(),
			FileId:       user.GetFileId(),
		}
		return nil
	})
}

func NewUserService() *UserService {
	return &UserService{}
}
