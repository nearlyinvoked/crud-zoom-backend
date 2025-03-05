package main

import (
	config "crud-zoom/config"
	database "crud-zoom/database"
	"crud-zoom/routes"

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

	routes.Init(r, cfg, logger)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
