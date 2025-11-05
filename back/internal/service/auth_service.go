package service

import (
	"github.com/katrin0929/Flowers-store/back/internal/model"
	"github.com/katrin0929/Flowers-store/back/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user model.User) error
	Login(email, password string) (*model.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Register(user model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	return s.userRepo.Create(&user)
}

func (s *authService) Login(email, password string) (*model.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
