package repositories

import (
	"database/sql"
	"golang-template/app/models"
)

type UserRepository interface {
	Create(user *models.UserRegister) error
	Update(user *models.UserUpdatePassword) error
	List() (*[]models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {

	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.UserRegister) error {

	query := `
		INSERT INTO users (username, email, password)
		VALUES (?, ?, ?)
	`
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Update(user *models.UserUpdatePassword) error {
	query := `
		UPDATE users SET username = ?, email = ?, password = ?
		WHERE username = ?
	`
	_, err := r.db.Exec(query, user.Username, user.NewPassword, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) List() (*[]models.User, error) {

	query := `
		SELECT username, email, created_at, updated_at FROM users
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}
