package service

import (
	"bonus/config"
	"bonus/internal/domain"
	"bonus/internal/repository"
	"context"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type IAuthServices interface {
	SendCode(sign *domain.Registry) error
	Registry(model *domain.RegistryRequest) (*domain.RegistryResponse, error)
	UpdateUser(userId string, model *domain.RegistryRequest) (*domain.RegistryResponse, error)
	Login(login *domain.Registry) (*domain.LoginResponse, error)
}

type IJWTServices interface {
	GenerateToken(email string, role string) (string, error)
	RefreshToken(tokenString string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type ICompanyService interface {
	CreateCompany(model *domain.CompanyRequest) (*domain.Company, error)
	GetCompanies() ([]*domain.Company, error)
}

type Services struct {
	AuthService    IAuthServices
	JWTService     IJWTServices
	CompanyService ICompanyService
}

func NewServices(ctx context.Context, appConfig *config.Config, zapLogger *zap.Logger, repo *repository.Repositories) *Services {
	jwtServices := NewJWTService(appConfig.SecretKey, appConfig.Issuer)
	return &Services{
		AuthService:    NewAuthService(ctx, appConfig, zapLogger, repo, jwtServices),
		CompanyService: NewCompanyService(ctx, appConfig, zapLogger, repo, jwtServices),
		JWTService:     jwtServices,
	}
}
