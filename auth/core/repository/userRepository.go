package repository

import (
	"cosmink/auth/core/entity"
	"cosmink/auth/infra/database"
	"fmt"
)

func Save(user entity.User) (bool, error) {
	db := database.GetConnection()
	defer db.Close()

	_, err := db.Exec("INSERT INTO users ( username, pass_hash) VALUES ($1, $2)", user.Username, user.PassHash)

	if err != nil {
		return false, fmt.Errorf("failed to register user: %v", err)
	}
	return true, nil
}
