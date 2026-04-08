package service

import (
	"go-project/internal/dto"
	"go-project/internal/models"
	"go-project/internal/repository"
	dbErrors "go-project/internal/service/errors"
)

type AssetService struct {
	repo *repository.AssetRepository
}

func NewAssetService(repo *repository.AssetRepository) *AssetService {
	return &AssetService{
		repo,
	}
}

func (s *AssetService) CreateAsset(userID uint, dto dto.CreateAssetDTO) (*models.Asset, error) {
	asset := &models.Asset{
		UserID:      userID,
		Name:        dto.Name,
		Description: dto.Description,
	}

	err := s.repo.CreateAsset(userID, asset)
	if err != nil {
		return nil, dbErrors.MapDBError(err)
	}

	return asset, err
}

func (s *AssetService) GetAssetsForUser(userID uint) ([]models.Asset, error) {
	assets, err := s.repo.GetAssetsForUser(userID)
	if err != nil {
		return nil, dbErrors.MapDBError(err)
	}

	return assets, err
}

func (s *AssetService) GetAssetById(userID uint, ID uint) (*models.Asset, error) {
	asset, err := s.repo.GetAssetById(userID, ID)
	if err != nil {
		return nil, dbErrors.MapDBError(err)
	}

	return asset, err
}

func (s *AssetService) GetAllAssets() ([]models.Asset, error) {
	assets, err := s.repo.GetAllAssets()
	if err != nil {
		return nil, dbErrors.MapDBError(err)
	}

	return assets, err
}

func (s *AssetService) GetAnyAssetById(ID uint) (*models.Asset, error) {
	asset, err := s.repo.GetAnyAssetById(ID)
	if err != nil {
		return nil, dbErrors.MapDBError(err)
	}

	return asset, err
}

func (s *AssetService) UpdateAsset(userID, ID uint, dto dto.UpdateAssetDTO) (*models.Asset, error) {
	asset := &models.Asset{
		UserID:      userID,
		Name:        dto.Name,
		Description: dto.Description,
	}

	asset, err := s.repo.UpdateAsset(userID, ID, asset)
	if err != nil {
		return nil, dbErrors.MapDBError(err)
	}

	return asset, err
}

func (s *AssetService) UpdateAnyAsset(ID uint, dto dto.UpdateAssetDTO) (*models.Asset, error) {
	asset := &models.Asset{
		Name:        dto.Name,
		Description: dto.Description,
	}

	asset, err := s.repo.UpdateAnyAsset(ID, asset)
	if err != nil {
		return nil, dbErrors.MapDBError(err)
	}

	return asset, err
}

func (s *AssetService) DeleteAsset(userID, ID uint) error {
	return dbErrors.MapDBError(s.repo.DeleteAsset(userID, ID))
}
