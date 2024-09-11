package repository

import (
	"database/sql"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type MySQLPortfolioRepository struct {
	conn *sql.DB
}

// NewMySQLPortfolioRepository は新しいPortfolioリポジトリを作成する
func NewMySQLPortfolioRepository(conn *sql.DB) domain.PortfolioRepository {
	return &MySQLPortfolioRepository{conn: conn}
}

// Save はPortfolioエンティティを保存する
func (r *MySQLPortfolioRepository) Save(portfolio domain.Portfolio) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}

	// ポートフォリオ自体を保存
	res, err := tx.Exec("INSERT INTO portfolios (name) VALUES (?)", portfolio.Name)
	if err != nil {
		tx.Rollback()
		return err
	}

	portfolioID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}

	// 各資産を保存
	for _, asset := range portfolio.Assets {
		_, err := tx.Exec("INSERT INTO portfolio_assets (portfolio_id, ticker, weight) VALUES (?, ?, ?)", portfolioID, asset.Ticker, asset.Weight)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// GetByID はIDでPortfolioエンティティを取得する
func (r *MySQLPortfolioRepository) GetByID(id int) (*domain.Portfolio, error) {
	portfolio := &domain.Portfolio{}

	// ポートフォリオ本体を取得
	err := r.conn.QueryRow("SELECT id, name FROM portfolios WHERE id = ?", id).Scan(&portfolio.ID, &portfolio.Name)
	if err != nil {
		return nil, err
	}

	// 資産情報を取得
	rows, err := r.conn.Query("SELECT ticker, weight FROM portfolio_assets WHERE portfolio_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []domain.Asset
	for rows.Next() {
		var asset domain.Asset
		if err := rows.Scan(&asset.Ticker, &asset.Weight); err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	portfolio.Assets = assets
	return portfolio, nil
}

// Delete はPortfolioエンティティを削除する
func (r *MySQLPortfolioRepository) Delete(id int) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}

	// ポートフォリオに紐づく資産を削除
	_, err = tx.Exec("DELETE FROM portfolio_assets WHERE portfolio_id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// ポートフォリオ自体を削除
	_, err = tx.Exec("DELETE FROM portfolios WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
