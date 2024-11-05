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
	GetUserInfo(email string) (*domain.LoginResponse, error)
	GetUserInfoTg(email string) (*domain.LoginResponse, error)
	GetUserTransaction(userId string) ([]string, error)
	DeleteUser(uuid string) error
}

type IJWTServices interface {
	GenerateToken(email string, role string) (string, error)
	RefreshToken(tokenString string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	GetUserId(tokenString string) (string, error)
	GetCompanyId(tokenString string) (string, error)
	GetCompanyObjectId(tokenString string) (string, error)
}

type ICompanyService interface {
	CreateCompany(model *domain.CompanyRequest) (*domain.Company, error)
	CreateCompanyObject(model *domain.CompanyObject) (*domain.CompanyObject, error)
	GetCompanies() ([]*domain.Company, error)
	AddBonusUser(transaction *domain.UserTransaction) (*domain.LoginResponse, error)
	GetCompanyObjectInfo(companyId string) (*domain.CompanyObject, error)
	GetCompanyObjects(companyId string) ([]*domain.CompanyObject, error)
	GetCompanyObjectTransAction(companyId string) (*domain.CompanyObjectTransAction, error)
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
