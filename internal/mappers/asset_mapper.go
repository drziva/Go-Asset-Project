package mappers

import (
	"go-project/internal/dto"
	"go-project/internal/models"
)

func ToAssetResponse(asset models.Asset) dto.AssetResponse {
	return dto.AssetResponse{
		ID:          asset.ID,
		Name:        asset.Name,
		Description: asset.Description,
		UserID:      asset.UserID,
	}
}
