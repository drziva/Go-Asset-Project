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

func (s *AssetService) CreateAsset(userId uint, dto dto.CreateAssetDTO) (*models.Asset, error) {
	asset := &models.Asset{
		UserID:      userId,
		Name:        dto.Name,
		Description: dto.Description,
	}

	err := s.repo.CreateAsset(userId, asset)
	if err != nil {
		return nil, dbErrors.MapDBError(err)
	}

	return asset, err
}
