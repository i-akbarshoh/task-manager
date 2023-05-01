package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/i-akbarshoh/task-manager/pkg/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Tokens struct to describe tokens object.
type Tokens struct {
	Access    string
	AccExpire int64
	Refresh   string
}

// GenerateNewTokens func for generate a new Access & Refresh token
func GenerateNewTokens(id string, credentials map[string]string) (*Tokens, error) {
	// Generate JWT Access token.
	accessToken, expire, err := generateNewAccessToken(id, credentials)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	// Generate JWT Refresh token.
	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Tokens{
		Access:    accessToken,
		Refresh:   refreshToken,
		AccExpire: expire,
	}, nil
}

func generateNewAccessToken(id string, credentials map[string]string) (string, int64, error) {
	claims := jwt.MapClaims{}

	claims["id"] = id
	claims["role"] = credentials["role"]

	// in local server access token ttl = 31 days
	if config.C.Environment == "development" {
		claims["expires"] = time.Now().Add(time.Hour * 24 * 31).Unix()
	} else {
		// in staging server access token ttl = a day
		claims["expires"] = time.Now().Add(time.Hour * time.Duration(config.C.JWT.Expire)).Unix()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.C.JWT.SigningKey))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", 0, err
	}

	return t, claims["expires"].(int64), nil
}

func generateNewRefreshToken() (string, error) {
	sha256Hash := sha256.New()

	refresh := config.C.JWT.RefreshKey + time.Now().String()

	_, err := sha256Hash.Write([]byte(refresh))
	if err != nil {
		return "", err
	}

	expireTime := fmt.Sprint(time.Now().Add(time.Hour * time.Duration(config.C.JWT.RExpire)).Unix())

	// Create a new refresh token (sha256 string with salt + expire time)
	t := hex.EncodeToString(sha256Hash.Sum(nil)) + "." + expireTime

	return t, nil
}
