package repository

import (
	"bonus/internal/domain"
	"database/sql"
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

	q := `INSERT INTO code_cache(email, code, created_at) VALUES($1, $2, $3)`
	_, err := r.db.Exec(q, user.Email, user.Code)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) CheckCode(user *domain.Registry) (bool, error) {
	var createdAt time.Time

	q := `SELECT created_at FROM code_cache WHERE email = $1 AND code = $2`
	err := r.db.QueryRow(q, user.Email, user.Code).Scan(&createdAt)
	if err != nil {
		return false, err
	}

	if time.Since(createdAt) > 10*time.Minute {
		return false, nil
	}

	return true, nil
}

func (r *AuthRepository) InsertUser(user *domain.RegistryRequest) (int64, error) {
	q := `
        INSERT INTO customer(user_name, user_last_name, email, locations, city, qr, bonus, token, isDeleted) 
        VALUES($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id
    `

	var userID int64
	err := r.db.QueryRow(q,
		user.UserName,
		user.UserLastName,
		user.Email,
		user.Locations,
		user.City,
		user.QR,
		user.Bonus,
		user.IsDeleted,
	).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("could not insert user: %v", err)
	}

	return userID, nil
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
