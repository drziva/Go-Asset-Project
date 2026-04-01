package dto

type CreateAssetDTO struct {
	Name        string `json:"name"binding:"required"`
	Description string `json:"description"binding:"required"`
}

type AssetResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}
