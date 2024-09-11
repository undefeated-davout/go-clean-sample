package domain

type Portfolio struct {
	ID     int
	Name   string
	Assets []Asset
}

type PortfolioRepository interface {
	Save(portfolio Portfolio) error
	GetByID(id int) (*Portfolio, error)
}
