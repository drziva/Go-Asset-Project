package repository

import (
	"go-project/internal/dto"
	"go-project/internal/models"

	"gorm.io/gorm"
)

type VerificationCodeRepository struct {
	db *gorm.DB
}

func NewVerificationCodeRepository(db *gorm.DB) *VerificationCodeRepository {
	return &VerificationCodeRepository{
		db,
	}
}

func (r *VerificationCodeRepository) CreateCode(userID uint, codeHash string, codeType dto.RequestType) error {
	verificationCode := &models.VerificationCode{
		UserID:   userID,
		CodeHash: codeHash,
		Type:     codeType,
	}

	err := r.db.Create(&verificationCode).Error

	return err
}

func (r *VerificationCodeRepository) DeleteCode(userID, ID uint) error {
	var code models.VerificationCode
	err := r.db.Where("id = ? AND user_id = ?", userID, ID).First(&code).Error

	return err
}
