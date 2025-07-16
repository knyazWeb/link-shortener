package auth

import (
	"errors"
	"go/http/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(name, email, password string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)

	if existedUser != nil {
		return "", errors.New(ErrUserExisted)
	}

	user := user.NewUser(name, email, "")
	result := service.UserRepository.Database.DB.Create(user)

	if result.Error != nil {
		return "", result.Error
	}

	return user.Email, nil
}
