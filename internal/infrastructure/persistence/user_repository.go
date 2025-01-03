package persistence

import (
	"context"
	"k3_api/internal/domain/model"
	"k3_api/internal/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(ctx context.Context, user *model.User) error {
	return r.db.Save(user).Error
}
