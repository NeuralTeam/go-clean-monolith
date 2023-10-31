package service

import (
	"go-clean-monolith/internal/entity"
	"go-clean-monolith/internal/repository"
	"go-clean-monolith/pkg/logger"
	"gorm.io/gorm"
)

// UsersService description
type UsersService struct {
	repository *repository.UsersRepository
	logger     logger.Logger
}

// WithTrx description
func (s *UsersService) WithTrx(trxHandle *gorm.DB) *UsersService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// CreateUser description
func (s *UsersService) CreateUser(user entity.Users) error {
	err := s.repository.CreateUser(user)
	return err
}
