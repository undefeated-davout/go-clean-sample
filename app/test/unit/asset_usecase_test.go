package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/undefeated-davout/portfolio-simulator/app/domain"
	"github.com/undefeated-davout/portfolio-simulator/app/test/mock"
	"github.com/undefeated-davout/portfolio-simulator/app/usecase/asset"
)

func TestAddAsset(t *testing.T) {
	// モックリポジトリのセットアップ
	mockRepo := mock.NewMockAssetRepository()
	addAssetUC := asset.NewAddAssetUseCase(mockRepo)

	// アセット追加のテスト
	err := addAssetUC.AddAsset("AAPL", 0.5)
	assert.NoError(t, err, "Error should not have occurred when adding asset")

	// アセットが正しく追加されたか確認
	result, err := mockRepo.GetByTicker("AAPL")
	assert.NoError(t, err, "Asset should have been found")
	assert.Equal(t, "AAPL", result.Ticker)
	assert.Equal(t, 0.5, result.Weight)
}

func TestRemoveAsset(t *testing.T) {
	// モックリポジトリのセットアップ
	mockRepo := mock.NewMockAssetRepository()
	removeAssetUC := asset.NewRemoveAssetUseCase(mockRepo)

	// 事前にアセットを追加
	mockRepo.Add(domain.Asset{Ticker: "AAPL", Weight: 0.5})

	// アセット削除のテスト
	err := removeAssetUC.RemoveAsset("AAPL")
	assert.NoError(t, err, "Error should not have occurred when removing asset")

	// アセットが削除されたか確認
	_, err = mockRepo.GetByTicker("AAPL")
	assert.Error(t, err, "Asset should not be found after deletion")
}
