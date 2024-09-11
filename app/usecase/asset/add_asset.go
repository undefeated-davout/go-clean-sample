package asset

import (
	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type AddAssetUseCase struct {
	assetRepo domain.AssetRepository
}

func NewAddAssetUseCase(assetRepo domain.AssetRepository) *AddAssetUseCase {
	return &AddAssetUseCase{assetRepo: assetRepo}
}

func (uc *AddAssetUseCase) AddAsset(ticker string, weight float64) error {
	asset := domain.Asset{
		Ticker: ticker,
		Weight: weight,
	}

	// リポジトリを使って資産を追加
	return uc.assetRepo.Add(asset)
}
