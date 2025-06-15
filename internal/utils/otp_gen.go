package utils

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func generateOTP() string {
	return fmt.Sprintf("%06d", rand.IntN(1000000))
}

func GenerateOTPWithExpiry() (string, time.Time) {
	otp := generateOTP()
	expiry := time.Now().Add(5 * time.Minute)
	return otp, expiry

}
