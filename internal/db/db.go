package db

import (
	"go-docker/config"
	"go-docker/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := config.GetDBConfig()
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %V", err)
	}

	err = DB.AutoMigrate(
		&models.EmailVerification{},
		&models.Expedition{},
		&models.ExpeditionImage{},
		&models.ExpeditionLike{},
		&models.FavoriteTeam{},
		&models.Game{},
		&models.GameScore{},
		&models.League{},
		&models.Payment{},
		&models.Sport{},
		&models.Stadium{},
		&models.Team{},
		&models.User{},
		&models.VisitedFacility{},

	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %V", err)
	}
}