package models

type Asset struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string

	UserID uint // the relation gets defined in the "User" model
	User   User `grom:"constraint:onDelete:CASCADE;"`
}
