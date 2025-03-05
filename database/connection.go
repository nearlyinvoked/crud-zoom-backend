package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"crud-zoom/config"
)

var (
    db  *gorm.DB
    err error
)

const MAX_IDLE_CONNECTIONS = 10
const MAX_OPEN_CONNECTIONS = 10
const MAX_LIFETIMES = time.Hour

func Init(cfg *config.Config) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
        cfg.DBHost, cfg.DBUser, cfg.DBPasswd, cfg.DBName, cfg.DBPort)

    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Warn),
    })
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
        return
    }

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Failed to get database instance: %v", err)
    }

    sqlDB.SetMaxIdleConns(MAX_IDLE_CONNECTIONS)
    sqlDB.SetMaxOpenConns(MAX_OPEN_CONNECTIONS)
    sqlDB.SetConnMaxLifetime(MAX_LIFETIMES)
}

func GetReadDB() *gorm.DB {
	return db
}

func GetWriteDB() *gorm.DB {
	return db
}
