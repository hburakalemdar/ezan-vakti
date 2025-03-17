package main

import (
	"fmt"
	"log"

	"github.com/hayrat/ezan-vakti/config"
	"github.com/hayrat/ezan-vakti/internal/services"
)

func main() {
	cfg := config.LoadConfig()

	authService := services.NewAuthService(cfg)

	token, err := authService.GetAccessToken()
	if err != nil {
		log.Fatalf("Token alınamadı: %v", err)
	}

	fmt.Println("Elde edilen JWT Token:", token)
}
