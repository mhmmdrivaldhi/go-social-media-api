package repository

import (
	"github.com/mhmmdrivaldhi/go-social-media-api/model/entity"
	"gorm.io/gorm"
)

type AuthRepository interface{
	Register(user *entity.User) error
	IsEmailExists(email string) bool
	GetUserByEmail(email string) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (ar *authRepository) IsEmailExists(email string) bool {
	var user entity.User
	err := ar.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return false
	}

	return err == nil
}

func (ar *authRepository) Register(user *entity.User) error {
	err := ar.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (ar *authRepository) GetUserByEmail(email string) (*entity.User, error) {
	err := ar.db.Where("email = ?", email).First(&entity.User{}).Error
	if err != nil {
		return nil, err
	}

	return &entity.User{}, nil
}