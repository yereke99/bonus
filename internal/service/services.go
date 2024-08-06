package service

import (
	"bonus/config"
	"context"

	"go.uber.org/zap"
)

type IAuthServices interface {
	Registry() error
	Login() error
}

type Services struct {
	AuthService IAuthServices
}

func NewServices(ctx context.Context, appConfig *config.Config, zapLogger *zap.Logger) *Services {

	return &Services{
		AuthService: NewAuthService(ctx, appConfig, zapLogger),
	}
}
