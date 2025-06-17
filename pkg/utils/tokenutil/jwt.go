package tokenutil

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("sso-go-gin-secret")

type JWTTokenParams struct {
	Username string
	Email  string
	Nonce  string
	TTL    time.Duration
}

func GenerateJWTToken(params JWTTokenParams) (string, error) {
	claims := jwt.MapClaims{
		"sub":   params.Username,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Second * time.Duration(params.TTL)).Unix(), // Token valid for 24 hours
		"email": params.Email,
		"nonce": params.Nonce,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseAndValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			return nil, jwt.ErrTokenExpired
		}
	}

	return claims, nil
}
