package handlers

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"hauparty/services/auth-service/services"
	"hausparty/libs/common/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	authService services.AuthService
}

var (
	verifier, challenge = generateCodeVerifier()
)

func NewAuthHandler(s services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: s,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	// Implementation here
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	// Implementation here
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (h *AuthHandler) AuthProvider(c *gin.Context) {
	provider := c.Param("provider")
	// Validate incoming request
	var req AuthRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
		return
	}
	oauthConfig := utils.GetOAuthConfig(provider)
	println("provider", provider)
	if oauthConfig.ClientID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported provider"})
		return
	}
	// authURL := oauthConfig.AuthCodeURL(req.State,
	// 	oauth2.SetAuthURLParam("redirect_uri", req.RedirectURI),
	// )
	println("req", req.State)
	//authURL := oauthConfig.AuthCodeURL(req.State)

	authURL := oauthConfig.AuthCodeURL(req.State,
		oauth2.SetAuthURLParam("redirect_uri", req.RedirectURI),
		oauth2.SetAuthURLParam("code_challenge", challenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)
	println("auth URL", authURL)
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

type AuthRequest struct {
	RedirectURI string `form:"redirect_uri" binding:"required"`
	State       string `form:"state" binding:"required"`
}

func (h *AuthHandler) AuthCallback(c *gin.Context) {
	println("In Callback", c.Param("provider"))
	println("🔁 Callback URL:", c.Request.URL.String())
	provider := c.Param("provider") // or c.Query("provider")
	code := c.Query("code")
	redirect_uri := c.Query("redirect_uri")
	config := utils.GetOAuthConfig(provider)
	if config.ClientID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported provider"})
		return
	}
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}
	config.RedirectURL = redirect_uri
	println("code", code)
	println("code", config.RedirectURL)

	// token, err := config.Exchange(context.Background(), code)
	token, err := config.Exchange(context.Background(), code,
		oauth2.SetAuthURLParam("code_verifier", verifier),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token exchange failed", "details": err.Error()})
		return
	}
	println("token", token)
	client := config.Client(context.Background(), token)

	userInfoURL := utils.GetUserInfoURL(provider)
	resp, err := client.Get(userInfoURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user info"})
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode user info"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "provider": provider, "user": userInfo})
}

func generateCodeVerifier() (string, string) {
	// 32 bytes is enough
	verifierBytes := make([]byte, 32)
	_, err := rand.Read(verifierBytes)
	if err != nil {
		panic(err)
	}
	verifier := base64.RawURLEncoding.EncodeToString(verifierBytes)

	// SHA256 it for the challenge
	hash := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(hash[:])

	return verifier, challenge
}
