package main

import (
	"backend/config"
	expensehandler "backend/internal/expense/handler"
	expenserepo "backend/internal/expense/repo"
	expenseusecase "backend/internal/expense/usecase"
	"backend/internal/payment"
	"backend/internal/task"
	userhandler "backend/internal/user/handler"
	userrepo "backend/internal/user/repo"
	userusecase "backend/internal/user/usecase"
	"backend/pkg/httputil"
	"backend/pkg/token"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Configuration initialization.
	err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	cfg := config.Cfg

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

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	asynqClient := asynq.NewClientFromRedisClient(rdb)

	// Package initialization.
	token := token.NewToken()
	httpClient := httputil.NewClient()

	// Repo initialization.
	userRepo := userrepo.NewUserRepo(db)
	expenseRepo := expenserepo.NewExpenseRepo(db)

	// Usecase initialization.
	payment := payment.NewPayment(httpClient)
	task := task.NewTask(expenseRepo, payment)
	userUsecase := userusecase.NewUserUsecase(userRepo, token)
	expenseUsecase := expenseusecase.NewExpenseUsecase(expenseRepo, task, asynqClient)

	// Handler initialization.
	userHandler := userhandler.NewUserHandler(userUsecase)
	expenseHandler := expensehandler.NewExpenseHandler(expenseUsecase)

	// Router inititialization.
	router := gin.Default()

	api := router.Group("/api")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	userhandler.RegisterRoutes(api, userHandler)
	expensehandler.RegisterRoutes(api, expenseHandler)

	httpPort := ":" + cfg.HTTP.Port
	router.Run(httpPort)
}
