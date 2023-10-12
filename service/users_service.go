package service

import (
	"go-clean-monolith/entity"
	"go-clean-monolith/pkg/logger"
	"go-clean-monolith/repository"
	"gorm.io/gorm"
)

var _ IUsersService = &UsersService{}

type (
	// UsersService description
	UsersService struct {
		repository repository.IUsersRepository
		logger     logger.Logger
	}

	// IUsersService description
	IUsersService interface {
		WithTrx(trxHandle *gorm.DB) IUsersService
		CreateUser(user entity.Users) (err error)
	}
)

// NewUsersService description
func NewUsersService(repository repository.IUsersRepository, logger logger.Logger) IUsersService {
	return &UsersService{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx description
func (s *UsersService) WithTrx(trxHandle *gorm.DB) IUsersService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// CreateUser description
func (s *UsersService) CreateUser(user entity.Users) error {
	err := s.repository.CreateUser(user)
	return err
}
