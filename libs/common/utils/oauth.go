package utils

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetOAuthConfig(provider string) *oauth2.Config {
	return oauthConfigs[provider]
}

var oauthConfigs = map[string]*oauth2.Config{
	"google": {
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		//RedirectURL:  "http://localhost:3001/api/v1/auth/google/callback",
		Scopes:   []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	},
	"twitter": {
		ClientID:     os.Getenv("TWITTER_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITTER_CLIENT_SECRET"),
		//RedirectURL:  "http://localhost:3001/api/v1/auth/twitter/callback",
		Scopes: []string{"tweet.read", "users.read"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://twitter.com/i/oauth2/authorize",
			TokenURL: "https://api.twitter.com/2/oauth2/token",
		},
	},
	"instagram": {
		ClientID:     os.Getenv("INSTAGRAM_CLIENT_ID"),
		ClientSecret: os.Getenv("INSTAGRAM_CLIENT_SECRET"),
		Scopes:       []string{"user_profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.instagram.com/oauth/authorize",
			TokenURL: "https://api.instagram.com/oauth/access_token",
		},
	},
}

func GetUserInfoURL(provider string) string {
	switch provider {
	case "google":
		return "https://www.googleapis.com/oauth2/v2/userinfo"
	case "instagram":
		return "https://graph.instagram.com/me?fields=id,username"
	case "twitter":
		// For Twitter v2 with OAuth2
		return "https://api.twitter.com/2/users/me"
	default:
		return ""
	}
}
