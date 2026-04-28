package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-project/internal/constants"
	appErrors "go-project/internal/errors"
	dbErrors "go-project/internal/service/errors"
	"net/http"

	"go-project/internal/dto"
	"go-project/internal/models"
	"go-project/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type AuthService struct {
	repo         *repository.UserRepository
	googleConfig *oauth2.Config
	jwtService   *JWTService
}

type AuthResult struct {
	User      *models.User
	Linked    bool
	LinkToken string
}

func NewAuthservice(repo *repository.UserRepository, googleConfig *oauth2.Config, jwtService *JWTService) *AuthService {
	return &AuthService{
		repo,
		googleConfig,
		jwtService,
	}
}

func (s *AuthService) SignUp(dto dto.SignUpDTO) (*AuthResult, error) {
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

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		user, err = s.repo.GetUserByEmail(user.Email)
		if err != nil {
			fmt.Print("----------------USER----------------", user)
			if user.AuthProvider == constants.AuthProviderGoogle {
				linkToken, err := s.jwtService.GenerateLinkToken(user.Email)
				if err != nil {
					return nil, err
				}
				return &AuthResult{User: nil, Linked: false, LinkToken: linkToken}, err
			}
		}
	}

	return &AuthResult{User: user, Linked: false, LinkToken: "obaj"}, dbErrors.MapDBError(err)
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

func (s *AuthService) GetGoogleLoginURL() string {
	return s.googleConfig.AuthCodeURL("random-state")
}

func (s *AuthService) HandleGoogleCallback(ctx context.Context, code string) (*AuthResult, string, error) {
	token, err := s.googleConfig.Exchange(ctx, code)
	if err != nil {
		return nil, "", err
	}

	client := s.googleConfig.Client(ctx, token)

	resp, err := client.Get("https://openidconnect.googleapis.com/v1/userinfo")
	if err != nil {
		return nil, "", err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("google userinfo fialed: %s", resp.Status)
	}
	defer resp.Body.Close()

	var data struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, "", err
	}

	user, err := s.repo.GetUserByEmail(data.Email)

	if err == nil {
		if user.AuthProvider == constants.AuthProviderLocal {
			linkToken, err := s.jwtService.GenerateLinkToken(user.Email)
			if err != nil {
				return nil, "", err
			}
			return &AuthResult{User: user, Linked: false, LinkToken: linkToken}, "", err
		}
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = &models.User{
				Email:        data.Email,
				Name:         data.Name,
				AuthProvider: constants.AuthProviderGoogle,
			}
			err = s.repo.CreateUser(user)
			if err != nil {
				return nil, "", dbErrors.MapDBError(err)
			}
		} else {
			return nil, "", dbErrors.MapDBError(err)
		}
	}

	jwt, err := s.jwtService.GenerateAccessToken(user.ID, user.IsAdmin)
	if err != nil {
		return nil, "", err
	}

	return &AuthResult{User: user, Linked: true, LinkToken: ""}, jwt, nil
}

func (s *AuthService) LinkAndLogin(linkRequest dto.LinkRequest) (*models.User, string, error) {
	tokenClaims, err := s.jwtService.ValidateLinkToken(linkRequest.LinkToken)
	if err != nil {
		return nil, "", err
	}

	userEmail := tokenClaims.Email

	loginDTO := &dto.LoginDTO{
		Email:    userEmail,
		Password: linkRequest.Password,
	}

	return s.Login(*loginDTO)
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(hash), err
}
