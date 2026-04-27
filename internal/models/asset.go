package models

type Asset struct {
	ID          uint   `gorm:"primaryKey"`
	FileName    string `gorm:"not null"`
	StoredName  string `gorm:"not null"`
	FilePath    string `gorm:"not null"`
	MimeType    string
	FileSize    int64
	Description string

	UserID uint // the relation gets defined in the "User" model
	User   User `gorm:"constraint:onDelete:CASCADE;"`
}
