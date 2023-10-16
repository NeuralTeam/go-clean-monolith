package repository

import (
	"go-clean-monolith/entity"
	"go-clean-monolith/internal/gateway"
	"go-clean-monolith/pkg/logger"
	"gorm.io/gorm"
)

var _ IUsersRepository = &UsersRepository{}

type (
	// UsersRepository description
	UsersRepository struct {
		db     *gateway.PostgresGateway
		logger logger.Logger
	}

	// IUsersRepository description
	IUsersRepository interface {
		WithTrx(trxHandle *gorm.DB) IUsersRepository
		CreateUser(user entity.Users) error
		GetUser(id int64) (user entity.Users, err error)
		UpdateUser(user entity.Users) error
		DeleteUser(id int64) error
	}
)

// NewUsersRepository description
func NewUsersRepository(db *gateway.PostgresGateway, logger logger.Logger) IUsersRepository {
	return &UsersRepository{
		db:     db,
		logger: logger,
	}
}

// WithTrx description
func (r *UsersRepository) WithTrx(trxHandle *gorm.DB) IUsersRepository {
	if trxHandle == nil {
		r.logger.Error().Msg("transaction database not found in context.")
		return r
	}
	r.db.DB = trxHandle
	return r
}

// CreateUser description
func (r *UsersRepository) CreateUser(user entity.Users) error {
	sql := r.db.Create(&user)
	return sql.Error
}

// GetUser description
func (r *UsersRepository) GetUser(id int64) (user entity.Users, err error) {
	sql := r.db.Find(&user, id)
	return user, sql.Error
}

// UpdateUser description
func (r *UsersRepository) UpdateUser(user entity.Users) error {
	sql := r.db.Save(&user)
	return sql.Error
}

// DeleteUser description
func (r *UsersRepository) DeleteUser(id int64) error {
	sql := r.db.Delete(&entity.Users{}, id)
	return sql.Error
}
