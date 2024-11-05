package repository

import (
	"bonus/internal/domain"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type CompanyRepository struct {
	db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

// Функция для создания компании
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
		&insertedCompany.ID, // ID теперь генерируется в базе
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

	// Устанавливаем значение isDeleted в объекте компании
	insertedCompany.IsDeleted = isDeleted.Valid && isDeleted.Bool

	return insertedCompany, nil
}

// Функция для создания объекта компании
func (r *CompanyRepository) CreateCompanyObject(object *domain.CompanyObject) (*domain.CompanyObject, error) {
	// Проверяем валидность данных
	if object == nil {
		return nil, errors.New("company object is nil")
	}

	query := `
	INSERT INTO business_types 
	(company_id, business_type, city, email, working_time, trc, business_address, floor, business_line, business_number) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id`

	// Выполняем запрос без генерации UUID в коде
	err := r.db.QueryRow(
		query,
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

	return object, nil
}

// Функция для получения списка компаний
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
			&company.ID, // ID теперь генерируется в базе
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

		// Устанавливаем значение isDeleted в объекте компании
		company.IsDeleted = isDeleted.Valid && isDeleted.Bool

		companies = append(companies, company)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

func (r *CompanyRepository) GetCompanyObjectTransAction(companyId string) (*domain.CompanyObjectTransAction, error) {

	return &domain.CompanyObjectTransAction{}, nil
}

func (r *CompanyRepository) GetCompanyObjectInfo(companyId string) (*domain.CompanyObject, error) {
	query := `
        SELECT
            id,
            company_id,
            business_type AS typeBusiness,
            business_line AS businessName,
            city,
            email,
            working_time AS businessTime,
            trc,
            business_address AS businessAddress,
            floor,
            business_line AS column,
            business_number AS numberColumn,
            false AS isDeleted
        FROM business_types
        WHERE company_id = $1
        LIMIT 1
    `

	row := r.db.QueryRow(query, companyId)

	var obj domain.CompanyObject
	err := row.Scan(
		&obj.ID,
		&obj.CompanyID,
		&obj.TypeBusines,
		&obj.BusinesName,
		&obj.City,
		&obj.Email,
		&obj.BusinessTime,
		&obj.Trc,
		&obj.BusinessAddress,
		&obj.Floor,
		&obj.Column,
		&obj.NumberColumn,
		&obj.IsDeleted,
	)
	if err == sql.ErrNoRows {
		return nil, nil // Вернуть nil, если компания с указанным ID не найдена
	} else if err != nil {
		return nil, err // Возвращаем ошибку, если возникла другая ошибка
	}

	return &obj, nil
}

func (r *CompanyRepository) AddBonusUser(transaction *domain.UserTransaction) (*domain.LoginResponse, error) {
	// Обновляем бонус пользователя
	updateQuery := `
        UPDATE customer
        SET bonus = bonus + $1
        WHERE id = $2
        RETURNING id, user_name, user_last_name, email, locations, city, qr, bonus, token, isDeleted
    `

	var response domain.LoginResponse
	err := r.db.QueryRow(
		updateQuery,
		transaction.Bonus,
		transaction.UserId,
	).Scan(
		&response.ID,
		&response.UserName,
		&response.UserLastName,
		&response.Email,
		&response.Locations,
		&response.City,
		&response.QR,
		&response.Bonus,
		&response.Token,
		&response.IsDeleted,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *CompanyRepository) RemoveBonusUser(transaction *domain.UserTransaction) (*domain.LoginResponse, error) {
	// Проверяем текущий бонус пользователя
	var currentBonus int
	checkQuery := `SELECT bonus FROM customer WHERE id = $1`
	err := r.db.QueryRow(checkQuery, transaction.UserId).Scan(&currentBonus)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Проверка на то, что бонус не станет отрицательным
	if currentBonus < transaction.Bonus {
		return nil, errors.New("insufficient bonus")
	}

	// Обновляем бонус пользователя
	updateQuery := `
        UPDATE customer
        SET bonus = bonus - $1
        WHERE id = $2
        RETURNING id, user_name, user_last_name, email, locations, city, qr, bonus, token, isDeleted
    `

	var response domain.LoginResponse
	err = r.db.QueryRow(
		updateQuery,
		transaction.Bonus,
		transaction.UserId,
	).Scan(
		&response.ID,
		&response.UserName,
		&response.UserLastName,
		&response.Email,
		&response.Locations,
		&response.City,
		&response.QR,
		&response.Bonus,
		&response.Token,
		&response.IsDeleted,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *CompanyRepository) GetCompanyObjects(uuid string) ([]*domain.CompanyObject, error) {
	q := `SELECT id, company_id, typeBusiness, businessName, city, email, businessTime, trc, businessAddress, floor, column, numberColumn, isDeleted FROM business_types WHERE id=$1`

	// Создаем слайс для хранения объектов CompanyObject
	var companyObjects []*domain.CompanyObject

	// Выполняем запрос к базе данных
	rows, err := r.db.Query(q, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Итерируемся по результатам запроса и маппим их на CompanyObject
	for rows.Next() {
		var obj domain.CompanyObject
		err := rows.Scan(
			&obj.ID,
			&obj.CompanyID,
			&obj.TypeBusines,
			&obj.BusinesName,
			&obj.City,
			&obj.Email,
			&obj.BusinessTime,
			&obj.Trc,
			&obj.BusinessAddress,
			&obj.Floor,
			&obj.Column,
			&obj.NumberColumn,
			&obj.IsDeleted,
		)
		if err != nil {
			return nil, err
		}
		companyObjects = append(companyObjects, &obj)
	}

	// Проверяем наличие ошибок после итерации
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companyObjects, nil
}
