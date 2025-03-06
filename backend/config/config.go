package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// GetAPIBaseURL returns the base URL for the API
func GetAPIBaseURL() (string, error) {
	err := godotenv.Load()
	if err != nil {
		err = fmt.Errorf("Error loading .env file : %v ", err)
		return "", err
	}

	baseURL := os.Getenv("SNCF_API_URL")
	return baseURL, nil

}
