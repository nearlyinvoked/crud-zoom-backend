package routes

import (
	config "crud-zoom/config"
	database "crud-zoom/database"
	repositories "crud-zoom/repositories"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type route struct {
	cfg       config.Config
	logger    *zap.Logger
	repo      *repositories.Repository
}

func newRoute(cfg config.Config, logger *zap.Logger) *route {
	return &route{
		cfg:       cfg,
		logger:    logger,
		repo:      repositories.NewRepository(logger, database.GetReadDB(), database.GetWriteDB()),
	}
}

func Init(r *gin.Engine, cfg config.Config, logger *zap.Logger) {

	route := newRoute(cfg, logger)
	app := r.Group("/v1")

	// init routes
	route.initZoom(app)     // /v1/zoom
}
