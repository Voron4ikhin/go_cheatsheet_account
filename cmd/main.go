package main

import (
	"log"

	"account/internal/config"
	"account/internal/logger"
	"account/internal/repository"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialization logger with server name
	logger := logger.New(cfg)

	db, err := gorm.Open(postgres.Open(cfg.DbDsn), &gorm.Config{})
	if err != nil {
		logger.Error().Msgf("Failed to connect to database: %v", err)
		return
	}
	logger.Info().Msg("Connected to database")

	// Repository initialization
	repo := repository.NewRepository(db, &logger)
	_ = repo

	router := gin.Default()
	err = router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		logger.Error().Msgf("Failed to set trusted proxies: %v", err)
		return
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", PingExample)
	router.Run(":8080")

	logger.Info().Msg("Service starting up")
}

// PingExample godoc
// @Summary 		Проверка доступности сервиса
// @Description 	Возвращает pong
// @Tags 			health
// @Success			200 {string} string "pong"
// @Router			/ping [get]
func PingExample(c *gin.Context) {
	c.String(200, "pong")
}
