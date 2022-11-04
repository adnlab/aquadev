package repository

import (
	"user-management/entity"
	"user-management/entity/response"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user entity.User) error
	GetAll() ([]entity.User, error)
	GetByID(req response.DeleteUserRequest) error
	Update(user entity.User) error
	Delete(user entity.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) Create(user entity.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) GetAll() ([]entity.User, error) {
	var users []entity.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, nil
	}
	return users, nil
}

func (u UserRepository) GetByID(req response.DeleteUserRequest) error {
	var user entity.User
	if err := u.db.First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) Update(user entity.User) error {
	if err := u.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) Delete(user entity.User) error {
	if err := u.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
