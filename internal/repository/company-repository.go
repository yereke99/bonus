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
			  VALUES ($1, $2, $3, $4, $5, $6, $7)
			  RETURNING id, company, company_name, email, city, company_addres, company_iin, bonus, isDeleted`

	insertedCompany := &domain.Company{}

	var isDeleted sql.NullBool
	err := r.db.QueryRow(query, company.Company, company.CompanyName, company.Email, company.City, company.CompanyAddress, company.CompanyIIN, company.Bonus).Scan(
		&insertedCompany.ID,
		&insertedCompany.Company,
		&insertedCompany.CompanyName,
		&insertedCompany.Email,
		&insertedCompany.City,
		&insertedCompany.CompanyAddress,
		&insertedCompany.CompanyIIN,
		&insertedCompany.Bonus,
		&isDeleted,
	)
	if err != nil {
		return nil, err
	}

	// Устанавливаем значение isDeleted в insertedCompany
	insertedCompany.IsDeleted = isDeleted.Valid && isDeleted.Bool

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
		var isDeleted sql.NullBool

		err := rows.Scan(
			&company.ID,
			&company.Company,
			&company.CompanyName,
			&company.Email,
			&company.City,
			&company.CompanyAddress,
			&company.CompanyIIN,
			&company.Bonus,
			&isDeleted,
		)
		if err != nil {
			return nil, err
		}

		// Устанавливаем значение isDeleted в компании
		company.IsDeleted = isDeleted.Valid && isDeleted.Bool

		companies = append(companies, company)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
