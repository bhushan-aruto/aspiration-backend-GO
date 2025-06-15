package models

type UserSignupRequest struct {
	UserName    string `json:"user_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

type UserSignupResponse struct {
	ID          string `json:"id"`
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type GetUserResponse struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	OTP         string `json:"otp"`
	NewPassword string `json:"new_password"`
}
