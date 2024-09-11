package domain

// Assetエンティティ
type Asset struct {
	Ticker string
	Weight float64
}

// AssetRepositoryインターフェース
type AssetRepository interface {
	Add(asset Asset) error
	Remove(ticker string) error
}
