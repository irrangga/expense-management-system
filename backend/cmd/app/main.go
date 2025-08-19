package main

import (
	"backend/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Configuration initialization.
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Router inititialization.
	router := gin.Default()

	api := router.Group("/api")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	httpPort := ":" + cfg.HTTP.Port
	router.Run(httpPort)
}
