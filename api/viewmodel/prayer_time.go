package viewmodel

import "github.com/hayrat/ezan-vakti/common/model"

type PrayerTimeResponse struct {
	Fajr                      string  `json:"fajr"`
	Sunrise                   string  `json:"sunrise"`
	Dhuhr                     string  `json:"dhuhr"`
	Asr                       string  `json:"asr"`
	Maghrib                   string  `json:"maghrib"`
	Isha                      string  `json:"isha"`
	AstronomicalSunset        string  `json:"astronomicalSunset"`
	AstronomicalSunrise       string  `json:"astronomicalSunrise"`
	HijriDateShort            string  `json:"hijriDateShort"`
	HijriDateShortIso8601     *string `json:"hijriDateShortIso8601"`
	HijriDateLong             string  `json:"hijriDateLong"`
	HijriDateLongIso8601      *string `json:"hijriDateLongIso8601"`
	QiblaTime                 string  `json:"qiblaTime"`
	GregorianDateShort        string  `json:"gregorianDateShort"`
	GregorianDateShortIso8601 string  `json:"gregorianDateShortIso8601"`
	GregorianDateLong         string  `json:"gregorianDateLong"`
	GregorianDateLongIso8601  string  `json:"gregorianDateLongIso8601"`
	GreenwichMeanTimeZone     int     `json:"greenwichMeanTimeZone"`
	CityID                    int64   `json:"city_id"`
}

func (vm *PrayerTimeResponse) ToDBModel() any {
	return &model.PrayerTime{
		Fajr:                      vm.Fajr,
		Sunrise:                   vm.Sunrise,
		Dhuhr:                     vm.Dhuhr,
		Asr:                       vm.Asr,
		Maghrib:                   vm.Maghrib,
		Isha:                      vm.Isha,
		AstronomicalSunset:        vm.AstronomicalSunset,
		AstronomicalSunrise:       vm.AstronomicalSunrise,
		HijriDateShort:            vm.HijriDateShort,
		HijriDateShortIso8601:     vm.HijriDateShortIso8601,
		HijriDateLong:             vm.HijriDateLong,
		HijriDateLongIso8601:      vm.HijriDateLongIso8601,
		QiblaTime:                 vm.QiblaTime,
		GregorianDateShort:        vm.GregorianDateShort,
		GregorianDateShortIso8601: vm.GregorianDateShortIso8601,
		GregorianDateLong:         vm.GregorianDateLong,
		GregorianDateLongIso8601:  vm.GregorianDateLongIso8601,
		GreenwichMeanTimeZone:     vm.GreenwichMeanTimeZone,
		CityID:                    vm.CityID,
	}
}
