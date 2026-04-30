package models

import (
	"go-project/internal/dto"
	"time"
)

type VerificationCode struct {
	ID          uint            `gorm:"primaryKey"`
	UserID      uint            `gorm:"not null;index"`
	CodeHash    string          `gorm:"type:text;not null"`
	Type        dto.RequestType `gorm:"type:text;not null;index"` // verify_email / reset_password / link_account
	ExpiresAt   time.Time       `gorm:"not null"`
	Attempts    int             `gorm:"default:0"`
	MaxAttempts int             `gorm:"default:5"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}
