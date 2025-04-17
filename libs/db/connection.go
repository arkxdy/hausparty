package db

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

// Connect establishes a connection to PostgreSQL using GORM
func Connect() (*gorm.DB, error) {
	var err error

	once.Do(func() {
		if os.Getenv("ENV") != "production" {
			_ = godotenv.Load(filepath.Join("..", "..", ".env.dev"))
		}

		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			err = fmt.Errorf("DATABASE_URL not set")
			return
		}

		db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if dbErr != nil {
			err = fmt.Errorf("failed to connect to database: %w", dbErr)
			return
		}

		sqlDB, sqlErr := db.DB()
		if sqlErr != nil {
			err = fmt.Errorf("failed to get DB instance: %w", sqlErr)
			return
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		log.Println("[DB] connected successfully")
		dbInstance = db
	})

	return dbInstance, err
}
