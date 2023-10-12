package config

import (
	"fmt"
	"go-clean-monolith/pkg/dotenv"
	"go.uber.org/fx"
	"log"
	"net"
)

// Env has environment stored
type Env struct {
	Environment string `json:"ENVIRONMENT"`
	ServerHost  net.IP `json:"SERVER_HOST"`
	ServerPort  int    `json:"SERVER_PORT"`

	LogMode     []string `json:"LOG_MODE"`
	LogDirName  string   `json:"LOG_OUTPUT_DIR_NAME"`
	LogFileName string   `json:"LOG_OUTPUT_FILE_NAME"`

	MasterDatabaseHost     string `json:"MASTER_DATABASE_HOST"`
	MasterDatabasePort     int    `json:"MASTER_DATABASE_PORT"`
	MasterDatabaseName     string `json:"MASTER_DATABASE_NAME"`
	MasterDatabaseUser     string `json:"MASTER_DATABASE_USER"`
	MasterDatabasePassword string `json:"MASTER_DATABASE_PASSWORD"`
	MasterDatabaseTimezone string `json:"MASTER_DATABASE_TIMEZONE"`
}

// New creates a new environment
func New() (env Env) {
	err := dotenv.LoadInStruct(&env, ".env")
	if err != nil {
		log.Fatalln(fmt.Sprintf("can not read configuration file. error: %s", err))
	}
	return env
}

// DependencyModules exports dependency for initializing application
var DependencyModules = fx.Options(
	fx.Provide(New),
)
