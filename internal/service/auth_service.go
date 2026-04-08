package service

import (
	appErrors "go-project/internal/errors"
	dbErrors "go-project/internal/service/errors"

	"go-project/internal/dto"
	"go-project/internal/models"
	"go-project/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo       *repository.UserRepository
	jwtService *JWTService
}

func NewAuthservice(repo *repository.UserRepository, jwtService *JWTService) *AuthService {
	return &AuthService{
		repo,
		jwtService,
	}
}

func (s *AuthService) SignUp(dto dto.SignUpDTO) (*models.User, error) {
	hashedPassword, err := hashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hashedPassword,
	}

	//CREATE USER AND MAP POTENTIAL ERROR
	err = s.repo.CreateUser(user)
	return user, dbErrors.MapDBError(err)
}

func (s *AuthService) Login(dto dto.LoginDTO) (*models.User, string, error) {
	user, err := s.repo.GetUserByEmail(dto.Email)

	if err != nil {
		return nil, "", dbErrors.MapDBError(err)
	}

	passwordMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if passwordMatch != nil {
		return nil, "", appErrors.ErrInvalidCredentials
	}

	token, err := s.jwtService.GenerateAccessToken(user.ID, user.IsAdmin)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) Me(id uint) (*models.User, error) {
	user, err := s.repo.GetUserById(id)

	if err != nil {
		return nil, dbErrors.MapDBError(err)
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(hash), err
}
