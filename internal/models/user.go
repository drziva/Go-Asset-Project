package models

import (
	"go-project/internal/constants"
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"not null;uniqueIndex"`
	Password     string
	IsAdmin      bool `gorm:"not null;default:false"`
	CreatedAt    time.Time
	AuthProvider constants.AuthProvider `gorm:"type:varchar(20);not null;default:local;check:provider IN ('local','google')"`

	Assets            []Asset            `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE;"` // O:M
	VerificationCodes []VerificationCode `gorm:"foreignKey:UserID"`
}
