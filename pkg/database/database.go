package database

import (
	"bonus/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var sqls = []string{}

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

	log.Println("Connected to PostgreSQL database")

	return db, nil
}
