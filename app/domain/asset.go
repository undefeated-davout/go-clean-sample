package domain

type Asset struct {
	Ticker string
	Weight float64
}

type AssetRepository interface {
	Add(asset Asset) error
	Remove(ticker string) error
}
