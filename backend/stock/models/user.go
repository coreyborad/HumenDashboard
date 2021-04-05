package models

import (
	"errors"
	"time"
)

// Error constants
var (
	ErrUserNotExist = errors.New("User Not Exist")
)

type User struct {
	ID              uint64     `json:"id" gorm:"primaryKey"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Password        string     `json:"password"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}
