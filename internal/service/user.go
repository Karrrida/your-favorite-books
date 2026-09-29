package service

import (
	"fmt"

	"github.com/Karrrida/your-favorite-books/internal/auth"
	"github.com/Karrrida/your-favorite-books/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Register(email, name, password string) error {

	password_hash, err := auth.Hash(password)

	if err != nil {
		fmt.Println("Error in crypto", err)
		return err
	}

	return s.userRepository.CreateUser(email, name, string(password_hash))
}

func (s *UserService) GetByEmail(email string) (*repository.User, error) {
	return s.userRepository.GetUserByEmail(email)
}
