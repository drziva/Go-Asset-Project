package service

import (
	"errors"
	"go-project/internal/dto"
	appErrors "go-project/internal/errors"
	"go-project/internal/models"
	"go-project/internal/repository"

	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthservice(repo *repository.UserRepository) *AuthService {
	return &AuthService{
		repo,
	}
}

func (s *AuthService) SignUp(user *models.User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	//CREATE USER AND MAP POTENTIAL ERROR
	err = s.repo.CreateUser(user)

	return mapDBError(err)
}

func (s *AuthService) Login(dto dto.LoginDTO) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(dto.Email)

	if err != nil {
		return nil, err
	}

	passwordMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if passwordMatch != nil {
		return nil, appErrors.ErrInvalidCredentials
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

func mapDBError(err error) error {

	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {

		switch pgErr.Code {

		case "23505":
			return appErrors.ErrEmailAlreadyExists
		}
	}

	return err
}
