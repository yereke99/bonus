package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	secretKey string
	issuer    string
}

func NewJWTService(secretKey, issuer string) *JWTService {
	return &JWTService{
		secretKey: secretKey,
		issuer:    issuer,
	}
}

// GenerateToken generates a new JWT token
func (s *JWTService) GenerateToken(email string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires in 1 hour
	claims["iss"] = s.issuer

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// RefreshToken generates a new token by refreshing the existing token
func (s *JWTService) RefreshToken(tokenString string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return "", err
	}

	// Remove existing expiration and set a new one
	delete(claims, "exp")
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Refresh token for another hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newTokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return newTokenString, nil
}

// ValidateToken validates a token and returns the parsed token
func (s *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})

	return token, err
}
