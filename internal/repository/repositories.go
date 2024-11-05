package repository

import (
	"bonus/internal/domain"
	"database/sql"
)

type IAuthRepository interface {
	InsertCode(user *domain.Registry) error
	CheckCode(user *domain.Registry) (bool, error)
	InsertUser(user *domain.RegistryRequest) (*domain.RegistryResponse, error)
	UpdateUser(userID string, user *domain.RegistryRequest) (*domain.RegistryResponse, error)
	CheckUser(email string) (bool, error)
	GetUser(email string) (*domain.LoginResponse, error)
	GetUserTransaction(userId string) ([]string, error)
	DeleteUser(uuid string) error
}

type ICompanyRepository interface {
	CreateCompany(company *domain.CompanyRequest) (*domain.Company, error)
	CreateCompanyObject(object *domain.CompanyObject) (*domain.CompanyObject, error)
	GetCompanies() ([]*domain.Company, error)
	AddBonusUser(transaction *domain.UserTransaction) (*domain.LoginResponse, error)
	RemoveBonusUser(transaction *domain.UserTransaction) (*domain.LoginResponse, error)
	GetCompanyObjectInfo(companyId string) (*domain.CompanyObject, error)
	GetCompanyObjects(uuid string) ([]*domain.CompanyObject, error)
	GetCompanyObjectTransAction(companyId string) (*domain.CompanyObjectTransAction, error)
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
