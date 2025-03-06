package main

import (
	"log"
	"os"

	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	routes.RegisterRoutes(router)

	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
  
