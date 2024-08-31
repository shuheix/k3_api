package repository

import (
	"k3_api/internal/domain/models"

	"gorm.io/gorm"
)

type UserRepository interface{
	List() ([]*models.User, error)
	Create(user *models.User) error
	Delete(user *models.User) error
}

type userRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) List() ([]*models.User, error) {
	var users []*models.User
	result := repo.db.Find(&users)
	
	return users, result.Error
}

func (repo *userRepository) Create(user *models.User) (error) {
	result := repo.db.Create(user)

	return result.Error
}

func (repo *userRepository) Delete(user *models.User) (error) {
	result := repo.db.Delete(user)

	return result.Error
}