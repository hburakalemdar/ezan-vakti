package clients

import (
	"encoding/json"
	"fmt"
	"github.com/hayrat/ezan-vakti/api/viewmodel"
	"github.com/hayrat/ezan-vakti/common/model"
)

func GetMonthlyPrayerTimes(apiUrl, token string, cityId int64) ([]*model.PrayerTime, error) {
	url := BuildURL(apiUrl, "/api/PrayerTime/Monthly/%d", cityId)

	respBody, err := MakeAPIRequest("GET", url, token, nil)
	if err != nil {
		return nil, err
	}

	var apiResp ApiResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("API yanıtı ayrıştırılamadı: %v", err)
	}
	if !apiResp.Success {
		return nil, fmt.Errorf("API başarısız yanıt döndürdü: %s", apiResp.Message)
	}

	var prayerTimeResps []viewmodel.PrayerTimeResponse
	if err := json.Unmarshal(apiResp.Data, &prayerTimeResps); err != nil {
		return nil, fmt.Errorf("namaz vakti verileri ayrıştırılamadı: %v", err)
	}

	prayerTimes := make([]*model.PrayerTime, len(prayerTimeResps))
	for i, resp := range prayerTimeResps {
		resp.CityID = cityId
		prayerTimes[i] = resp.ToDBModel().(*model.PrayerTime)
	}

	return prayerTimes, nil
}
