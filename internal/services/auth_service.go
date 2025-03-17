package services

import (
	"errors"
	"github.com/hayrat/ezan-vakti/config"
	"github.com/hayrat/ezan-vakti/internal/api"
	"time"
)

type AuthService struct {
	Config       config.AppConfig
	AccessToken  string
	RefreshToken string
	TokenExpiry  time.Time
}

func NewAuthService(cfg config.AppConfig) *AuthService {
	return &AuthService{
		Config: cfg,
	}
}

func (a *AuthService) Authenticate() error {
	authResp, err := api.Login(a.Config.ApiBaseUrl, a.Config.Email, a.Config.Password)
	if err != nil || !authResp.Success {
		return errors.New("authentication failed")
	}

	a.AccessToken = authResp.Data.AccessToken
	a.RefreshToken = authResp.Data.RefreshToken
	a.TokenExpiry = time.Now().Add(30 * time.Minute) // 30 dakka

	return nil
}

func (a *AuthService) GetAccessToken() (string, error) {
	// Token geçerliyse mevcut token'i kullan
	if time.Now().Before(a.TokenExpiry.Add(-5 * time.Minute)) {
		return a.AccessToken, nil
	}

	// Token süresi dolmak üzereyse veya dolduysa yeniden kimlik doğrula
	err := a.Authenticate()
	if err != nil {
		return "", err
	}

	return a.AccessToken, nil
}
