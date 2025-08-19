package main

import (
	"backend/config"
	userhandler "backend/internal/user/handler"
	userrepo "backend/internal/user/repo"
	userusecase "backend/internal/user/usecase"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Configuration initialization.
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Database initialization.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Name,
		cfg.Postgres.Port,
		cfg.Postgres.SslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Repo initialization.
	userRepo := userrepo.NewUserRepo(db)

	// Usecase initialization.
	userUsecase := userusecase.NewUserUsecase(userRepo)

	// Handler initialization.
	userHandler := userhandler.NewUserHandler(userUsecase)

	// Router inititialization.
	router := gin.Default()

	api := router.Group("/api")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	userhandler.RegisterRoutes(api, userHandler)

	httpPort := ":" + cfg.HTTP.Port
	router.Run(httpPort)
}
