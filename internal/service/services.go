package service

import (
	"bonus/config"
	"bonus/internal/domain"
	"context"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type IAuthServices interface {
	SendCode(code *domain.CodeRequest) error
	Registry(model *domain.RegistryRequest) (*domain.RegistryResponse, error)
	Login() error
}

type IJWTServices interface {
	GenerateToken(userID string, role string) (string, error)
	RefreshToken(tokenString string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type Services struct {
	AuthService IAuthServices
	JWTService  IJWTServices
}

func NewServices(ctx context.Context, appConfig *config.Config, zapLogger *zap.Logger) *Services {

	return &Services{
		AuthService: NewAuthService(ctx, appConfig, zapLogger),
	}
}
