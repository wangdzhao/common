package secret

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TokenService struct for handling tokens with a specific key
type TokenService struct {
	jwtKey []byte
}

// NewTokenService creates a new instance of TokenService with the given key
func NewTokenService(key string) *TokenService {
	return &TokenService{
		jwtKey: []byte(key),
	}
}

// Claims struct to embed user information in the token
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// GenerateTokenPair generates a new JWT AccessToken and RefreshToken for a user
func (ts *TokenService) GenerateTokenPair(userID string) (accessToken string, refreshToken string, err error) {
	// AccessToken expires after 24 hours
	accessToken, err = ts.generateToken(userID, 24*time.Hour)
	if err != nil {
		return "", "", err
	}

	// RefreshToken expires after 30 days
	refreshToken, err = ts.generateToken(userID, 720*time.Hour)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// generateToken is a helper function to generate JWT tokens with different expiration times
func (ts *TokenService) generateToken(userID string, duration time.Duration) (string, error) {
	expirationTime := time.Now().Add(duration)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(ts.jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the JWT token and returns user ID
func (ts *TokenService) ValidateToken(tokenString string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return ts.jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", jwt.NewValidationError("invalid token", jwt.ValidationErrorClaimsInvalid)
	}

	return claims.UserID, nil
}
