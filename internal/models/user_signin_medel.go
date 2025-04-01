package models

type UserSigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSigninResponse struct {
	Token string `json:"token"`
}
