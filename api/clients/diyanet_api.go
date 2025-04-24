package clients

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func MakeAPIRequest(method, url, token string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("istek oluşturulurken hata: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	client := &http.Client{
		Timeout: 180 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("API isteğinde hata: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("yanıt okunamadı: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API hata kodu döndürdü: %d, yanıt: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func BuildURL(apiURL, endpoint string, params ...interface{}) string {
	if len(params) > 0 {
		endpoint = fmt.Sprintf(endpoint, params...)
	}
	return fmt.Sprintf("%s%s", apiURL, endpoint)
}
