package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hayrat/ezan-vakti/common/model"
)

func Login(apiUrl, email, password string) (model.AuthResponse, error) {
	loginUrl := BuildURL(apiUrl, "/Auth/Login")

	requestData := map[string]string{
		"email":    email,
		"password": password,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return model.AuthResponse{}, fmt.Errorf("JSON'a dönüştürme hatası: %v", err)
	}

	respBody, err := MakeAPIRequest("POST", loginUrl, "", bytes.NewBuffer(jsonData))
	if err != nil {
		return model.AuthResponse{}, err
	}

	var authResp model.AuthResponse
	if err := json.Unmarshal(respBody, &authResp); err != nil {
		return model.AuthResponse{}, fmt.Errorf("yanıt ayrıştırma hatası: %v", err)
	}

	if !authResp.Success {
		return authResp, fmt.Errorf("giriş başarısız: %s", authResp.Message)
	}

	log.Printf("Giriş başarılı")
	return authResp, nil
}
