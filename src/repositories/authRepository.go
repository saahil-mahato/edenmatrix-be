package repositories

import (
	"github.com/saahil-mahato/edenmatrix-be/src/database"
	"github.com/saahil-mahato/edenmatrix-be/src/models"
)

type AuthRepository struct{}

func (r *AuthRepository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *AuthRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error

	return &user, err
}
