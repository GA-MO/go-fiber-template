package services

import (
	"golang-template/app/models"
	"golang-template/app/repositories"
)

type UserService interface {
	Register(user *models.UserRegister) error
	Update(user *models.UserUpdatePassword) error
	List() (*[]models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) Register(user *models.UserRegister) error {
	return s.userRepository.Create(user)
}

func (s *userService) Update(user *models.UserUpdatePassword) error {
	return s.userRepository.Update(user)
}

func (s *userService) List() (*[]models.User, error) {
	return s.userRepository.List()
}
