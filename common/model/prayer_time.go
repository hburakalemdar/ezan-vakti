package model

import (
	"github.com/uptrace/bun"
)

type PrayerTime struct {
	bun.BaseModel `bun:"table:prayer_times"`

	ID     int64 `bun:",pk,autoincrement"`
	CityID int64 `json:"city_id"`

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

	City *City `bun:"rel:belongs-to,join:city_id=id"`
}
