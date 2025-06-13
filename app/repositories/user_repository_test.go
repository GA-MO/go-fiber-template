package repositories

import (
	"database/sql"
	"golang-template/app/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %v", err)
	}
	return db, mock
}

func TestUserRepository_Create(t *testing.T) {
	db, mock := setupTestDB(t)
	defer db.Close()

	repo := NewUserRepository(db)

	testCaseList := []struct {
		name          string
		user          *models.UserRegister
		mockSetup     func(sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "successful creation",
			user: &models.UserRegister{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO users").
					WithArgs("testuser", "test@example.com", "password123").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "database error",
			user: &models.UserRegister{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO users").
					WithArgs("testuser", "test@example.com", "password123").
					WillReturnError(sql.ErrConnDone)
			},
			expectedError: sql.ErrConnDone,
		},
	}

	for _, testCase := range testCaseList {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockSetup(mock)
			err := repo.Create(testCase.user)
			assert.Equal(t, testCase.expectedError, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUserRepository_Update(t *testing.T) {
	db, mock := setupTestDB(t)
	defer db.Close()

	repo := NewUserRepository(db)

	testCaseList := []struct {
		name          string
		user          *models.UserUpdatePassword
		mockSetup     func(sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "successful update",
			user: &models.UserUpdatePassword{
				Username:    "testuser",
				NewPassword: "newpassword123",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE users").
					WithArgs("testuser", "newpassword123", "testuser").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "database error",
			user: &models.UserUpdatePassword{
				Username:    "testuser",
				NewPassword: "newpassword123",
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE users").
					WithArgs("testuser", "newpassword123", "testuser").
					WillReturnError(sql.ErrConnDone)
			},
			expectedError: sql.ErrConnDone,
		},
	}

	for _, testCase := range testCaseList {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockSetup(mock)
			err := repo.Update(testCase.user)
			assert.Equal(t, testCase.expectedError, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}

func TestUserRepository_List(t *testing.T) {
	db, mock := setupTestDB(t)
	defer db.Close()

	repo := NewUserRepository(db)
	testTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	testCaseList := []struct {
		name          string
		mockSetup     func(sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "successful list",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"username", "email", "created_at", "updated_at"}).
					AddRow("user1", "user1@example.com", testTime, testTime)
				mock.ExpectQuery("SELECT username, email, created_at, updated_at FROM users").
					WillReturnRows(rows)
			},

			expectedError: nil,
		},
		{
			name: "database error",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT username, email, created_at, updated_at FROM users").
					WillReturnError(sql.ErrConnDone)
			},
			expectedError: sql.ErrConnDone,
		},
		{
			name: "row scan error",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"username", "email", "created_at", "updated_at"}).
					AddRow("user1", "user1@example.com", testTime, testTime)

				rows.AddRow(nil, "user1@example.com", testTime, testTime)
				mock.ExpectQuery("SELECT").
					WillReturnRows(rows)
			},
			expectedError: assert.AnError,
		},
	}

	for _, testCase := range testCaseList {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockSetup(mock)
			_, err := repo.List()
			if testCase.expectedError == assert.AnError {
				assert.Error(t, err)
			} else {
				assert.Equal(t, testCase.expectedError, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
