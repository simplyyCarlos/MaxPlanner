package services

import (
	"backend/config"
	"backend/utils"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

// FetchSNCFData requests data from the SNCF API with dynamic parameters
func FetchSNCFData(c *gin.Context) (interface{}, error) {
	baseURL := config.GetAPIBaseURL()
	params := url.Values{}

	// Get query params from request, or set defaults
	params.Set("select", "*")
	params.Set("where", c.DefaultQuery("where", "(origine like \"Paris\" OR origine like \"BOURG EN BRESSE\") AND (destination like \"PARIS\" OR destination like \"BOURG EN BRESSE\")"))
	params.Set("order_by", c.DefaultQuery("order_by", "date DESC"))
	params.Set("limit", c.DefaultQuery("limit", "100"))

	// Construct full URL
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Make HTTP request
	response, err := utils.MakeRequest(fullURL)
	if err != nil {
		return nil, err
	}

	return response, nil
}
