package repositories

import (
	"golang-template/app/models"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (m *UserRepositoryMock) Create(user *models.UserRegister) error {
	args := m.Mock.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) Update(user *models.UserUpdatePassword) error {
	args := m.Mock.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) List() (*[]models.User, error) {
	args := m.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]models.User), args.Error(1)
}
