package services

import (
	"golang-template/app/models"
	"golang-template/app/repositories"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_Register(t *testing.T) {
	testCaseList := []struct {
		name          string
		data          *models.UserRegister
		mockSetup     func(*repositories.UserRepositoryMock)
		expectedError error
	}{
		{
			name: "successful registration",
			data: &models.UserRegister{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			mockSetup: func(m *repositories.UserRepositoryMock) {
				m.On("Create", mock.Anything).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "repository error",
			data: &models.UserRegister{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			mockSetup: func(m *repositories.UserRepositoryMock) {
				m.On("Create", mock.Anything).Return(assert.AnError)
			},
			expectedError: assert.AnError,
		},
	}

	for _, testCase := range testCaseList {
		t.Run(testCase.name, func(t *testing.T) {
			repoMock := repositories.NewUserRepositoryMock()
			testCase.mockSetup(repoMock)

			service := NewUserService(repoMock)
			err := service.Register(testCase.data)

			assert.Equal(t, testCase.expectedError, err)
			repoMock.AssertExpectations(t)
		})
	}
}

func TestUserService_Update(t *testing.T) {
	testCaseList := []struct {
		name          string
		user          *models.UserUpdatePassword
		mockSetup     func(*repositories.UserRepositoryMock)
		expectedError error
	}{
		{
			name: "successful update",
			user: &models.UserUpdatePassword{
				Username:    "testuser",
				NewPassword: "newpassword123",
			},
			mockSetup: func(m *repositories.UserRepositoryMock) {
				m.On("Update", mock.Anything).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "repository error",
			user: &models.UserUpdatePassword{
				Username:    "testuser",
				NewPassword: "newpassword123",
			},
			mockSetup: func(m *repositories.UserRepositoryMock) {
				m.On("Update", mock.Anything).Return(assert.AnError)
			},
			expectedError: assert.AnError,
		},
	}

	for _, testCase := range testCaseList {
		t.Run(testCase.name, func(t *testing.T) {
			repoMock := repositories.NewUserRepositoryMock()
			testCase.mockSetup(repoMock)

			service := NewUserService(repoMock)
			err := service.Update(testCase.user)

			assert.Equal(t, testCase.expectedError, err)
			repoMock.AssertExpectations(t)
		})
	}
}

func TestUserService_List(t *testing.T) {
	testTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	testCaseList := []struct {
		name          string
		mockSetup     func(*repositories.UserRepositoryMock)
		expectedUsers *[]models.User
		expectedError error
	}{
		{
			name: "successful list",
			mockSetup: func(m *repositories.UserRepositoryMock) {
				users := &[]models.User{
					{
						Username:  "user1",
						Email:     "user1@example.com",
						CreatedAt: testTime,
						UpdatedAt: testTime,
					},
				}
				m.On("List").Return(users, nil)
			},
			expectedUsers: &[]models.User{
				{
					Username:  "user1",
					Email:     "user1@example.com",
					CreatedAt: testTime,
					UpdatedAt: testTime,
				},
			},
			expectedError: nil,
		},
		{
			name: "repository error",
			mockSetup: func(m *repositories.UserRepositoryMock) {
				m.On("List").Return(nil, assert.AnError)
			},
			expectedUsers: nil,
			expectedError: assert.AnError,
		},
	}

	for _, testCase := range testCaseList {
		t.Run(testCase.name, func(t *testing.T) {
			repoMock := repositories.NewUserRepositoryMock()
			testCase.mockSetup(repoMock)

			service := NewUserService(repoMock)
			users, err := service.List()

			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedUsers, users)
			repoMock.AssertExpectations(t)
		})
	}
}
