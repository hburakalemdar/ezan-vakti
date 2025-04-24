package service

import (
	"errors"

	"github.com/hayrat/ezan-vakti/api/clients"
	"github.com/hayrat/ezan-vakti/pkg"
)

type AuthService struct {
	Config      pkg.ApiConfig
	AccessToken string
}

func NewAuthService(cfg pkg.ApiConfig) *AuthService {
	return &AuthService{
		Config: cfg,
	}
}

func (a *AuthService) Authenticate() error {
	authResp, err := clients.Login(a.Config.BaseUrl, a.Config.Email, a.Config.Password)
	if err != nil || !authResp.Success {
		return errors.New("authentication failed")
	}

	a.AccessToken = authResp.Data.AccessToken
	return nil
}

func (a *AuthService) GetAccessToken() (string, error) {
	if a.AccessToken != "" {
		return a.AccessToken, nil
	}

	err := a.Authenticate()
	if err != nil {
		return "", err
	}

	return a.AccessToken, nil
}
