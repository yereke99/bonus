package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock JWT secret and issuer
const mockSecretKey = "mockSecret"
const mockIssuer = "mockIssuer"

func TestJWTService_GetUserId_ValidToken(t *testing.T) {
	// Arrange
	jwtService := NewJWTService(mockSecretKey, mockIssuer)
	email := "test@example.com"
	role := "user"

	// Generate a valid token
	tokenString, err := jwtService.GenerateToken(email, role)
	assert.NoError(t, err)

	// Act
	email_, err := jwtService.GetUserId(tokenString)
	fmt.Println(email)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, email, email_)
}

func TestJWTService_GetCompanyId_ValidToken(t *testing.T) {
	jwtService := NewJWTService(mockSecretKey, mockIssuer)

	email := "test@example.com"
	role := "company"

	// Generate a valid token
	tokenString, err := jwtService.GenerateToken(email, role)
	assert.NoError(t, err)

	email_, err := jwtService.GetCompanyId(tokenString)
	assert.NoError(t, err)

	fmt.Println(email)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, email, email_)
}

func TestJWTService_GetCompanyObjectId_ValidToken(t *testing.T) {
	jwtService := NewJWTService(mockSecretKey, mockIssuer)

	email := "test@example.com"
	role := "company"

	// Generate a valid token
	tokenString, err := jwtService.GenerateToken(email, role)
	assert.NoError(t, err)

	email_, err := jwtService.GetCompanyId(tokenString)
	assert.NoError(t, err)

	fmt.Println(email)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, email, email_)
}
