package config

import (
	"time"

	"github.com/trangnkp/my_books/src/internal/db"
)

const (
	serverHost = "0.0.0.0"
	serverPort = ":8896"

	dbUser                      = "admin"
	dbPassword                  = "password"
	dbServer                    = "0.0.0.0:3306"
	dbSchema                    = "my_books"
	dbConnectionLifetimeSeconds = 300
	dbMaxIdleConnection         = 50
	dbMaxOpenConnection         = 100
)

type AppConfig struct {
	Server *ServerConfig
	DB     *db.MySQLConfig
}

type ServerConfig struct {
	Port              string
	Address           string
	ReadHeaderTimeout time.Duration
}

func New() *AppConfig {
	return &AppConfig{
		Server: NewServerConfig(),
		DB:     NewDBConfig(),
	}
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:              serverPort,
		Address:           serverHost + serverPort,
		ReadHeaderTimeout: 5 * time.Second,
	}
}

func NewDBConfig() *db.MySQLConfig {
	return &db.MySQLConfig{
		Server:                    dbServer,
		Schema:                    dbSchema,
		User:                      dbUser,
		Password:                  dbPassword,
		ConnectionLifetimeSeconds: dbConnectionLifetimeSeconds,
		MaxIdleConnections:        dbMaxIdleConnection,
		MaxOpenConnections:        dbMaxOpenConnection,
	}
}
