package database

import (
	"bonus/config"
	"bonus/internal/domain"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
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

	log.Println("Connected to PostgreSQL database")

	return db, nil
}

func Migrate(db *sql.DB, zapLogger *zap.Logger) error {
	// First, ensure the `uuid-ossp` extension is enabled.
	extensionQuery := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
	_, err := db.Exec(extensionQuery)
	if err != nil {
		zapLogger.Error("failed to enable uuid-ossp extension", zap.Error(err))
		return fmt.Errorf("failed to enable uuid-ossp extension: %w", err)
	}

	// List of table creation SQL statements
	sqls := []string{
		customerTable,
		codeCacheTable,
		companyTable,
		businessTypesTable,
	}

	// Check if one of the tables already exists (e.g., `code_cache`)
	query := `SELECT 1 FROM code_cache LIMIT 1`
	_, err = db.Query(query)
	if err == nil {
		return domain.ErrExistsTable
	}

	// Execute each SQL statement
	for _, sqlQuery := range sqls {
		_, err := db.Exec(sqlQuery)
		if err != nil {
			zapLogger.Error("failed to execute migration", zap.Error(err), zap.String("query", sqlQuery))
			return fmt.Errorf("failed to execute migration: %w", err)
		}
		zapLogger.Info("migration executed", zap.String("query", sqlQuery))
	}

	zapLogger.Info("All migrations executed successfully")
	return nil
}
