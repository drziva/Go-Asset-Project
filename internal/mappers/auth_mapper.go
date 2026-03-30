package mappers

import (
	"go-project/internal/dto"
	"go-project/internal/models"
)

func ToLoginResponse(user models.User) dto.LoginResponse {
	return dto.LoginResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
