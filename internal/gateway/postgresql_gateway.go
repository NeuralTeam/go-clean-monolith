package gateway

import (
	"fmt"
	"go-clean-monolith/config"
	"go-clean-monolith/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresGateway description
type PostgresGateway struct {
	*gorm.DB
}

// NewPostgresGateway description
func NewPostgresGateway(env config.Env, logger logger.Logger) *PostgresGateway {
	url := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s",
		env.MasterDatabaseHost, env.MasterDatabaseUser, env.MasterDatabasePassword,
		env.MasterDatabaseName, env.MasterDatabasePort, env.MasterDatabaseTimezone,
	)

	logger.Info().Msg("attempt to connect to the database")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{Logger: logger.Gorm})
	if err != nil {
		logger.Fatal().Err(err).Msg(url)
	}
	logger.Info().Msg("database is connected")

	return &PostgresGateway{
		DB: db,
	}
}
