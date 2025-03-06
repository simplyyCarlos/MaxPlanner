package controllers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSNCFData handles API requests
func GetSNCFData(c *gin.Context) {
	data, err := services.FetchSNCFData(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
