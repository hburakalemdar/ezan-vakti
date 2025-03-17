package models

type AuthResponse struct {
	Data struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	} `json:"data"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
