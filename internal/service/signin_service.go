package service

import (
	"log"
	db "user-profile-service/internal/database/repositoies"
	"user-profile-service/internal/models"
	"user-profile-service/internal/utils"
)

func SinginService(request models.UserSigninRequest) (*models.UserSigninResponse, error) {
	result, err := db.UserSingin(request.Username, request.Password)
	if err != nil {
		log.Fatal("Error signin process")
		return nil, err
	}

	token, err := utils.GenerateJWT(result.UserId)
	if err != nil {
		log.Fatal("Error generate token")
		return nil, err
	}

	return &models.UserSigninResponse{
		Token: token,
	}, nil
}
