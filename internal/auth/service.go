package auth

import (
	"errors"
	"go/http/internal/user"

	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	user := user.NewUser(name, email, string(hashedPassword))
	result := service.UserRepository.Database.DB.Create(user)

	if result.Error != nil {
		return "", result.Error
	}

	return user.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)

	if existedUser == nil {
		return "", errors.New(ErrUserNotFound)
	}

	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))

	if err != nil {
		return "", errors.New(ErrWrongPassword)

	}

	return existedUser.Email, nil
}
