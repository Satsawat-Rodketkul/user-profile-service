package utils

import (
	"github.com/google/uuid"
)

func GenerateUserId() string {
	return uuid.NewString()[:8]
}
