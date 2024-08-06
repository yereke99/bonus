package service

import (
	"bonus/config"
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

func (s *AuthService) Registry() error {

	return nil
}

func (s *AuthService) Login() error {

	return nil
}
