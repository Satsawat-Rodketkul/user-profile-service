package models

type UserSignupRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
}
