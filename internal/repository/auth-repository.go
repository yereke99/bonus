package repository

import (
	"bonus/internal/domain"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {

	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) InsertCode(user *domain.Registry) error {

	currentTime := time.Now()

	// Check if the email already exists
	var count int
	checkQuery := `SELECT COUNT(*) FROM code_cache WHERE email = $1`
	err := r.db.QueryRow(checkQuery, user.Email).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		// If email exists, update the code and created_at
		updateQuery := `UPDATE code_cache SET code = $1, created_at = $2 WHERE email = $3`
		_, err := r.db.Exec(updateQuery, user.Code, currentTime, user.Email)
		if err != nil {
			return err
		}
	} else {
		// If email does not exist, insert a new record
		insertQuery := `INSERT INTO code_cache(email, code, created_at) VALUES($1, $2, $3)`
		_, err := r.db.Exec(insertQuery, user.Email, user.Code, currentTime)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *AuthRepository) CheckCode(user *domain.Registry) (bool, error) {
	var createdAt time.Time
	var code int

	q := `SELECT created_at, code FROM code_cache WHERE email = $1`
	err := r.db.QueryRow(q, user.Email).Scan(&createdAt, &code)
	if err != nil {
		fmt.Println("here")
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("here")
			return false, errors.New("code not found for the provided email")
		}
		fmt.Println(err)
		return false, err
	}

	// Check if the code has expired
	if time.Since(createdAt) > 10*time.Minute {
		return false, errors.New("code has expired")
	}

	// Validate the code
	if code != user.Code {
		fmt.Println("here")
		return false, errors.New("invalid code")
	}

	return true, nil
}

func (r *AuthRepository) InsertUser(user *domain.RegistryRequest) (*domain.RegistryResponse, error) {
	q := `
        INSERT INTO customer(user_name, user_last_name, email, locations, city, qr, bonus, token, isDeleted) 
        VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, user_name, user_last_name, email, locations, city, qr, bonus, token, isDeleted
    `

	var resp domain.RegistryResponse
	err := r.db.QueryRow(q,
		user.UserName,
		user.UserLastName,
		user.Email,
		user.Locations,
		user.City,
		user.QR,
		user.Bonus,
		user.Token,
		user.IsDeleted,
	).Scan(
		&resp.ID,
		&resp.UserName,
		&resp.UserLastName,
		&resp.Email,
		&resp.Locations,
		&resp.City,
		&resp.QR,
		&resp.Bonus,
		&resp.Token,
		&resp.IsDeleted,
	)
	if err != nil {
		return nil, fmt.Errorf("could not insert user: %v", err)
	}

	return &resp, nil
}

func (r *AuthRepository) UpdateUser(userID int64, user *domain.RegistryRequest) (*domain.RegistryResponse, error) {
	// Начальный шаблон для запроса
	query := "UPDATE customer SET"
	args := []interface{}{}
	argIndex := 1

	// Проверяем и добавляем в запрос соответствующие поля, если они не пустые
	if user.UserName != "" {
		query += fmt.Sprintf(" user_name = $%d,", argIndex)
		args = append(args, user.UserName)
		argIndex++
	}
	if user.UserLastName != "" {
		query += fmt.Sprintf(" user_last_name = $%d,", argIndex)
		args = append(args, user.UserLastName)
		argIndex++
	}
	if user.Locations != "" {
		query += fmt.Sprintf(" locations = $%d,", argIndex)
		args = append(args, user.Locations)
		argIndex++
	}
	if user.City != "" {
		query += fmt.Sprintf(" city = $%d,", argIndex)
		args = append(args, user.City)
		argIndex++
	}

	// Удаляем последнюю запятую и добавляем условие WHERE
	query = query[:len(query)-1]
	query += fmt.Sprintf(" WHERE id = $%d RETURNING id, user_name, user_last_name, email, locations, city, qr, bonus, token, isDeleted", argIndex)
	args = append(args, userID)

	// Выполнение запроса
	var resp domain.RegistryResponse
	err := r.db.QueryRow(query, args...).Scan(
		&resp.ID,
		&resp.UserName,
		&resp.UserLastName,
		&resp.Email,
		&resp.Locations,
		&resp.City,
		&resp.QR,
		&resp.Bonus,
		&resp.Token,
		&resp.IsDeleted,
	)
	if err != nil {
		return nil, fmt.Errorf("could not update user: %v", err)
	}

	return &resp, nil
}

func (r *AuthRepository) ChecUser(email string) (bool, error) {
	q := `SELECT email FROM customer WHERE email = $1`

	var retrievedEmail string
	err := r.db.QueryRow(q, email).Scan(&retrievedEmail)
	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, nil
	}

	// if user exists
	return true, nil
}

func (r *AuthRepository) GetUser(email string) (*domain.LoginResponse, error) {
	q := `
        SELECT id, user_name, user_last_name, email, locations, city, qr, bonus, token, isDeleted 
        FROM customer 
        WHERE email = $1 AND isDeleted = false
    `

	var user domain.LoginResponse

	err := r.db.QueryRow(q, email).Scan(
		&user.ID,
		&user.UserName,
		&user.UserLastName,
		&user.Email,
		&user.Locations,
		&user.City,
		&user.QR,
		&user.Bonus,
		&user.Token,
		&user.IsDeleted,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user with email %s not found", email)
	}

	if err != nil {
		return nil, fmt.Errorf("error retrieving user: %v", err)
	}

	return &user, nil
}
