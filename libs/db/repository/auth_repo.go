package repository

import (
	"context"
	"encoding/json"
	"fmt"
	models "hausparty/libs/db/models/users"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	CreateAuthCredentials(ctx context.Context, creds *models.AuthCredentials) error
	UpdatePassword(ctx context.Context, userID string, newHash string) error
	GetCredentialsByOAuthID(ctx context.Context, provider, oauthID string) (*models.AuthCredentials, error)
	GetCredentialsByEmail(ctx context.Context, email string) (*models.AuthCredentials, error)
	UpdateEmail(ctx context.Context, userID string, newEmail string) error
}

type authRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

// GetCredentialsByEmail implements IAuthRepository.
func (r *authRepository) GetCredentialsByEmail(ctx context.Context, email string) (*models.AuthCredentials, error) {
	panic("unimplemented")
}

// UpdateEmail implements IAuthRepository.
func (r *authRepository) UpdateEmail(ctx context.Context, userID string, newEmail string) error {
	panic("unimplemented")
}

func (r *authRepository) CreateAuthCredentials(ctx context.Context, creds *models.AuthCredentials) error {
	return r.db.WithContext(ctx).Create(creds).Error
}

func (r *authRepository) UpdatePassword(ctx context.Context, userID string, newHash string) error {
	// First update DB
	err := r.db.WithContext(ctx).
		Model(&models.AuthCredentials{}).
		Where("user_id = ?", userID).
		Update("password_hash", newHash).Error

	// Then invalidate related cache entries
	r.redis.Del(ctx,
		fmt.Sprintf("user:%s:credentials", userID),
		fmt.Sprintf("user:%s:sessions", userID),
	)

	return err
}

func (r *authRepository) GetCredentialsByOAuthID(ctx context.Context, provider string, oauthID string) (*models.AuthCredentials, error) {
	cacheKey := fmt.Sprintf("oauth:%s:%s", provider, oauthID)

	// Try cache first
	cachedVal, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var creds models.AuthCredentials
		if err := json.Unmarshal([]byte(cachedVal), &creds); err == nil {
			return &creds, nil
		}
	}

	// Cache miss - query DB
	var creds models.AuthCredentials
	err = r.db.WithContext(ctx).
		Where("oauth_provider = ? AND oauth_id = ?", provider, oauthID).
		First(&creds).Error

	if err != nil {
		return nil, err
	}

	// Cache result with TTL
	serialized, _ := json.Marshal(creds)
	r.redis.Set(ctx, cacheKey, serialized, 15*time.Minute) // TTL = 15m

	return &creds, nil
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return &authRepository{db: db}
}
