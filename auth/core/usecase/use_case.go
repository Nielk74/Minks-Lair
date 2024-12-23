package usecase

import (
	"cosmink/auth/core/entity"
	"crypto/sha256"
	"encoding/hex"
)

type IUserRepository interface {
	Save(user entity.User) (bool, error)
	FindByUsername(username string) (entity.User, error)
}

type AuthUseCase struct {
	userRepository IUserRepository
}

func NewAuthUseCase(userRepository IUserRepository) AuthUseCase {
	return AuthUseCase{
		userRepository: userRepository,
	}
}

func (a AuthUseCase) Register(username, password string) (bool, error) {
	passhash := sha256.Sum256([]byte(password))
	user := entity.NewUser(username, hex.EncodeToString(passhash[:]))
	return a.userRepository.Save(*user)
}