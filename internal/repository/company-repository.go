package repository

import (
	"bonus/internal/domain"
	"database/sql"
)

type CompanyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (r *CompanyRepository) CreateCompany(company *domain.Company) (*domain.Company, error) {
	query := `INSERT INTO company (id, company, company_name, email, city, company_addres, company_iin, bonus, isDeleted) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
			  RETURNING id, company, company_name, email, city, company_addres, company_iin, bonus, isDeleted`

	// Create a variable to hold the inserted company details
	insertedCompany := &domain.Company{}

	// Execute the query and scan the returned row into the insertedCompany struct
	err := r.db.QueryRow(query, company.ID, company.Company, company.CompanyName, company.Email, company.City, company.CompanyAddress, company.CompanyIIN, company.Bonus, company.IsDeleted).Scan(
		&insertedCompany.ID,
		&insertedCompany.Company,
		&insertedCompany.CompanyName,
		&insertedCompany.Email,
		&insertedCompany.City,
		&insertedCompany.CompanyAddress,
		&insertedCompany.CompanyIIN,
		&insertedCompany.Bonus,
		&insertedCompany.IsDeleted,
	)
	if err != nil {
		return nil, err
	}

	return insertedCompany, nil
}
