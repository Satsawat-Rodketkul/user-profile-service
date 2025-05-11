package service

import (
	"log"
	db "user-profile-service/internal/database/repositoies"
	"user-profile-service/internal/models"
	"user-profile-service/internal/redis/model"
	rd "user-profile-service/internal/redis/repositoies"
)

func UserProfileService(userId string) (*models.UserProfileResponse, error) {
	dataRedis, err := rd.UserProfileGet(userId)
	if err != nil {
		log.Fatal("Error get user profile from redis")
		return nil, err
	}

	if dataRedis != nil {
		return &models.UserProfileResponse{
			Username:     dataRedis.Username,
			Email:        dataRedis.Email,
			MobileNumber: dataRedis.MobileNumber,
		}, nil
	} else {
		dataDatabase, err := db.UserProfileGet(userId)
		if err != nil {
			log.Fatal("Error get user profile from database")
			return nil, err
		}

		rd.UserProfileSet(userId, model.UserProfile{
			Username:     dataDatabase.Username,
			Email:        dataDatabase.Email,
			MobileNumber: dataDatabase.MobileNumber,
		})

		return &models.UserProfileResponse{
			Username:     dataDatabase.Username,
			Email:        dataDatabase.Email,
			MobileNumber: dataDatabase.MobileNumber,
		}, nil
	}
}
