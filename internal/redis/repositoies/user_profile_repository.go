package repositoies

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"user-profile-service/internal/redis"
	"user-profile-service/internal/redis/model"
)

func UserProfileGet(userId string) (model.UserProfile, error) {
	redisKey := fmt.Sprintf("USER:PROFILE:%s", userId)
	userProfile := model.UserProfile{}

	result, err := redis.Redis.Get(context.Background(), redisKey).Result()
	if err != nil {
		return userProfile, fmt.Errorf("Error get redis: %s", err)
	}

	err = json.Unmarshal([]byte(result), userProfile)
	if err != nil {
		return userProfile, fmt.Errorf("Unmarshal error: %s", err)
	}

	return userProfile, nil
}

func UserProfileSet(userId string, userProfile model.UserProfile) error {
	redisKey := fmt.Sprintf("USER:PROFILE:%s", userId)
	expiration := time.Hour * 24

	data, err := json.Marshal(userProfile)
	if err != nil {
		return fmt.Errorf("Marshal error: %s", err)
	}

	err = redis.Redis.Set(context.Background(), redisKey, string(data), expiration).Err()
	if err != nil {
		return fmt.Errorf("Error set redis: %s", err)
	}

	return nil
}
