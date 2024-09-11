package repository

import (
	"database/sql"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type MySQLAssetRepository struct {
	conn *sql.DB
}

func NewMySQLAssetRepository(conn *sql.DB) domain.AssetRepository {
	return &MySQLAssetRepository{conn: conn}
}

func (r *MySQLAssetRepository) Add(asset domain.Asset) error {
	_, err := r.conn.Exec("INSERT INTO assets (ticker, weight) VALUES (?, ?)", asset.Ticker, asset.Weight)
	return err
}

func (r *MySQLAssetRepository) Remove(ticker string) error {
	_, err := r.conn.Exec("DELETE FROM assets WHERE ticker = ?", ticker)
	return err
}
