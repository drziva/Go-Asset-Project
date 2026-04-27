package dto

type UpdateAssetDTO struct {
	FileName    string `json:"name"binding:"required"`
	Description string `json:"description"binding:"required"`
}

type ServiceCreateAssetDTO struct {
	FileName    string
	StoredName  string
	FilePath    string
	FileSize    int64
	MimeType    string
	Description string
}

type AssetResponse struct {
	ID          uint   `json:"id"`
	FileName    string `json:"name"`
	FileSize    string `json:"file_size"`
	MimeType    string `json:"mime_type"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}
