package usecase

import (
	"user-management/entity"
	"user-management/entity/response"
	"user-management/repository"

	"github.com/jinzhu/copier"
)

type IUserUsecase interface {
	CreateUser(user response.CreateUserRequest) error
	GetListUser() ([]response.GetUserResponse, error)
	DeleteUser(user response.DeleteUserRequest) error
}

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUsecase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository}
}

func (u UserUsecase) CreateUser(req response.CreateUserRequest) error {
	user := entity.User{}
	copier.Copy(&user, &req)

	err := u.userRepository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUsecase) GetList() ([]response.GetUserResponse, error) {
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, nil
	}
	userResponse := []response.GetUserResponse{}
	copier.Copy(&userResponse, &users)
	return userResponse, nil
}

func (u UserUsecase) UpdateUser(req response.UpdateUserRequest) error {
	user := entity.User{}
	copier.Copy(&user, &req)

	err := u.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUsecase) DeleteUser(req response.DeleteUserRequest) error {
	if err := u.userRepository.GetByID(req); err != nil {
		return err
	}

	user := entity.User{}
	copier.Copy(&user, &req)

	if err := u.userRepository.Delete(user); err != nil {
		return err
	}
	return nil
}
