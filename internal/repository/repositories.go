package repository

import (
	"bonus/internal/domain"
	"database/sql"
)

type IAuthRepository interface {
	InsertCode(user *domain.Registry) error
	CheckCode(user *domain.Registry) (bool, error)
	InsertUser(user *domain.RegistryRequest) (*domain.RegistryResponse, error)
	ChecUser(email string) (bool, error)
	GetUser(email string) (*domain.LoginResponse, error)
}

type ICompanyRepository interface {
	CreateCompany(company *domain.Company) (*domain.Company, error)
}

type Repositories struct {
	AuthRepository    IAuthRepository
	CompanyRepository ICompanyRepository
}

func NewRepository(db *sql.DB) *Repositories {

	return &Repositories{
		AuthRepository:    NewAuthRepository(db),
		CompanyRepository: NewCompanyRepository(db),
	}
}
