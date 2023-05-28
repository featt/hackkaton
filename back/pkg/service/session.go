package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

type SessionService struct {
	JWTSecretKey []byte
}

func NewSessionService(secretKey string) *SessionService {
	return &SessionService{
		JWTSecretKey: []byte(secretKey),
	}
}

func (s *SessionService) ExtractUserIdFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no token provided")
	}

	// The Bearer token usually comes in format: "Bearer <token>"
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", errors.New("invalid token format")
	}

	token := parts[1]

	tokenData, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return s.JWTSecretKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := tokenData.Claims.(jwt.MapClaims); ok && tokenData.Valid {
		userIdFloat, ok := claims["user_id"].(float64)
		if !ok {
			return "", errors.New("invalid token")
		}

		userId := fmt.Sprintf("%.0f", userIdFloat)
		return userId, nil
	}

	return "", errors.New("invalid token")
}