package configs

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	PgSqlDSN      string
	WebServerPort string
	SecretKey     string
}

func New() (*Config, error) {
	pqSqlHost := os.Getenv("PGSQL_HOST")
	if len(pqSqlHost) == 0 {
		return nil, errors.New("please specify PGSQL_HOST variable in env")
	}

	pqSqlDB := os.Getenv("PGSQL_DB")
	if len(pqSqlDB) == 0 {
		return nil, errors.New("please specify PGSQL_DB variable in env")
	}

	pqSqlUser := os.Getenv("PGSQL_USER")
	if len(pqSqlUser) == 0 {
		return nil, errors.New("please specify PGSQL_USER variable in env")
	}

	pqSqlPassword := os.Getenv("PGSQL_PASSWORD")
	if len(pqSqlPassword) == 0 {
		return nil, errors.New("please specify PGSQL_PASSWORD variable in env")
	}

	pqSqlPort := os.Getenv("PGSQL_PORT")
	if len(pqSqlPort) == 0 {
		return nil, errors.New("please specify PGSQL_PORT variable in env")
	}

	pgSqlSSLMode := os.Getenv("PGSQL_SSL_MODE")
	if len(pgSqlSSLMode) == 0 {
		pgSqlSSLMode = "disable"
	}

	pgSqlDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		pqSqlUser, pqSqlPassword, pqSqlHost, pqSqlPort, pqSqlDB, pgSqlSSLMode)
	webServerPort := os.Getenv("WEB_SERVER_PORT")
	if len(webServerPort) == 0 {
		return nil, errors.New("please specify WEB_SERVER_PORT variable in env")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if len(secretKey) == 0 {
		return nil, errors.New("please specify SECRET_KEY variable in env")
	}

	return &Config{
		PgSqlDSN:      pgSqlDSN,
		WebServerPort: webServerPort,
		SecretKey:     secretKey,
	}, nil
}
