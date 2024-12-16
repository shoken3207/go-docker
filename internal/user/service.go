package user

import (
	"go-docker/internal/db"
	"go-docker/models"
	"log"
)

type UserService struct{}

func (s *UserService) findUserById(userId uint) (*models.User, error) {
	user := models.User{}
	if err := db.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) updateUser(userId *uint, request *UpdateUserRequestBody) (*models.User, error) {
	user := models.User{}
	if err := db.DB.Where("id = ?", &userId).First(&user).Error; err != nil {
		log.Println("ユーザー検索エラー:", err)
		return nil, err
	}

	user.Name = request.Name
	user.Description = request.Description
	user.ProfileImage = request.ProfileImage

	if err := db.DB.Model(&user).Updates(models.User{Name: request.Name, Description: request.Description, ProfileImage: request.ProfileImage}).Error; err != nil {
		log.Println("ユーザー更新エラー:", err)
		return nil, err
	}

	return &user, nil
}

func NewUserService() *UserService {
	return &UserService{}
}
