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

// NewMySQLConnection は MySQL データベースへの接続を確立します
func NewMySQLConnection() (*MySQLConnection, error) {
	// 環境変数から MySQL の接続情報を取得
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	// DSN (Data Source Name) の作成
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)

	// MySQL データベースへ接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	// 接続が成功しているか確認
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %v", err)
	}

	return &MySQLConnection{DB: db}, nil
}

// Close はデータベース接続をクローズします
func (conn *MySQLConnection) Close() error {
	if conn.DB != nil {
		return conn.DB.Close()
	}
	return nil
}
