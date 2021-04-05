package repositories

import (
	"stock/models"

	"github.com/jinzhu/gorm"
)

// UserHmiRepository UserHmi Repository
type UserRepository struct {
	db *gorm.DB
}

// UserRepository New UserHmi UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Find Find
func (r *UserRepository) Find(id uint64) (*models.User, error) {
	user := new(models.User)
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}
