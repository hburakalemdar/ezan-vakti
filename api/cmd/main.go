package main

import (
	"context"
	"log"
	"time"

	"github.com/hayrat/ezan-vakti/common/service"
	"github.com/hayrat/ezan-vakti/pkg"
	"github.com/uptrace/bun"
)

func main() {
	log.Println("Bismillah...")

	cfg, err := pkg.Setup()
	if err != nil {
		log.Fatalf("Yapılandırma yüklenemedi: %v", err)
	}

	pkg.InitDB(cfg.Database)

	db := pkg.DB
	if db == nil {
		log.Fatalf("Veritabanı bağlantısı kurulamadı")
	}

	updateData(cfg, db)

	ticker := time.NewTicker(cfg.UpdatePeriod)
	defer ticker.Stop()

	log.Printf("Düzenli güncelleme zamanlayıcısı başlatıldı (Periyot: %s)", cfg.UpdatePeriod)
	for {
		select {
		case <-ticker.C:
			log.Println("Zamanlanmış güncelleme başlatılıyor...")
			updateData(cfg, db)
		}
	}
}

func updateData(cfg *pkg.Config, db *bun.DB) {
	ctx := context.Background()
	log.Println("Veri güncelleme işlemi başlıyor...")

	authService := service.NewAuthService(cfg.API)
	token, err := authService.GetAccessToken()
	if err != nil {
		log.Printf("Token alınamadı, güncelleme yapılamayacak: %v", err)
		return
	}

	placeService := service.NewPlaceService(db, cfg.API.BaseUrl, token)
	prayerService := service.NewPrayerService(db, cfg.API.BaseUrl, token, placeService)

	pkg.Migrate()

	cityIDs, err := placeService.SyncLocations(ctx)
	if err != nil {
		log.Printf("Konum verileri güncellenirken hata: %v", err)
		return
	}
	if err := prayerService.SyncPrayerTimes(ctx, cityIDs); err != nil {
		log.Printf("Namaz vakitleri güncellenirken hata: %v", err)
	}

	log.Printf("Bir sonraki güncelleme %s sonra yapılacak: %s",
		cfg.UpdatePeriod,
		time.Now().Add(cfg.UpdatePeriod).Format("2006-01-02 15:04:05"))
}
