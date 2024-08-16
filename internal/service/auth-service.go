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

	// here we are using method to send code to email while just code 1111

	return nil
}

func (s *AuthService) Registry(model *domain.RegistryRequest) (*domain.RegistryResponse, error) {

	return nil, nil
}

func (s *AuthService) Login() error {

	return nil
}
