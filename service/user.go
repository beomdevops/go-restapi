package service

import (
	"github.com/beomdevops/go-restapi/dto"
	"github.com/beomdevops/go-restapi/models"
	"github.com/beomdevops/go-restapi/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(di_userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: di_userRepo,
	}
}

func (u *UserService) SignUser(createUser *dto.CreateUserRequest) (*models.User, error) {
	user, err := u.userRepo.CreateUser(createUser.ToUserEntity())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetUserById(id int) (*models.User, error) {
	user, err := u.userRepo.FindById(id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetUserByName(name string) (*models.User, error) {
	user, err := u.userRepo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
