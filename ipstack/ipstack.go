package ipstack

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IPStackClient struct {
	APIKey string
}

type GeolocationData struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (client *IPStackClient) GetGeolocation(ip_address string) (*GeolocationData, error) {
	url := fmt.Sprintf("http://api.ipstack.com/%s?access_key=%s", ip_address, client.APIKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("IPStack API error: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data GeolocationData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
