package clients

import (
	"encoding/json"
	"fmt"
	"github.com/hayrat/ezan-vakti/api/viewmodel"
	"github.com/hayrat/ezan-vakti/common/model"
)

type ApiResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func GetLocations(apiUrl, token string, locationType model.LocationType, parentID *int64) ([]model.Location, []int64, error) {
	var endpoint string

	switch locationType {
	case model.LocationTypeCountry:
		endpoint = "/api/Place/Countries"
	case model.LocationTypeState:
		if parentID == nil {
			return nil, nil, fmt.Errorf("eyaletler için ülke ID'si gereklidir")
		}
		endpoint = fmt.Sprintf("/api/Place/States/%d", *parentID)
	case model.LocationTypeCity:
		if parentID == nil {
			return nil, nil, fmt.Errorf("şehirler için eyalet ID'si gereklidir")
		}
		endpoint = fmt.Sprintf("/api/Place/Cities/%d", *parentID)
	default:
		return nil, nil, fmt.Errorf("geçersiz konum tipi: %d", locationType)
	}

	url := BuildURL(apiUrl, endpoint)

	respBody, err := MakeAPIRequest("GET", url, token, nil)
	if err != nil {
		return nil, nil, err
	}

	var apiResp ApiResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, nil, fmt.Errorf("API yanıtı ayrıştırılamadı: %v", err)
	}
	if !apiResp.Success {
		return nil, nil, fmt.Errorf("API başarısız yanıt döndürdü: %s", apiResp.Message)
	}

	var responses []viewmodel.LocationResponse
	if err := json.Unmarshal(apiResp.Data, &responses); err != nil {
		return nil, nil, fmt.Errorf("lokasyon verileri ayrıştırılamadı: %v", err)
	}

	for i := range responses {
		responses[i].Type = int(locationType)
		if parentID != nil && locationType != model.LocationTypeCountry {
			responses[i].ParentID = *parentID
		}
	}

	locations := make([]model.Location, len(responses))
	apiIDs := make([]int64, len(responses))

	for i, resp := range responses {
		dbModel := resp.ToDBModel().(*model.Location)
		locations[i] = *dbModel
		apiIDs[i] = resp.ID
	}

	return locations, apiIDs, nil
}
