package models

type AdminSignupRequest struct {
	AdminName   string `json:"user_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type AdimnSignupResponse struct {
	ID          string `json:"id"`
	AdminName   string `json:"userName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
