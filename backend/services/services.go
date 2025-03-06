package services

import (
	"backend/config"
	"backend/utils"
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

func FetchSNCFData(c *gin.Context) (interface{}, error) {
	baseURL, err := config.GetAPIBaseURL()
	if err != nil {
		return nil, err
	}
	params := url.Values{}

	params.Set("select", "*")
	params.Set("where", c.DefaultQuery("where", "(origine like \"Paris\" OR origine like \"BOURG EN BRESSE\") AND (destination like \"PARIS\" OR destination like \"BOURG EN BRESSE\")"))
	params.Set("order_by", c.DefaultQuery("order_by", "date DESC"))
	params.Set("limit", c.DefaultQuery("limit", "100"))

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	response, err := utils.MakeRequest(fullURL)
	if err != nil {
		return nil, err
	}

	return response, nil
}
