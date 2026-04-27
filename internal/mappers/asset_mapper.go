package mappers

import (
	"fmt"
	"go-project/internal/dto"
	"go-project/internal/models"
)

func ToAssetResponse(asset models.Asset) dto.AssetResponse {
	return dto.AssetResponse{
		ID:          asset.ID,
		FileName:    asset.FileName,
		FileSize:    formatFileSize(asset.FileSize),
		MimeType:    asset.MimeType,
		Description: asset.Description,
		UserID:      asset.UserID,
	}
}

func formatFileSize(size int64) string {
	const unit = 1024

	if size < unit {
		return fmt.Sprintf("%d B", size)
	}

	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
