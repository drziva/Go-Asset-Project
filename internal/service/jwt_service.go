package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret []byte
	ttl    time.Duration
}

type AccessTokenClaims struct {
	ID      uint `json:"id"`
	IsAdmin bool `json:"is_admin"`

	jwt.RegisteredClaims
}

func NewJWTService(secret string, ttlSeconds int) *JWTService {
	return &JWTService{
		secret: []byte(secret),
		ttl:    time.Duration(ttlSeconds) * time.Second,
	}
}

func (s *JWTService) GenerateAccessToken(ID uint, isAdmin bool) (string, error) {
	now := time.Now()

	claims := AccessTokenClaims{
		ID:      ID,
		IsAdmin: isAdmin,

		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(ID)),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.ttl)),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.secret)
}

func (s *JWTService) ValidateAccessToken(tokenString string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&AccessTokenClaims{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return s.secret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
