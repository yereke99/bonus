package service

import (
	"bonus/config"
	"bonus/internal/domain"
	"bonus/internal/repository"
	"bonus/traits"
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type AuthService struct {
	ctx        context.Context
	appConfig  *config.Config
	zapLogger  *zap.Logger
	repo       *repository.Repositories
	jwtService *JWTService
}

func NewAuthService(ctx context.Context, appConfig *config.Config, zapLogger *zap.Logger, repo *repository.Repositories, jwtService *JWTService) *AuthService {
	return &AuthService{
		ctx:        ctx,
		appConfig:  appConfig,
		zapLogger:  zapLogger,
		repo:       repo,
		jwtService: jwtService,
	}
}

func (s *AuthService) SendCode(sign *domain.Registry) error {

	// here we are using method to send code to email while just code 1111
	sign.Code = 1111
	err := s.repo.AuthRepository.InsertCode(sign)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *AuthService) CheckCode(code *domain.Registry) (bool, error) {
	valid, err := s.repo.AuthRepository.CheckCode(code)
	if err != nil {
		return false, err
	}
	if !valid {
		return false, errors.New("invalid or expired code")
	}

	return true, nil
}

func (s *AuthService) Registry(model *domain.RegistryRequest) (*domain.RegistryResponse, error) {

	accessToken, err := s.jwtService.GenerateToken(model.Email, "user")
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.jwtService.RefreshToken(accessToken)
	if err != nil {
		return nil, err
	}

	// Create a hash token for the QR code
	qrToken := traits.GenerateQRToken(model.UserName, model.UserLastName, model.Email)

	model.QR = qrToken
	model.Token = refreshToken
	model.IsDeleted = false

	fmt.Println(model)
	user, err := s.repo.AuthRepository.InsertUser(model)
	if err != nil {
		return nil, err
	}

	user.AccessToken = accessToken
	return user, nil
}

func (s *AuthService) UpdateUser(userId string, model *domain.RegistryRequest) (*domain.RegistryResponse, error) {

	fmt.Println(model)
	user, err := s.repo.AuthRepository.UpdateUser(userId, model)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) CheckUser(email string) (bool, error) {
	return s.repo.AuthRepository.CheckUser(email)
}

func (s *AuthService) Login(login *domain.Registry) (*domain.LoginResponse, error) {

	valid, err := s.repo.AuthRepository.CheckCode(login)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.New("invalid code")
	}

	userExist, err := s.CheckUser(login.Email)
	if err != nil {
		return nil, err
	}

	if !userExist {
		return nil, errors.New("user does not exist")
	}

	user, err := s.repo.AuthRepository.GetUser(login.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) GetUserInfo(email string) (*domain.LoginResponse, error) {
	return s.repo.AuthRepository.GetUser(email)
}

func (s *AuthService) GetUserInfoTg(email string) (*domain.LoginResponse, error) {
	return s.repo.AuthRepository.GetUser(email)
}

func (s *AuthService) GetUserTransaction(userId string) ([]string, error) {
	return s.repo.AuthRepository.GetUserTransaction(userId)
}

func (s *AuthService) DeleteUser(uuid string) error {
	return s.repo.AuthRepository.DeleteUser(uuid)
}
