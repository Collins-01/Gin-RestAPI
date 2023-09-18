package user

import (
	"database/sql"
	"fmt"
	"rest_api/internal/app/dto"
	"rest_api/internal/app/model"
	"rest_api/internal/utils"
)

type NewUserRepositoryImpl struct {
	db *sql.DB
}

// NewMySQLUserRepository creates a new instance of MySQLUserRepository.
func NewMySQLUserRepository(db *sql.DB) *NewUserRepositoryImpl {
	return &NewUserRepositoryImpl{db: db}
}

func (repo *NewUserRepositoryImpl) CreateUser(user *dto.CreateUserDTO) (int, error) {
	logger := utils.NewLogger(utils.Info)

	// Implement the SQL insert query here
	result, err := repo.db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", user.Username, user.Email)
	lastId, _ := result.LastInsertId()
	message := fmt.Sprintf("Response from creating user: %v", lastId)
	logger.Info(message)
	return int(lastId), err
}

// GetUserByID retrieves a user by their ID from the database.
func (repo *NewUserRepositoryImpl) GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := repo.db.QueryRow("SELECT id, username, email FROM users WHERE id=?", id).
		Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user's information in the database.
func (repo *NewUserRepositoryImpl) UpdateUser(user *model.User) error {
	// Implement the SQL update query here
	_, err := repo.db.Exec("UPDATE users SET username=?, email=? WHERE id=?", user.Username, user.Email, user.ID)
	return err
}

// DeleteUser deletes a user by their ID from the database.
func (repo *NewUserRepositoryImpl) DeleteUser(id int) error {
	// Implement the SQL delete query here
	_, err := repo.db.Exec("DELETE FROM users WHERE id=?", id)
	return err
}
