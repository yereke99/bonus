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

func (r *CompanyRepository) CreateCompany(company *domain.CompanyRequest) (*domain.Company, error) {
	query := `INSERT INTO company (company, company_name, email, city, company_addres, company_iin, bonus) 
			  VALUES (?, ?, ?, ?, ?, ?, ?)
			  RETURNING id, company, company_name, email, city, company_addres, company_iin, bonus, isDeleted`

	// Create a variable to hold the inserted company details
	insertedCompany := &domain.Company{}

	// Execute the query and scan the returned row into the insertedCompany struct
	err := r.db.QueryRow(query, company.Company, company.CompanyName, company.Email, company.City, company.CompanyAddress, company.CompanyIIN, company.Bonus).Scan(
		&insertedCompany.Company,
		&insertedCompany.CompanyName,
		&insertedCompany.Email,
		&insertedCompany.City,
		&insertedCompany.CompanyAddress,
		&insertedCompany.CompanyIIN,
		&insertedCompany.Bonus,
	)
	if err != nil {
		return nil, err
	}

	return insertedCompany, nil
}

func (r *CompanyRepository) GetCompanies() ([]*domain.Company, error) {
	query := `SELECT id, company, company_name, email, city, company_addres, company_iin, bonus, isDeleted FROM company`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []*domain.Company

	for rows.Next() {
		company := &domain.Company{}
		err := rows.Scan(
			&company.ID,
			&company.Company,
			&company.CompanyName,
			&company.Email,
			&company.City,
			&company.CompanyAddress,
			&company.CompanyIIN,
			&company.Bonus,
			&company.IsDeleted,
		)
		if err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
