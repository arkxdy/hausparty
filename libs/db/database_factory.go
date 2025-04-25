package db

import (
	"context"
	"errors"
	"fmt"
	partyModels "hausparty/libs/db/models/party"
	models "hausparty/libs/db/models/users"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
func (f *DBFactory) AutoMigrateParty() error {
	if f.pg == nil {
		return nil
	}
	return f.pg.AutoMigrate(
		&partyModels.Party{},
		&partyModels.Attendee{},
		&partyModels.Report{},
	)
}

// AutoMigrateRating runs GORM auto-migrations for rating models
func (f *DBFactory) AutoSetupRatingMongo() error {
	if f.mongo == nil {
		return errors.New("MongoDB client is nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ratingColl := f.mongo.Collection(f.config.MongoDBName)

	// Ensure indexes
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "party_id", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "user_id", Value: 1}},
		},
		{
			Keys: bson.D{
				{Key: "party_id", Value: 1},
				{Key: "user_id", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err := ratingColl.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		return fmt.Errorf("failed to create indexes: %w", err)
	}

	log.Println("MongoDB 'ratings' collection and indexes set up successfully")
	return nil
}
