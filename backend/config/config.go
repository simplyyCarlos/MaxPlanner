package config

import "os"

// GetAPIBaseURL returns the base URL for the API
func GetAPIBaseURL() string {
	baseURL := os.Getenv("SNCF_API_URL")
	if baseURL == "" {
		baseURL = "https://ressources.data.sncf.com/api/explore/v2.1/catalog/datasets/tgvmax/records"
	}
	return baseURL
}
