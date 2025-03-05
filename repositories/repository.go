package repositories

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Repository defines the interface for  operations.
type BSRepository interface{}

// Repository implements the Repository interface.
type Repository struct {
	logger  *zap.Logger
	dbRead  *gorm.DB
	dbWrite *gorm.DB
}

// NewRepository creates a new instance of Repository.
func NewRepository(logger *zap.Logger, dbRead, dbWrite *gorm.DB) *Repository {
	return &Repository{
		logger:  logger,
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}
