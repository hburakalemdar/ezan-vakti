package service

import (
	"context"
	"fmt"

	"github.com/hayrat/ezan-vakti/api/clients"
	"github.com/uptrace/bun"
)

type PrayerService struct {
	DB           *bun.DB
	ApiUrl       string
	Token        string
	PlaceService *PlaceService
}

func NewPrayerService(db *bun.DB, apiUrl, token string, placeService *PlaceService) *PrayerService {
	return &PrayerService{
		DB:           db,
		ApiUrl:       apiUrl,
		Token:        token,
		PlaceService: placeService,
	}
}

func (s *PrayerService) SyncPrayerTimes(ctx context.Context, cityApiIDs []int64) error {
	for _, cityID := range cityApiIDs {
		prayerTimes, err := clients.GetMonthlyPrayerTimes(s.ApiUrl, s.Token, cityID)
		if err != nil {
			fmt.Printf("Şehir ID: %d için namaz vakitleri çekilirken hata: %v\n", cityID, err)
			continue
		}
		if len(prayerTimes) == 0 {
			continue
		}

		_, err = s.DB.NewInsert().
			Model(&prayerTimes).
			Exec(ctx)

		if err != nil {
			fmt.Printf("Şehir ID: %d için namaz vakitleri kaydedilirken hata: %v\n", cityID, err)
		}
	}

	return nil
}
