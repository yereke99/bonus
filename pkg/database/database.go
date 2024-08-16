package database

import (
	"bonus/config"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func ConnectToDatabase(databaseConfig *config.DatabaseConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Database,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	pwd, _ := os.Getwd()

	log.Println("Current working directory:", pwd)

	// Создаем объект миграции
	m, err := migrate.New("file://./migration", connectionString)
	if err != nil {
		return nil, err
	}
	defer m.Close() // Закрываем объект миграции после использования

	// Применяем миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	log.Println("Миграция успешно выполнена")
	log.Println("Connected to PostgreSQL database")

	return db, nil
}
