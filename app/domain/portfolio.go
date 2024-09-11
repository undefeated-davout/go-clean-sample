package domain

// Portfolioエンティティ
type Portfolio struct {
	ID     int
	Name   string
	Assets []Asset
}

// PortfolioRepositoryインターフェース
type PortfolioRepository interface {
	Save(portfolio Portfolio) error
	GetByID(id int) (*Portfolio, error)
}
