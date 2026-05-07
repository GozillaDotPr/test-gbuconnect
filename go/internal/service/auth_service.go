package service

import (
	"errors"
	"time"

	"github.com/app/gin-postgres-api/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Login(username, password string) (string, error)
	ValidateToken(tokenStr string) (jwt.MapClaims, error)
}

type AuthServiceImpl struct {
	cfg *config.Config
}

func NewAuthService(cfg *config.Config) AuthService {
	return &AuthServiceImpl{cfg: cfg}
}

func (s *AuthServiceImpl) Login(username, password string) (string, error) {
	if username != s.cfg.AppUsername || password != s.cfg.AppPassword {
		return "", errors.New("invalid credentials")
	}

	expires := s.cfg.JWTExpires
	if expires == 0 {
		expires = 3600
	}

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Duration(expires) * time.Second).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWTSecret))
}

func (s *AuthServiceImpl) ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.cfg.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
