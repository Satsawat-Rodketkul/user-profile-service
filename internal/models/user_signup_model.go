package models

type UsersignupRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
}
