package ipstack

import (
	"os"
	"testing"
)

func TestGetGeolocation(t *testing.T) {

	apiKey := os.Getenv("IPSTACK_API_KEY") // Get the API key from the environment variable
	if apiKey == "" {
		t.Fatal("IPSTACK_API_KEY environment variable is not set.")
	}

	client := &IPStackClient{APIKey: apiKey}
	data, err := client.GetGeolocation("8.8.8.8")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if data.Latitude != 40.5369987487793 || data.Longitude != -82.12859344482422 {
		t.Errorf("Unexpected GeolocationData: %+v", data)
	}
}
