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

	err := r.db.Where("user_id = ?", userId).Find(&assets).Error

	return assets, err
}

func (r *AssetRepository) GetAssetById(userId, id uint) (*models.Asset, error) {
	var asset models.Asset

	err := r.db.Where("id = ? AND user_id = ?", id, userId).First(&asset).Error

	return &asset, err
}

func (r *AssetRepository) UpdateAsset(userID, ID uint, asset *models.Asset) (*models.Asset, error) {
	dbAsset := &models.Asset{}

	// db.Update does NOT load the data into the passed struct(dbAsset)
	err := r.db.
		Model(dbAsset).
		Where("id = ? AND user_id = ?", ID, userID).
		Updates(map[string]interface{}{
			"name":        asset.Name,
			"description": asset.Description,
		}).
		Error

	err = r.db.Where("id = ? AND user_id = ?", ID, userID).First(dbAsset).Error

	return dbAsset, err
}

func (r *AssetRepository) DeleteAsset(userID uint, ID uint) error {
	dbAsset := &models.Asset{}

	err := r.db.Where("id = ? AND user_id = ?", ID, userID).First(dbAsset).Error
	if err != nil {
		return err
	}

	return r.db.Delete(dbAsset).Error
}

// ADMIN FUNCTIONS
func (r *AssetRepository) GetAllAssets() ([]models.Asset, error) {
	var assets []models.Asset

	err := r.db.Find(&assets).Error

	return assets, err
}

func (r *AssetRepository) GetAnyAssetById(ID uint) (*models.Asset, error) {
	dbAsset := &models.Asset{}

	err := r.db.Where("id = ?", ID).First(dbAsset).Error
	return dbAsset, err
}

func (r *AssetRepository) UpdateAnyAsset(ID uint, asset *models.Asset) (*models.Asset, error) {
	dbAsset := &models.Asset{}

	err := r.db.
		Model(dbAsset).
		Where("id = ?", ID).
		Updates(map[string]interface{}{
			"name":        asset.Name,
			"description": asset.Description,
		}).
		Error

	err = r.db.Where("id = ?", ID).First(dbAsset).Error

	return dbAsset, err
}

func (r *AssetRepository) DeleteAnyAsset(ID uint) error {
	dbAsset := &models.Asset{}

	err := r.db.Where("id = ?", ID).First(&dbAsset).Error
	if err != nil {
		return err
	}

	return r.db.Delete(dbAsset).Error
}
