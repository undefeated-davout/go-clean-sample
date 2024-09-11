package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConnection struct {
	DB *sql.DB
}

func NewMySQLConnection() (*MySQLConnection, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	// check if the connection is alive
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %v", err)
	}

	return &MySQLConnection{DB: db}, nil
}

func (conn *MySQLConnection) Close() error {
	if conn.DB != nil {
		return conn.DB.Close()
	}
	return nil
}
