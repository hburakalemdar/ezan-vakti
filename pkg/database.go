package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/hayrat/ezan-vakti/common/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func InitDB(config DbConfig) {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		config.Host, config.Port, config.Name, config.Username, config.Password, config.SSLMode,
	)

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}

	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("Veritabanı bağlantısı test edilemedi: %v", err)
	}

	DB = bun.NewDB(sqlDB, pgdialect.New())

	log.Println("Veritabanı bağlantısı başarıyla kuruldu.")
}

func Migrate() {
	ctx := context.Background()

	models := []interface{}{
		(*model.Location)(nil),
		(*model.PrayerTime)(nil),
	}

	for _, m := range models {
		_, err := DB.NewDropTable().
			Model(m).
			IfExists().
			Cascade().
			Exec(ctx)
		if err != nil {
			log.Printf("Tablo silme hatası: %v", err)
		}
	}

	for _, m := range models {
		_, err := DB.NewCreateTable().
			Model(m).
			IfNotExists().
			Exec(ctx)
		if err != nil {
			log.Fatalf("Tablo işlemi sırasında hata: %v", err)
		}
	}

}
