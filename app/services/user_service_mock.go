package services

import (
	"golang-template/app/models"

	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func NewUserServiceMock() *UserServiceMock {
	return &UserServiceMock{}
}

func (m *UserServiceMock) Register(user *models.UserRegister) error {
	args := m.Mock.Called(user)
	return args.Error(0)
}

func (m *UserServiceMock) Update(user *models.UserUpdatePassword) error {
	args := m.Mock.Called(user)
	return args.Error(0)
}

func (m *UserServiceMock) List() (*[]models.User, error) {
	args := m.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]models.User), args.Error(1)
}
