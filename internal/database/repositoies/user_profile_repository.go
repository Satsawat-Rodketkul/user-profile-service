package repositoies

import (
	"log"
	"user-profile-service/internal/database"
	"user-profile-service/internal/database/entities"

	"gorm.io/gorm"
)

func UserSignup(userProfile *entities.UserProfile) error {
	tx := database.DB.Create(userProfile)
	if tx.Error != nil {
		log.Println("Error create user info:", tx.Error)
	} else {
		log.Println("Create user info success")
	}
	return tx.Error
}

func UserSingin(username string, password string) (entities.UserProfile, error) {
	userProfile := entities.UserProfile{}
	tx := database.DB.Where("username = ? AND password = ?", username, password).Find(&userProfile)
	if tx.Error != nil {
		log.Println("Error query user info:", tx.Error)
		return userProfile, tx.Error
	}
	if tx.RowsAffected <= 0 {
		log.Println("User info not found")
		return userProfile, gorm.ErrRecordNotFound
	}
	log.Println("Found user info:", userProfile)
	return userProfile, tx.Error
}

func UserProfileGet(userId string) (entities.UserProfile, error) {
	userProfile := entities.UserProfile{}
	tx := database.DB.Where("user_id = ?", userId).Find(&userProfile)
	if tx.Error != nil {
		log.Println("Error query user info:", tx.Error)
		return userProfile, tx.Error
	}
	if tx.RowsAffected <= 0 {
		log.Println("User info not found")
		return userProfile, gorm.ErrRecordNotFound
	}
	log.Println("Found user info:", userProfile)
	return userProfile, tx.Error
}
