package services

import (
	"errors"
	"time"

	"github.com/o5h/config"
	"github.com/o5h/quiz/pkg/token"
)

type AuthService interface {
	Login(username, password string) (string, error)
}

type authService struct {
	tokenTimeout int
}

func NewAuthService() AuthService {
	return &authService{
		tokenTimeout: config.Get("auth.token_timeout", 15*60), // 15 minutes
	}
}

func (s *authService) Login(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("username and password must be provided")
	}

	token, err := token.CreateAccessToken(token.Claims{
		UserName: username,
		Role:     "user",
	}, time.Duration(s.tokenTimeout)*time.Second)

	return token, err
}
