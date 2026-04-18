package repository

import (
	"blogSystem/dto"
	"blogSystem/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByUserName(username string) (*dto.UserDto, error) {
	var userDto dto.UserDto
	result := r.db.Model(&model.User{}).Where("username = ?", username).Scan(&userDto)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &userDto, nil
}

func (r *userRepository) Create(user *model.User) error {
	result := r.db.Model(&model.User{}).Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
