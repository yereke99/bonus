package repository

import (
	"bonus/internal/domain"
	"database/sql"
)

type IAuthRepository interface {
	InsertCode(user *domain.Registry) error
	CheckCode(user *domain.Registry) (bool, error)
	InsertUser(user *domain.RegistryRequest) (int64, error)
	ChecUser(email string) (bool, error)
	GetUser(email string) (*domain.LoginResponse, error)
}

type Repositories struct {
	AuthRepository IAuthRepository
}

func NewRepository(db *sql.DB) *Repositories {

	return &Repositories{
		AuthRepository: NewAuthRepository(db),
	}
}
