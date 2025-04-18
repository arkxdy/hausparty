package utils

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

type OAuthUser struct {
	ID    string
	Email string
	Name  string
	// Add other fields from providers
}

type OAuthConfig struct {
	Config      *oauth2.Config
	UserInfoURL string
}

func ExchangeOAuthCode(provider, code string, config OAuthConfig) (*OAuthUser, error) {
	token, err := config.Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %v", err)
	}

	// Get user info - example for Google
	client := config.Config.Client(context.Background(), token)
	resp, err := client.Get(config.UserInfoURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	defer resp.Body.Close()

	// Parse response (implementation varies by provider)
	var user OAuthUser
	// ... parsing logic ...

	return &user, nil
}
