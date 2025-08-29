package main

import (
	"backend/config"
	expenserepo "backend/internal/expense/repo"
	"backend/internal/payment"
	"backend/internal/task"
	"backend/pkg/constant"
	"backend/pkg/httputil"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// workers.go
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

	srv := asynq.NewServerFromRedisClient(
		rdb,
		asynq.Config{Concurrency: 10},
	)

	// Package initialization.
	httpClient := httputil.NewClient()

	// Repo initialization.
	expenseRepo := expenserepo.NewExpenseRepo(db)

	// Usecase initialization.
	payment := payment.NewPayment(httpClient)
	task := task.NewTask(expenseRepo, payment)

	mux := asynq.NewServeMux()
	mux.HandleFunc(constant.PaymentType, task.HandlePaymentProcessTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
