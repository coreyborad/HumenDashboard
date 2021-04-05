package services

import (
	"stock/auth"
	"stock/models"
	"stock/repositories"
)

// UserService User Service
type UserService struct {
	userRep *repositories.UserRepository
}

// NewUserService New User Service
func NewUserService(
	userRep *repositories.UserRepository,
) *UserService {
	return &UserService{
		userRep: userRep,
	}
}

// GetUserFromToken GetUserFromToken
func (s *UserService) GetUserFromToken(tokenString string) (*models.User, error) {
	claims, err := auth.Decode(tokenString)
	if err != nil {
		return nil, err
	}
	if claims.UserID > 0 {
		user, err := s.userRep.Find(claims.UserID)
		if err != nil {
			return nil, err
		}
		token, err := auth.ParseToken(tokenString)
		if err != nil {
			return nil, err
		}

		if token.Valid {
			return user, nil
		}

		return nil, ErrTokenInvalid
	}

	return nil, ErrUnsupportedClaims
}
