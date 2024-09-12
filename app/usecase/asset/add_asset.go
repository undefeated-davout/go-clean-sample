package asset

import (
	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type AddAssetUseCase struct {
	assetRepo domain.AssetRepository
}

type AssetAdder interface {
	AddAsset(ticker string, weight float64) error
}

var _ AssetAdder = (*AddAssetUseCase)(nil)

func NewAddAssetUseCase(assetRepo domain.AssetRepository) *AddAssetUseCase {
	return &AddAssetUseCase{assetRepo: assetRepo}
}

func (uc *AddAssetUseCase) AddAsset(ticker string, weight float64) error {
	asset := domain.Asset{
		Ticker: ticker,
		Weight: weight,
	}

	return uc.assetRepo.Add(asset)
}
