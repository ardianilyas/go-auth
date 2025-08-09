package repositories

import (
	model "github.com/ardianilyas/go-auth/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindById(id uint) (*model.User, error) {
	var user model.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *UserRepository) UpdateRefreshToken(userID uint, token string) error {
	return r.DB.Model(&model.User{}).Where("id = ?", userID).Update("refresh_token", token).Error
}