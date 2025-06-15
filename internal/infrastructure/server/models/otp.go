package models

import "time"

type OTP struct {
	Email     string    `bson:"email"`
	Otp       string    `bson:"otp"`
	ExpiresAt time.Time `bson:"expires_at"`
}

type UserVerifyOTPRequest struct {
	Email    string `json:"email"`
	Otp      string `json:"otp"`
	Password string `json:"password"`
}

type UserOTPRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
