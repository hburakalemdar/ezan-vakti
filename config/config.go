package config

type AppConfig struct {
	ApiBaseUrl string
	Email      string
	Password   string
}

func LoadConfig() AppConfig {
	return AppConfig{
		ApiBaseUrl: "https://awqatsalah.diyanet.gov.tr",
		Email:      "burakalemdar097@icloud.com",
		Password:   "eA6#8n?D",
	}
}
