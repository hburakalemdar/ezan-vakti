package service

import (
	"context"

	"github.com/hayrat/ezan-vakti/api/clients"
	"github.com/hayrat/ezan-vakti/common/model"
	"github.com/uptrace/bun"
)

type PlaceService struct {
	DB     *bun.DB
	ApiUrl string
	Token  string
}

func NewPlaceService(db *bun.DB, apiUrl, token string) *PlaceService {
	return &PlaceService{
		DB:     db,
		ApiUrl: apiUrl,
		Token:  token,
	}
}

func (s *PlaceService) SaveLocationsToDB(ctx context.Context, locations []model.Location) error {
	if len(locations) == 0 {
		return nil
	}

	_, err := s.DB.NewInsert().
		Model(&locations).
		On("CONFLICT (api_id, type) DO UPDATE").
		Set("name = EXCLUDED.name").
		Set("code = EXCLUDED.code").
		Set("parent_id = EXCLUDED.parent_id").
		Set("created_at = CURRENT_TIMESTAMP").
		Exec(ctx)

	return err
}

func (s *PlaceService) SyncLocations(ctx context.Context) ([]int64, error) {
	var allStates, allCities []model.Location
	var stateIDs []int64
	var allCityIDs []int64

	countries, countryIDs, err := clients.GetLocations(s.ApiUrl, s.Token, model.LocationTypeCountry, nil)
	if err != nil {
		return nil, err
	}

	if err := s.SaveLocationsToDB(ctx, countries); err != nil {
		return nil, err
	}

	for _, countryID := range countryIDs {
		states, ids, err := clients.GetLocations(s.ApiUrl, s.Token, model.LocationTypeState, &countryID)
		if err != nil {
			continue
		}

		allStates = append(allStates, states...)
		stateIDs = append(stateIDs, ids...)
	}

	if err := s.SaveLocationsToDB(ctx, allStates); err != nil {
		return nil, err
	}

	for _, stateID := range stateIDs {
		cities, cityIDs, err := clients.GetLocations(s.ApiUrl, s.Token, model.LocationTypeCity, &stateID)
		if err != nil {
			continue
		}

		allCities = append(allCities, cities...)
		allCityIDs = append(allCityIDs, cityIDs...)
	}

	if err := s.SaveLocationsToDB(ctx, allCities); err != nil {
		return nil, err
	}

	return allCityIDs, nil
}
