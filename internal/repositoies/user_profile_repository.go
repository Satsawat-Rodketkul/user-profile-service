package repositoies

import (
	"log"
	"user-profile-service/internal/database"
	"user-profile-service/internal/models"
)

func Usersignup(userProfile *models.UserProfile) error {
	tx := database.DB.Create(userProfile)
	if tx.Error != nil {
		log.Println("Error create user info:", tx.Error)
	} else {
		log.Println("Create user info success")
	}
	return tx.Error
}
