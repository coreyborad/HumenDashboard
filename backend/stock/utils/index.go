package utils

import (
	"github.com/google/uuid"
)

// UUID UUID
func UUID() string {
	return uuid.New().String()
}
