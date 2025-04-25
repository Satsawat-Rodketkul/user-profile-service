package repositoies

import "fmt"

func UserProfileGet(userId string) {
	redisKey := fmt.Sprintf("USER:PROFILE:%s", userId)

	print(redisKey)
}
