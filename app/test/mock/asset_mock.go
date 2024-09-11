package mock

import (
	"errors"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type MockAssetRepository struct {
	assets map[string]*domain.Asset
}

func NewMockAssetRepository() *MockAssetRepository {
	return &MockAssetRepository{
		assets: make(map[string]*domain.Asset),
	}
}

func (r *MockAssetRepository) Add(asset domain.Asset) error {
	r.assets[asset.Ticker] = &asset
	return nil
}

func (r *MockAssetRepository) GetByTicker(ticker string) (*domain.Asset, error) {
	asset, exists := r.assets[ticker]
	if !exists {
		return nil, errors.New("asset not found")
	}
	return asset, nil
}

func (r *MockAssetRepository) Remove(ticker string) error {
	_, exists := r.assets[ticker]
	if !exists {
		return errors.New("asset not found")
	}
	delete(r.assets, ticker)
	return nil
}
