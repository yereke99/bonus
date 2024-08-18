package service

import (
	"bonus/config"
	"bonus/internal/domain"
	"bonus/internal/repository"
	"context"
	"errors"

	"go.uber.org/zap"
)

type AuthService struct {
	ctx       context.Context
	appConfig *config.Config
	zapLogger *zap.Logger
	repo      *repository.Repositories
}

func NewAuthService(ctx context.Context, appConfig *config.Config, zapLogger *zap.Logger, repo *repository.Repositories) *AuthService {
	return &AuthService{
		ctx:       ctx,
		appConfig: appConfig,
		zapLogger: zapLogger,
		repo:      repo,
	}
}

func (s *AuthService) SendCode(code *domain.CodeRequest) error {

	// here we are using method to send code to email while just code 1111

	return nil
}

func (s *AuthService) Registry(model *domain.RegistryRequest) (*domain.RegistryResponse, error) {

	return nil, nil
}

func (s *AuthService) CheckUser(email string) (bool, error) {
	return s.repo.AuthRepository.ChecUser(email)
}

func (s *AuthService) Login(login *domain.Registry) (*domain.LoginResponse, error) {

	ok, err := s.repo.AuthRepository.CheckCode(login)
	if err != nil {
		return nil, err
	}
	if ok {
		ok, err := s.CheckUser(login.Email)
		if err != nil {
			return nil, err
		}
		if ok {
			user, err := s.repo.AuthRepository.GetUser(login.Email)
			if err != nil {
				return nil, err
			}
			return user, nil
		}

		return nil, errors.New("does not exists user")
	}

	return nil, errors.New("does not exists code")
}
