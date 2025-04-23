package db

import (
	"context"
	"fmt"
	models "hausparty/libs/db/models/users"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DBType string

const (
	PostgresDB DBType = "postgres"
	MongoDB    DBType = "mongo"
)

type Config struct {
	Type        DBType
	PostgresDSN string
	MongoURI    string
	MongoDBName string
}

type DBFactory struct {
	pg     *gorm.DB
	mongo  *mongo.Database
	config Config
}

func NewDBFactory(ctx context.Context, cfg Config) (*DBFactory, error) {
	factory := &DBFactory{config: cfg}

	switch cfg.Type {
	case PostgresDB:
		// Initialize Postgres
		println("in factory", cfg.PostgresDSN)
		pg, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: false},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to postgres: %w", err)
		}
		factory.pg = pg
	case MongoDB:
		// Initialize MongoDB
		clientOpts := options.Client().ApplyURI(cfg.MongoURI)
		client, err := mongo.Connect(ctx, clientOpts)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to mongo: %w", err)
		}
		// Ping to verify
		ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err = client.Ping(ctxPing, nil); err != nil {
			return nil, fmt.Errorf("failed to ping mongo: %w", err)
		}
		factory.mongo = client.Database(cfg.MongoDBName)
	default:
		return nil, fmt.Errorf("unsupported DB type: %s", cfg.Type)
	}

	return factory, nil
}

// GetPostgres returns the GORM DB instance (nil if not initialized)
func (f *DBFactory) GetPostgres() *gorm.DB {
	return f.pg
}

// GetMongo returns the Mongo Database instance (nil if not initialized)
func (f *DBFactory) GetMongo() *mongo.Database {
	return f.mongo.Client().Database(f.config.MongoDBName)
}

// AutoMigrateIdentity runs GORM auto-migrations for identity models

func (f *DBFactory) AutoMigrateIdentity() error {
	if f.pg == nil {
		return nil
	}
	return f.pg.AutoMigrate(
		&models.User{},
		&models.AuthCredentials{},
		&models.Role{},
		&models.UserRole{},
		&models.Session{},
		&models.AdminAction{},
	)
}

// AutoMigrateParty runs GORM auto-migrations for party models

// func (f *DBFactory) AutoMigrateParty() error {
// 	if f.pg == nil {
// 		return nil
// 	}
// 	return f.pg.AutoMigrate(
// 		&models.Party{},
// 		&models.Attendee{},
// 		&models.Report{},
// 	)
// }
