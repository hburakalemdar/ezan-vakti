package viewmodel

import "github.com/hayrat/ezan-vakti/common/model"

type DBConvertible interface {
	ToDBModel() any
}

type CountryResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func (vm *CountryResponse) ToDBModel() any {
	return &model.Country{
		ApiID: vm.ID,
		Name:  vm.Name,
		Code:  vm.Code,
	}
}

type StateResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	CountryID int64  `json:"country_id"`
}

func (vm *StateResponse) ToDBModel() any {
	return &model.State{
		ApiID:     vm.ID,
		Name:      vm.Name,
		Code:      vm.Code,
		CountryID: vm.CountryID,
	}
}

type CityResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	StateID int64  `json:"state_id"`
}

func (vm *CityResponse) ToDBModel() any {
	return &model.City{
		ApiID:   vm.ID,
		Name:    vm.Name,
		Code:    vm.Code,
		StateID: vm.StateID,
	}
}
