package config

import (
	"net"
)

// Config description
type Config struct {
	Version     string `dotenv:"VERSION"`
	Environment string `dotenv:"ENVIRONMENT"`
	ServerHost  net.IP `dotenv:"SERVER_HOST"`
	ServerPort  int    `dotenv:"SERVER_PORT"`

	LogMode     []string `dotenv:"LOG_MODE"`
	LogDirName  string   `dotenv:"LOG_DIR_NAME"`
	LogFileName string   `dotenv:"LOG_FILE_NAME"`
}
