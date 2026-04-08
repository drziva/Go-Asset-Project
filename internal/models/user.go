package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null;uniqueIndex"`
	Password  string `gorm:"not null"`
	IsAdmin   bool   `gorm:"not null;default:false"`
	CreatedAt time.Time

	Assets []Asset `grom:"foreignKey:UserID"` // O:M
}
