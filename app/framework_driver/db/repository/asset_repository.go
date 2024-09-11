package repository

import (
	"database/sql"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type MySQLAssetRepository struct {
	conn *sql.DB
}

// NewMySQLAssetRepository は新しいAssetリポジトリを作成する
func NewMySQLAssetRepository(conn *sql.DB) domain.AssetRepository {
	return &MySQLAssetRepository{conn: conn}
}

// Add はAssetエンティティを追加する
func (r *MySQLAssetRepository) Add(asset domain.Asset) error {
	_, err := r.conn.Exec("INSERT INTO assets (ticker, weight) VALUES (?, ?)", asset.Ticker, asset.Weight)
	return err
}

// Remove はAssetエンティティを削除する
func (r *MySQLAssetRepository) Remove(ticker string) error {
	_, err := r.conn.Exec("DELETE FROM assets WHERE ticker = ?", ticker)
	return err
}
