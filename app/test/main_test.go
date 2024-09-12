package test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func SetupTestDB() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", user, password, host, port, database)
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlFile, err := os.ReadFile("../sql/schema.sql")
	if err != nil {
		return nil, fmt.Errorf("failed to read SQL file: %w", err)
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return nil, fmt.Errorf("failed to execute migration: %w", err)
	}

	log.Println("Database migrated successfully.")
	return db, nil
}

func TestMain(m *testing.M) {
	db, err := SetupTestDB()
	if err != nil {
		log.Fatalf("Failed to set up test database: %v", err)
	}
	defer db.Close()

	code := m.Run()

	os.Exit(code)
}
