package service

import (
	"bonus/config"
	"bonus/internal/domain"
	"context"

	"go.uber.org/zap"
)

type AuthService struct {
	ctx       context.Context
	appConfig *config.Config
	zapLogger *zap.Logger
}

func NewAuthService(ctx context.Context, appConfig *config.Config, zapLogger *zap.Logger) *AuthService {
	return &AuthService{
		ctx:       ctx,
		appConfig: appConfig,
		zapLogger: zapLogger,
	}
}

func (s *AuthService) SendCode(code *domain.CodeRequest) error {

	return nil
}

func (s *AuthService) Registry(model *domain.RegistryRequest) (*domain.RegistryResponse, error) {

	return nil, nil
}

func (s *AuthService) Login() error {

	return nil
}
