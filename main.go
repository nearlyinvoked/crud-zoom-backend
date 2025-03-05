package main

import (
	config "crud-zoom/config"
	database "crud-zoom/database"
	"crud-zoom/routes"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
    r := gin.Default()

    logger, _ := zap.NewProduction()
    defer logger.Sync()

    logger.Info("Initiating Service....")

    cfg, err := config.NewConfig(logger)
    if err != nil {
        logger.Error(err.Error())
    }

    database.Init(&cfg)

    // CORS configuration
    corsConfig := cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
        AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }

    r.Use(cors.New(corsConfig))

    routes.Init(r, cfg, logger)

    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}