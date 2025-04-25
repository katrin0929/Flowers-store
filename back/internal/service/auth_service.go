package service

import (
	"Flowers-store/internal/repository"
)

type AuthService struct {
	repo repository.Repository
}

func NewAuthService(repo repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) Login(username, password string) bool {
	user, err := a.repo.GetByUsername(username)
	if err != nil || user == nil {
		return false
	}

	return user.Password == password
}
