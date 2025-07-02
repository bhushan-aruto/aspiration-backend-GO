package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecrete = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claim struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userId string) (string, error) {
	claims := Claim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecrete)
	if err != nil {
		return "", err
	}
	return signedToken, nil

}

func ValidateJWT(tokenStr string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecrete, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claim)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
