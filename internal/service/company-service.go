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

func (s *CompanyService) AddBonusUser(transaction *domain.UserTransaction) (*domain.LoginResponse, error) {
	return s.repo.CompanyRepository.AddBonusUser(transaction)
}

func (s *CompanyService) RemoveBonusUser(transaction *domain.UserTransaction) (*domain.LoginResponse, error) {
	return s.repo.CompanyRepository.RemoveBonusUser(transaction)
}

func (s *CompanyService) CreateCompany(model *domain.CompanyRequest) (*domain.Company, error) {
	return s.repo.CompanyRepository.CreateCompany(model)
}

func (s *CompanyService) CreateCompanyObject(model *domain.CompanyObject) (*domain.CompanyObject, error) {
	return s.repo.CompanyRepository.CreateCompanyObject(model)
}

func (s *CompanyService) GetCompanies() ([]*domain.Company, error) {
	return s.repo.CompanyRepository.GetCompanies()
}

func (s *CompanyService) GetCompanyObjects(companyId string) ([]*domain.CompanyObject, error) {
	return s.repo.CompanyRepository.GetCompanyObjects(companyId)
}

func (s *CompanyService) GetCompanyObjectInfo(companyId string) (*domain.CompanyObject, error) {
	return s.repo.CompanyRepository.GetCompanyObjectInfo(companyId)
}

func (s *CompanyService) GetCompanyObjectTransAction(companyId string) (*domain.CompanyObjectTransAction, error) {
	return s.repo.CompanyRepository.GetCompanyObjectTransAction(companyId)
}
