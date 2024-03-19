package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/harsha-s/geoip/ipstack"
)

func main() {
	var ipAddress string
	flag.StringVar(&ipAddress, "i", "", "IP Address to lookup")
	flag.Parse()

	if ipAddress == "" {
		fmt.Println("Usage: ./geoip -i <IP Address>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	apiKey := os.Getenv("IPSTACK_API_KEY") // Get the API key from the environment variable
	if apiKey == "" {
		log.Fatal("IPSTACK_API_KEY environment variable is not set.")
	}

	client := ipstack.IPStackClient{APIKey: apiKey} // Create client with your key
	data, err := client.GetGeolocation(ipAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Latitude: %.6f, Longitude: %.6f\n", data.Latitude, data.Longitude)
}
