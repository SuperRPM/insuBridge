package repository

import (
	"insuBridge/domain"
	"insuBridge/internal/config"
	"insuBridge/internal/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{
		db: config.DB,
	}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}
