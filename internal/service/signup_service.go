package service

import (
	"time"
	"user-profile-service/internal/database/entities"
	db "user-profile-service/internal/database/repositoies"
	"user-profile-service/internal/models"
	"user-profile-service/internal/utils"
)

func SingupService(request models.UserSignupRequest) error {
	userProfile := entities.UserProfile{
		UserId:       utils.GenerateUserId(),
		Username:     request.Username,
		Password:     request.Password,
		Email:        request.Email,
		MobileNumber: request.MobileNumber,
		CreateDate:   time.Now(),
		UpdateDate:   time.Now(),
	}

	response := db.UserSignup(&userProfile)

	return response
}
