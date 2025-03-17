package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hayrat/ezan-vakti/internal/models"
)

func Login(apiUrl, email, password string) (models.AuthResponse, error) {
	loginUrl := fmt.Sprintf("%s/Auth/Login", apiUrl)

	reqBody, _ := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})

	resp, err := http.Post(loginUrl, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return models.AuthResponse{}, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("Hata: Response body kapatılırken sorun oluştu:", err)
		}
	}()

	var authResp models.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {
		return models.AuthResponse{}, err
	}

	return authResp, nil
}
