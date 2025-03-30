package models

import "time"

type UserProfile struct {
	UserId       string
	Username     string
	Password     string
	Email        string
	MobileNumber string
	CreateDate   time.Time
	UpdateDate   time.Time
}

func (UserProfile) TableName() string {
	return "user_profile"
}
