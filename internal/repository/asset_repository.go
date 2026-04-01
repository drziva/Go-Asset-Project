package repository

import (
	"go-project/internal/models"

	"gorm.io/gorm"
)

type AssetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) *AssetRepository {
	return &AssetRepository{
		db,
	}
}

func (r *AssetRepository) CreateAsset(userId uint, asset *models.Asset) error {
	return r.db.Create(asset).Error
}

func (r *AssetRepository) GetAssetsForUser(userId uint) ([]models.Asset, error) {
	var assets []models.Asset

	err := r.db.Where("id = ?", userId).Find(&assets).Error

	return assets, err
}

func (r *AssetRepository) GetAssetById(userId, id uint) (models.Asset, error) {
	var asset models.Asset

	err := r.db.Where("id = ? AND user_id = ?", id, userId).First(&asset).Error

	return asset, err
}

func (r *AssetRepository) GetAllAssets() ([]models.Asset, error) {
	var assets []models.Asset

	err := r.db.Find(&assets).Error

	return assets, err
}
