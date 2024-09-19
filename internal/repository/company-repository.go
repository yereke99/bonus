package repository

import (
	"bonus/internal/domain"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
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
	query := `
		INSERT INTO company (company, company_name, email, city, company_address, company_iin, bonus) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, company, company_name, email, city, company_address, company_iin, bonus, isDeleted
	`

	insertedCompany := &domain.Company{}

	var isDeleted sql.NullBool
	err := r.db.QueryRow(query,
		company.Company,
		company.CompanyName,
		company.Email,
		company.City,
		company.CompanyAddress,
		company.CompanyIIN,
		company.Bonus,
	).Scan(
		&insertedCompany.ID, // ID is now a string
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

	// Set isDeleted value in insertedCompany
	insertedCompany.IsDeleted = isDeleted.Valid && isDeleted.Bool

	return insertedCompany, nil
}

func (r *CompanyRepository) CreateCompanyObject(object *domain.CompanyObject) (*domain.CompanyObject, error) {
	// Проверяем валидность данных
	if object == nil {
		return nil, errors.New("company object is nil")
	}

	query := `
	INSERT INTO business_types 
	(id, company_id, business_type, city, email, working_time, trc, business_address, floor, business_line, business_number) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING id`

	// Генерируем новый UUID для поля id
	newID := uuid.New()

	// Выполняем запрос с передачей параметров
	err := r.db.QueryRow(
		query,
		newID,
		object.CompanyID,
		object.TypeBusines,
		object.City,
		object.Email,
		object.BusinessTime,
		object.Trc,
		object.BusinessAddress,
		object.Floor,
		object.Column,
		object.NumberColumn,
	).Scan(&object.ID)

	if err != nil {
		log.Println("Failed to insert company object:", err)
		return nil, fmt.Errorf("error creating company object: %w", err)
	}

	object.ID = newID.String()

	return object, nil
}

func (r *CompanyRepository) GetCompanies() ([]*domain.Company, error) {
	query := `
		SELECT id, company, company_name, email, city, company_address, company_iin, bonus, isDeleted
		FROM company
	`

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
			&company.ID, // ID is now a string
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

		// Set isDeleted value in company
		company.IsDeleted = isDeleted.Valid && isDeleted.Bool

		companies = append(companies, company)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
