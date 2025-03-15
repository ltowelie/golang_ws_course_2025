package initialization

import (
	"errors"
	"log/slog"
	"os"
)

type Config struct {
	Repo ConfigRepo
}

type ConfigRepo struct {
	DBType string
	DBConn string
}

func initConfig() (*Config, error) {
	slog.Debug("Initializing config repository")

	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "sqlite"
	}
	dbConnStr := os.Getenv("DB_CONN_STR")
	if dbConnStr == "" {
		return nil, errors.New("DB_CONN_STR is empty")
	}
	configRepo := ConfigRepo{DBType: dbType, DBConn: dbConnStr}

	return &Config{Repo: configRepo}, nil
}
