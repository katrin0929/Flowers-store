package service

import (
	"Flowers-store/internal/model"
	"Flowers-store/internal/repository"
	"errors"
)

type RegistrationService struct {
	repo repository.Repository
}

func NewRegistrationService(repo repository.Repository) *RegistrationService {
	return &RegistrationService{
		repo: repo,
	}
}

func (rs *RegistrationService) Register(user *model.User) error {
	if len(user.Username) == 0 || len(user.Password) == 0 {
		return errors.New("invalid input data")
	}
	err := rs.repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
