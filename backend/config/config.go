package config

import (
	"log"
	"os"
)

// GetAPIBaseURL returns the base URL for the API
func GetAPIBaseURL() string {
	baseURL := os.Getenv("SNCF_API_URL")
	if baseURL == "" {
		log.Fatalf("SNCF_API_URL is not set")
	}
	return baseURL
}
