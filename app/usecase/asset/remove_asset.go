package asset

import (
	"github.com/undefeated-davout/portfolio-simulator/app/domain"
)

type RemoveAssetUseCase struct {
	assetRepo domain.AssetRepository
}

func NewRemoveAssetUseCase(assetRepo domain.AssetRepository) *RemoveAssetUseCase {
	return &RemoveAssetUseCase{assetRepo: assetRepo}
}

func (uc *RemoveAssetUseCase) RemoveAsset(ticker string) error {
	return uc.assetRepo.Remove(ticker)
}
