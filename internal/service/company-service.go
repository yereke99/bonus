package service

import (
	"bonus/config"
	"bonus/internal/domain"
	"bonus/internal/repository"
	"context"

	"go.uber.org/zap"
)

type CompanyService struct {
	ctx        context.Context
	appConfig  *config.Config
	zapLogger  *zap.Logger
	repo       *repository.Repositories
	jwtService *JWTService
}

func NewCompanyService(ctx context.Context, appConfig *config.Config, zapLogger *zap.Logger, repo *repository.Repositories, jwtService *JWTService) *CompanyService {
	return &CompanyService{
		ctx:        ctx,
		appConfig:  appConfig,
		zapLogger:  zapLogger,
		repo:       repo,
		jwtService: jwtService,
	}
}

func (s *CompanyService) CreateCompany(model *domain.CompanyRequest) (*domain.Company, error) {
	return s.repo.CompanyRepository.CreateCompany(model)
}

func (s *CompanyService) GetCompanies() ([]*domain.Company, error) {
	return s.repo.CompanyRepository.GetCompanies()
}
