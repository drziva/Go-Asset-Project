package repository

import (
	"go-project/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User

	err := r.db.
		Where("email = ?", email).
		First(&user).
		Error

	return user, err
}

func (r *UserRepository) GetUserById(id uint) (*models.User, error) {
	var user *models.User

	err := r.db.
		Where("id = ?", id).
		First(&user).
		Error

	return user, err
}
