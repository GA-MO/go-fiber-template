package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"golang-template/app/models"
	"golang-template/app/services"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name               string
	url                string
	method             string
	jsonBody           string
	expectedStatusCode int
	mockFunc           func(userServiceMock *services.UserServiceMock)
}

func TestUserHandler(t *testing.T) {

	testCaseList := []testCase{
		{
			name:               "Register Success",
			url:                "/register",
			method:             fiber.MethodPost,
			jsonBody:           `{"username": "test", "email": "test@test.com", "password": "test"}`,
			expectedStatusCode: 200,
			mockFunc: func(serviceMock *services.UserServiceMock) {
				serviceMock.On("Register", mock.Anything).Return(nil).Once()
			},
		},
		{
			name:               "Register Failed",
			url:                "/register",
			method:             fiber.MethodPost,
			jsonBody:           `{"username": "", "email": "test@test.com", "password": "test"}`,
			expectedStatusCode: 400,
			mockFunc: func(serviceMock *services.UserServiceMock) {
				serviceMock.On("Register", mock.Anything).Return(errors.New("error")).Once()
			},
		},
		{
			name:               "Register Body Empty",
			url:                "/register",
			method:             fiber.MethodPost,
			jsonBody:           "",
			expectedStatusCode: 400,
			mockFunc:           func(serviceMock *services.UserServiceMock) {},
		},
		{
			name:               "Update Password Success",
			url:                "/update",
			method:             fiber.MethodPut,
			jsonBody:           `{"username": "test", "newPassword": "test"}`,
			expectedStatusCode: 200,
			mockFunc: func(serviceMock *services.UserServiceMock) {
				serviceMock.On("Update", mock.Anything).Return(nil).Once()
			},
		},
		{
			name:               "Update Password Failed",
			url:                "/update",
			method:             fiber.MethodPut,
			jsonBody:           `{"username": "test", "newPassword": ""}`,
			expectedStatusCode: 400,
			mockFunc: func(serviceMock *services.UserServiceMock) {
				serviceMock.On("Update", mock.Anything).Return(errors.New("error")).Once()
			},
		},
		{
			name:               "Update Password Body Empty",
			url:                "/update",
			method:             fiber.MethodPut,
			jsonBody:           "",
			expectedStatusCode: 400,
			mockFunc:           func(serviceMock *services.UserServiceMock) {},
		},
		{
			name:               "List Success",
			url:                "/list",
			method:             fiber.MethodGet,
			expectedStatusCode: 200,
			mockFunc: func(serviceMock *services.UserServiceMock) {
				serviceMock.On("List", mock.Anything).Return(&[]models.User{}, nil).Once()
			},
		},
		{
			name:               "List Failed",
			url:                "/list",
			method:             fiber.MethodGet,
			expectedStatusCode: 500,
			mockFunc: func(serviceMock *services.UserServiceMock) {
				serviceMock.On("List", mock.Anything).Return(nil, errors.New("error")).Once()
			},
		},
	}

	app := fiber.New()
	userServiceMock := services.NewUserServiceMock()
	handler := NewUserHandler(userServiceMock)
	group := "/api/v1/user"
	RegisterUserRoutes(app.Group(group), handler)

	for _, testCase := range testCaseList {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockFunc(userServiceMock)
			req, _ := http.NewRequest(testCase.method, group+testCase.url, bytes.NewBufferString(testCase.jsonBody))
			req.Header.Set("Content-Type", "application/json")
			res, _ := app.Test(req, -1)
			body, _ := io.ReadAll(res.Body)
			fmt.Println(string(body))
			assert.Equal(t, testCase.expectedStatusCode, res.StatusCode)
		})
	}

}
