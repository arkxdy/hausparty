package db

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

var (
	once        sync.Once
	factoryInst *DBFactory
	factoryErr  error
)

func Connect(ctx context.Context) (*DBFactory, error) {
	once.Do(func() {
		// Load local env in non-prod
		if os.Getenv("ENV") != "production" {
			_ = godotenv.Load(filepath.Join("..", "..", ".env.dev"))
		}
		svc := os.Getenv("SERVICE_NAME")
		dbType := DBType(os.Getenv("DB_TYPE"))

		switch dbType {
		case PostgresDB:
			key := fmt.Sprintf("%s_DATABASE_URL", strings.ToUpper(svc))
			dsn := os.Getenv(key)
			if dsn == "" {
				factoryErr = fmt.Errorf("DATABASE_URL not set")
				return
			}
			cfg := Config{
				Type:        PostgresDB,
				PostgresDSN: dsn,
			}
			factoryInst, factoryErr = NewDBFactory(ctx, cfg)

		case MongoDB:
			uri := os.Getenv("MONGO_URI")
			dbName := os.Getenv("MONGO_DB_NAME")
			if uri == "" || dbName == "" {
				factoryErr = fmt.Errorf("MONGO_URI or MONGO_DB_NAME not set")
				return
			}
			cfg := Config{
				Type:        MongoDB,
				MongoURI:    uri,
				MongoDBName: dbName,
			}
			factoryInst, factoryErr = NewDBFactory(ctx, cfg)

		default:
			factoryErr = fmt.Errorf("unsupported DB_TYPE: %s", dbType)
		}
	})
	return factoryInst, factoryErr
}
