package repository

import (
	"go-project/internal/dto"
	"go-project/internal/models"
	"time"

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
		UserID:    userID,
		CodeHash:  codeHash,
		Type:      codeType,
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}

	err := r.db.Create(&verificationCode).Error

	return err
}

func (r *VerificationCodeRepository) GetCodeForUser(userID uint, codeType dto.RequestType) (*models.VerificationCode, error) {
	var code models.VerificationCode
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").First(&code).Error

	return &code, err
}

func (r *VerificationCodeRepository) DeleteCode(userID, ID uint) error {
	err := r.db.Where("id = ? AND user_id = ?", ID, userID).Delete(&models.VerificationCode{}).Error

	return err
}
