package service

import (
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

func (s *UserService) Register(email, name string) error {
	return s.userRepository.CreateUser(email, name)
}

func (s *UserService) GetByEmail(email string) (*repository.User, error) {
	return s.userRepository.GetUserByEmail(email)
}
