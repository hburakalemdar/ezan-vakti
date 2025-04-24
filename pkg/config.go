package pkg

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
)

var config *Config
var DB *bun.DB

type Config struct {
	API          ApiConfig
	Database     DbConfig
	UpdatePeriod time.Duration
}

type ApiConfig struct {
	BaseUrl  string
	Email    string
	Password string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
	SSLMode  string
}

func setDefaults() {

	viper.SetDefault("api.baseUrl", "https://awqatsalah.diyanet.gov.tr")
	viper.SetDefault("api.email", "burakalemdar097@icloud.com")
	viper.SetDefault("api.password", "eA6#8n?D")

	viper.SetDefault("database.name", "ezanvakti")
	viper.SetDefault("database.username", "burakalemdar")
	viper.SetDefault("database.password", "burakalemdar")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.sslmode", "disable")

	viper.SetDefault("updatePeriod", 29*24*time.Hour)
}

func bindEnvs() {
	bindEnv := func(input ...string) {
		err := viper.BindEnv(input...)
		if err != nil {
			log.Printf("Çevre değişkeni bağlanamadı: %v", err)
		}
	}

	bindEnv("api.baseUrl", "API_BASE_URL")
	bindEnv("api.email", "API_EMAIL")
	bindEnv("api.password", "API_PASSWORD")

	bindEnv("database.name", "DB_NAME")
	bindEnv("database.username", "DB_USERNAME")
	bindEnv("database.password", "DB_PASSWORD")
	bindEnv("database.host", "DB_HOST")
	bindEnv("database.port", "DB_PORT")
	bindEnv("database.sslmode", "DB_SSLMODE")

	bindEnv("updatePeriod", "UPDATE_PERIOD")
}

func Setup() (*Config, error) {
	setDefaults()
	bindEnvs()

	viper.AutomaticEnv()

	config = &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("yapılandırma yüklenemedi: %v", err)
	}

	return config, nil
}
