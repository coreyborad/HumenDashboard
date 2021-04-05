package services

import "errors"

// Error constants
var (
	ErrAuthFailed        = errors.New("Client Authentication failed")
	ErrUnsupportedClaims = errors.New("Unsupported claims")
	ErrTokenInvalid      = errors.New("Token is invalid")
	ErrTokenExpired      = errors.New("Token is expired")
)
