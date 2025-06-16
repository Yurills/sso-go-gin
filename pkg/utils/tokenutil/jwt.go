package tokenutil

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("sso-go-gin-secret")

func GenerateJWTToken(userID string, userEmail string, nonce string, ttl int) (string, error) {
	claims := jwt.MapClaims{
		"sub":   userID,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Second * time.Duration(ttl)).Unix(), // Token valid for 24 hours
		"email": userEmail,
		"nonce": nonce,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
