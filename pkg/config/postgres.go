package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/samber/do/v2"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func NewPostgresConfig(_ do.Injector) (*PostgresConfig, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	if host == "" || port == "" || user == "" || password == "" || dbName == "" {
		return nil, errors.New("missing required environment variables")
	}

	return &PostgresConfig{
		Host:     host,
		Port:     portInt,
		User:     user,
		Password: password,
		DBName:   dbName,
	}, nil
}
