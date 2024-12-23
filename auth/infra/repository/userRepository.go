package repository

import (
	"cosmink/auth/core/entity"
	"cosmink/auth/infra/database"
	"fmt"
)

type UserRepository struct{}

func (r UserRepository) Save(user entity.User) (bool, error) {
	db := database.GetConnection()
	defer db.Close()

	_, err := db.Exec("INSERT INTO users ( username, passhash) VALUES ($1, $2)", user.Username, user.PassHash)

	if err != nil {
		return false, fmt.Errorf("failed to register user: %v", err)
	}
	return true, nil
}

func (r UserRepository) FindByUsername(username string) (entity.User, error) {
	db := database.GetConnection()
	defer db.Close()

	var user entity.User
	err := db.QueryRow("SELECT username, passhash FROM users WHERE username = $1", username).Scan(&user.Username, &user.PassHash)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to find user: %v", err)
	}
	return user, nil
}
