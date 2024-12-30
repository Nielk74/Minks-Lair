package usecase

import (
	"cosmink/auth/core/entity"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
	"cosmink/libs/utils"
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	TokenString string
	ExpiresAt   time.Time
}

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

func (a AuthUseCase) Login(username, password string) (Token, error) {
	user, err := a.userRepository.FindByUsername(username)
	if err != nil {
		return Token{}, err
	}
	passhash := sha256.Sum256([]byte(password))
	if user.PassHash != hex.EncodeToString(passhash[:]) {
		return Token{}, fmt.Errorf("invalid password")
	}
	token, err := generateToken(username, user.PassHash)
	return token, nil
}

func generateToken(username string, passhash string) (Token, error) {
	tokenSecret := utils.GetEnvValue("TOKEN_SECRET")
	if tokenSecret == "" {
		return Token{}, errors.New("missing TOKEN_SECRET environment variable")
	}

	// Define token claims
	claims := jwt.MapClaims{
		"username": username,
		"passhash": passhash,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return Token{}, fmt.Errorf("failed to sign token: %v", err)
	}

	return Token{
		TokenString: tokenString,
		ExpiresAt:   time.Now().Add(24 * time.Hour),
	}, nil
}

func (a AuthUseCase) GetUserByToken(tokenString string) (entity.User, error) {
	tokenSecret := utils.GetEnvValue("TOKEN_SECRET")
	// check if the token is valid
	if !validateToken(tokenString, tokenSecret) {
		return entity.User{}, errors.New("invalid token")
	}
	parsedToken , err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return entity.User{}, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return entity.User{}, errors.New("invalid token")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return entity.User{}, errors.New("invalid token")
	}
	return a.userRepository.FindByUsername(username)

}
func validateToken(tokenString string, tokenSecret string) bool {
	if tokenSecret == "" {
		return false
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}
	exp, ok := claims["exp"].(float64)
	if !ok {
		return false
	}
	if time.Unix(int64(exp), 0).Before(time.Now()) {
		return false
	}
	return true
}