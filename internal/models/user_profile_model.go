package models

type UserProfileResponse struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobileNumber"`
}
