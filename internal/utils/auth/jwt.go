package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAuthenticator struct {
	secret string
}

func NewJWTAuthenticator(secret string) *JWTAuthenticator {
	return &JWTAuthenticator{
		secret: secret,
	}
}

func (auth *JWTAuthenticator) GenerateToken(userId string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp": time.Now().Add(30 * 24 * time.Hour).Unix(),
			"iat": time.Now().Unix(),
			"sub": userId,
		})
	s, err := t.SignedString([]byte(auth.secret))
	if err != nil {
		return "", fmt.Errorf("failed to create JWT token: %v", err)
	}
	return s, nil
}
